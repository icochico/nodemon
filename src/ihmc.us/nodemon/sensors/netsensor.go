package sensors

import (
	log "github.com/Sirupsen/logrus"
	"net"
	"strconv"
	"github.com/golang/protobuf/proto"
	"ihmc.us/nodemon/measure"
	"ihmc.us/nodemon/sensors/netsensor"
	"ihmc.us/nodemon/subjects/traffic"
	"ihmc.us/nodemon/subjects/host"
	"ihmc.us/nodemon/subjects/network"
	"ihmc.us/nodemon/util"
	"time"
	"github.com/golang/protobuf/ptypes/timestamp"
	"errors"
)

// values for type
const ()

type NetSensor struct {
	port     uint16                  // the port of the sensor
	traffic  chan<- *measure.Measure // a channel to return traffic data from the sensor
	host     chan<- *measure.Measure // a channel to return host data from the sensor
	network  chan<- *measure.Measure // a channel to return network data from the sensor
	logDebug bool                    // decides whether enabling debug log from this specific sensor
	tag      string                  // a tag for the log
}

func NewNetSensor(port uint16, traffic, host, network chan<- *measure.Measure, logDebug bool) *NetSensor {

	ns := new(NetSensor)
	ns.port = port
	ns.traffic = traffic
	ns.host = host
	ns.network = network
	ns.logDebug = logDebug
	ns.tag = "NetSensor"
	//ns.wg = wg
	return ns
}

func (ns *NetSensor) Start() (err error) {

	conn, err := createUDPConn(int(ns.port), ns.TAG())
	if err != nil {
		return err
	}
	go ns.handleConn(conn)
	return nil
}

func (ns *NetSensor) handleConn(conn *net.UDPConn) {
	for {
		var buf [netsensor.UDPPacketMaxSize]byte
		size, addr, err := conn.ReadFromUDP(buf[0:])
		if err != nil {
			log.Error(err.Error())
			continue
		}

		log.Info(ns.TAG(), "received packet, size ", strconv.Itoa(size), " src: ", addr.String())

		go ns.handlePacket(buf[:size])
	}
}

func (ns *NetSensor) handlePacket(buf []byte) {
	start := time.Now()

	nsc := &netsensor.NetSensorContainer{}
	err := proto.Unmarshal(buf, nsc)
	if err != nil {
		log.Warn(ns.TAG()+"Unmarshalling error: ", err)
		//ns.wg.Done()
		return
	}

	//parsing NetSensorTag's Container
	//log.Debug(ns.TAG() + "parsePacket() packet: " + nsc.String())
	ns.parsePacket(nsc)

	end := time.Since(start)
	if ns.logDebug {
		log.Debug(ns.TAG() + "handlePacket() ex time: " + end.String())
	}
}

func (ns *NetSensor) parsePacket(nsc *netsensor.NetSensorContainer) {
	start := time.Now()
	countTraffic, countHost, countNetwork := 0, 0, 0

	if ns.logDebug {
		log.Debug(ns.TAG() + nsc.String())
	}

	switch nsc.DataType {
	case netsensor.DataType_TRAFFIC:
		for _, tbi := range nsc.GetTrafficByInterfaces() {

			//current sensor ip is represented in monitoring interface (NetSensor)
			sensorIp := tbi.GetMonitoringInterface()

			for _, flow := range tbi.GetMicroflows() {
				//fetch src and dest host for this Flow
				srcIP := util.UInt32ToIP(flow.GetIpSrc()).String()
				destIP := util.UInt32ToIP(flow.GetIpDst()).String()
				for _, stat := range flow.GetStats() {
					if ns.logDebug {
						log.Debug(ns.TAG(), stat)
					}

					//fetch src and dest port from this Stat
					srcPort := stat.GetSrcPort()
					destPort := stat.GetDstPort()
					protocol := stat.GetProtocol()
					//initialize Measure
					m := &measure.Measure{
						Subject:   measure.Subject_traffic,
						Strings:   make(map[string]string),
						Integers:  make(map[string]int64),
						Doubles:   make(map[string]float64),
						Timestamp: nsc.GetTimestamp(),
					}
					m.GetStrings()[traffic.Str_sensor_ip.String()] = sensorIp
					m.GetStrings()[traffic.Str_src_ip.String()] = srcIP
					m.GetStrings()[traffic.Str_dest_ip.String()] = destIP
					m.GetStrings()[traffic.Str_src_port.String()] = strconv.Itoa(int(srcPort))
					m.GetStrings()[traffic.Str_dest_port.String()] = strconv.Itoa(int(destPort))
					m.GetStrings()[traffic.Str_protocol.String()] = protocol
					statType := stat.GetStatType()
					switch statType {
					case netsensor.StatType_TRAFFIC_AVERAGE:
						avgs := stat.GetAverages()
						if avgs == nil || len(avgs) == 0 {
							if ns.logDebug {
								log.Warn(ns.TAG() +
									statType.String() +
									" value not found")
							}
							continue
						}
						//save only first average (should be the only one)
						//only if value is != 0
						var isDataSane bool = false
						sent := avgs[0].GetSent()
						if sent != 0 {
							m.GetIntegers()[traffic.Int_sent.String()] = int64(sent)
							isDataSane = true
						}
						received := avgs[0].GetReceived()
						if received != 0 {
							m.GetIntegers()[traffic.Int_received.String()] = int64(received)
							isDataSane = true
						}
						observed := avgs[0].GetObserved()
						if observed != 0 {
							m.GetIntegers()[traffic.Int_observed.String()] = int64(observed)
							isDataSane = true
						}
						resolution := avgs[0].GetResolution()
						if resolution != 0 {
							m.GetIntegers()[traffic.Int_resolution.String()] = int64(resolution)
						}

						if !isDataSane {
							log.Warn("None of sent, received or observed are filled in " +
								srcIP + " -> " + destIP)
						}

						//append measure to slice
						select {
						case ns.traffic <- m:
							countTraffic++
						default:
							log.Error(ns.TAG() +
								"parsePacket(): traffic channel is full, dropping measure")
						}
						countTraffic++
					case netsensor.StatType_PACKETS:
						if ns.logDebug {
							log.Warn(ns.TAG() + statType.String() +
								" not supported")
						}
						continue
					}
				}
			}
		}
	case netsensor.DataType_TOPOLOGY:
		for _, topo := range nsc.GetTopologies() {
			if ns.logDebug {
				log.Debug(ns.TAG(), topo)
			}

			ni := topo.GetNetworkInfo()
			mNet := networkInfoToMeasure(ni, nsc.GetTimestamp(), "") //no type
			err := ns.sendTo(ns.network, mNet)
			if err == nil {
				countNetwork++
			}

			for _, internal := range topo.GetInternals() {
				mHost := hostToMeasure(ni, internal, nsc.Timestamp,
					netsensor.HostTypeDefault)
				//send host information
				err := ns.sendTo(ns.host, mHost)
				if err == nil {
					countHost++
				}
			}

			for _, localGw := range topo.GetLocalGws() {

				mLocalGw := hostToMeasure(ni, localGw, nsc.Timestamp,
					netsensor.HostTypeGateway)
				//send host information
				err := ns.sendTo(ns.host, mLocalGw)
				if err == nil {
					countHost++
				}
			}
		}

	case netsensor.DataType_NETPROXY:
		npi := nsc.GetNetProxyInfo()
		if npi == nil {
			log.Error(ns.TAG() + "parsePacket(): NetProxyInfo is null, skipping")
			return
		}

		if npi.GetInternal() == nil {
			log.Warn(ns.TAG() + " parsePacket(): NetProxyInfo internal is null, skipping")
			return

		}
		mInternal := networkInfoToMeasure(npi.GetInternal(), nsc.GetTimestamp(), netsensor.NetworkTypeInternal)
		err := ns.sendTo(ns.network, mInternal)
		if ns.logDebug {
			log.Debug(ns.TAG() + "Measure: " + mInternal.String())
		}
		if err == nil {
			countNetwork++
		}

		if npi.GetExternal() == nil {
			log.Warn(ns.TAG() + " parsePacket(): NetProxyInfo external is null, skipping")
			return
		}
		mExternal := networkInfoToMeasure(npi.GetExternal(), nsc.GetTimestamp(), netsensor.NetworkTypeExternal)
		if ns.logDebug {
			log.Debug(ns.TAG() + "Measure: " + mExternal.String())
		}
		err = ns.sendTo(ns.network, mExternal)
		if err == nil {
			countNetwork++
		}

		//create one host for each remote NetProxy
		for _, rNetProxyIP := range npi.GetRemoteNetProxyIPs() {
			mHost := &measure.Measure{
				Subject:   measure.Subject_host,
				Strings:   make(map[string]string),
				Integers:  make(map[string]int64),
				Doubles:   make(map[string]float64),
				Timestamp: nsc.GetTimestamp(),
			}
			mHost.GetStrings()[host.Str_sensor_ip.String()] = util.UInt32ToIP(npi.GetExternal().InterfaceIp).String()
			mHost.GetStrings()[host.Str_host_ip.String()] = util.UInt32ToIP(rNetProxyIP).String()
			mHost.GetStrings()[host.Str_type.String()] = netsensor.HostTypeNetProxy
			if ns.logDebug {
				log.Debug(ns.TAG() + "Measure: " + mHost.String())
			}
			err := ns.sendTo(ns.host, mHost)
			if err == nil {
				countHost++
			}
		}
	}

	end := time.Since(start)
	if ns.logDebug {
		log.Debug(ns.TAG() + "parsePacket() type: " + nsc.DataType.String() + " produced measures." +
			" traffic: " + strconv.Itoa(countTraffic) +
			" host: " + strconv.Itoa(countHost) +
			" network: " + strconv.Itoa(countNetwork) +
			" in time: " + end.String())
	}
}

func (ns *NetSensor) sendTo(channel chan<- *measure.Measure, measure *measure.Measure) error {

	select {
	case channel <- measure:
		return nil
	default:
		errorStr := "parsePacket(): " + measure.Subject.String() + " channel is full, dropping measure";
		log.Error(ns.TAG() + errorStr)
		return errors.New(errorStr)
	}
}

func networkInfoToMeasure(ni *netsensor.NetworkInfo, timestamp *timestamp.Timestamp,
	typeStr string) *measure.Measure {

	sensorIp := util.UInt32ToIP(ni.GetInterfaceIp()).String()
	networkIp := ni.GetNetworkName()
	networkMask := ni.GetNetworkNetmask()
	mNet := &measure.Measure{
		Subject:   measure.Subject_network,
		Strings:   make(map[string]string),
		Integers:  make(map[string]int64),
		Doubles:   make(map[string]float64),
		Timestamp: timestamp,
	}
	mNet.GetStrings()[network.Str_sensor_ip.String()] = sensorIp
	mNet.GetStrings()[network.Str_network_ip.String()] = networkIp
	mNet.GetStrings()[network.Str_netmask.String()] = networkMask
	mNet.GetStrings()[network.Str_type.String()] = typeStr

	return mNet

}

func hostToMeasure(ni *netsensor.NetworkInfo, h *netsensor.Host,
	timestamp *timestamp.Timestamp, typeStr string) *measure.Measure {

	sensorIp := util.UInt32ToIP(ni.GetInterfaceIp()).String()
	networkIp := ni.GetNetworkName()
	networkMask := ni.GetNetworkNetmask()

	mHost := &measure.Measure{
		Subject:   measure.Subject_host,
		Strings:   make(map[string]string),
		Integers:  make(map[string]int64),
		Doubles:   make(map[string]float64),
		Timestamp: timestamp,
	}

	mHost.GetStrings()[host.Str_sensor_ip.String()] = sensorIp
	mHost.GetStrings()[host.Str_host_ip.String()] = util.UInt32ToIP(h.Ip).String()
	mHost.GetStrings()[host.Str_mac.String()] = h.Mac
	mHost.GetStrings()[host.Str_network_ip.String()] = networkIp
	mHost.GetStrings()[host.Str_netmask.String()] = networkMask
	mHost.GetStrings()[host.Str_type.String()] = typeStr
	return mHost
}

func (ns *NetSensor) Stop() {
	//ns.wg.Done()
}

func (ns *NetSensor) TAG() string {
	return ns.tag + " [" + util.GetGIDString() + "] "
}
