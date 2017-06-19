package sensors

import (
	"net"
	"bytes"
	"gopkg.in/vmihailenco/msgpack.v2"
	log "github.com/Sirupsen/logrus"
	"ihmc.us/nodemon/sensors/disservice"
	"ihmc.us/nodemon/measure"
	"strconv"
	subjdisservice "ihmc.us/nodemon/subjects/disservice"
	"ihmc.us/nodemon/util"
	"github.com/golang/protobuf/ptypes/timestamp"
	"time"
)

type DisServiceSensor struct {
	port       uint16                  // the port of the sensor
	disservice chan<- *measure.Measure // a channel to return disservice data from the sensor
	logDebug   bool                    // decides whether enabling debug log from this specific sensor
	tag        string                  // a tag for the log
}

func NewDisServiceSensor(port uint16, disservice chan<- *measure.Measure, logDebug bool) *DisServiceSensor {

	ds := new(DisServiceSensor)
	ds.port = port
	ds.disservice = disservice
	ds.tag = "DisService"
	ds.logDebug = logDebug

	return ds
}

func (ds *DisServiceSensor) Start() (err error) {
	conn, err := createUDPConn(int(ds.port), ds.TAG())
	if err != nil {
		return err
	}
	go ds.handleConn(conn)
	return nil
}

func (ds *DisServiceSensor) handleConn(conn *net.UDPConn) {
	for {
		var buf [disservice.UDPPacketMaxSize]byte
		size, addr, err := conn.ReadFromUDP(buf[0:])
		if err != nil {
			log.Error(err.Error())
			continue
		}

		log.Info(ds.TAG(), "received packet, size ", strconv.Itoa(size), " src: ", addr.String())
		go ds.parsePacket(addr.IP, buf[:size])
		//log.Debug(TAG(MocketsTag) + "handleConn() generated so far: " + strconv.Itoa(counter) + " measures")
	}
}

func (ds *DisServiceSensor) parsePacket(sensorIP net.IP, buf []byte) {
	var (
		decoder *msgpack.Decoder
		err     error
	)
	reader := bytes.NewReader(buf[0:])
	decoder = msgpack.NewDecoder(reader)
	var (
		header         int16
		peerId         string
		packetType     int16
		basicStat      *disservice.BasicStatisticsInfo
		statInfo       *disservice.StatsInfo
		dupTrafficInfo *disservice.DuplicateTrafficInfo
		ordinal        int16
	)
	header, err = decoder.DecodeInt16()
	checkError(err)
	peerId, err = decoder.DecodeString()
	checkError(err)
	if ds.logDebug {
		log.Debug(ds.TAG(), "header: ", header, " peerId: ", peerId)
	}
	packetType, err = decoder.DecodeInt16()
	checkError(err)
	if ds.logDebug {
		log.Debug(ds.TAG(), "packetType: ", packetType)
	}
	basicStat = new(disservice.BasicStatisticsInfo)
	basicStat.DataMessageReceived, err = decoder.DecodeInt64()
	checkError(err)
	basicStat.DataBytesReceived, err = decoder.DecodeInt64()
	checkError(err)
	basicStat.DataFragmentsReceived, err = decoder.DecodeInt64()
	checkError(err)
	basicStat.DataFragmentBytesReceived, err = decoder.DecodeInt64()
	checkError(err)
	basicStat.MissingFragmentRequestMessagesSent, err = decoder.DecodeInt64()
	checkError(err)
	basicStat.MissingFragmentRequestBytesSent, err = decoder.DecodeInt64()
	checkError(err)
	basicStat.MissingFragmentRequestMessagesReceived, err = decoder.DecodeInt64()
	checkError(err)
	basicStat.MissingFragmentRequestBytesReceived, err = decoder.DecodeInt64()
	checkError(err)
	basicStat.DataCacheQueryMessagesSent, err = decoder.DecodeInt64()
	checkError(err)
	basicStat.DataCacheQueryBytesSent, err = decoder.DecodeInt64()
	checkError(err)
	basicStat.DataCacheQueryMessagesReceived, err = decoder.DecodeInt64()
	checkError(err)
	basicStat.DataCacheQueryBytesReceived, err = decoder.DecodeInt64()
	checkError(err)
	basicStat.TopologyStateMessagesSent, err = decoder.DecodeInt64()
	checkError(err)
	basicStat.TopologyStateBytesSent, err = decoder.DecodeInt64()
	checkError(err)
	basicStat.TopologyStateMessagesReceived, err = decoder.DecodeInt64()
	checkError(err)
	basicStat.TopologyStateBytesReceived, err = decoder.DecodeInt64()
	checkError(err)
	basicStat.KeepAliveMessagesSent, err = decoder.DecodeInt64()
	checkError(err)
	basicStat.KeepAliveMessagesReceived, err = decoder.DecodeInt64()
	checkError(err)
	basicStat.QueryMessagesSent, err = decoder.DecodeInt64()
	checkError(err)
	basicStat.QueryMessagesReceived, err = decoder.DecodeInt64()
	checkError(err)
	basicStat.QueryHitsMessagesSent, err = decoder.DecodeInt64()
	checkError(err)
	basicStat.QueryHitsMessagesReceived, err = decoder.DecodeInt64()
	checkError(err)


	// sending measure back
	m := toDisServiceMeasure(basicStat)
	if ds.logDebug {
		log.Debug(ds.TAG(), "Measure: ", m.String())
	}

	//putting measure inside channel
	ds.disservice <- m

	ordinal, err = decoder.DecodeInt16()
	if ordinal == int16(disservice.DSSF_OverallStats) {
		statInfo = &disservice.StatsInfo{PeerId: peerId}
		statInfo.ClientMessagesPushed, err = decoder.DecodeInt64()
		checkError(err)
		statInfo.ClientBytesPushed, err = decoder.DecodeInt64()
		checkError(err)
		statInfo.ClientMessagesMadeAvailable, err = decoder.DecodeInt64()
		checkError(err)
		statInfo.ClientBytesMadeAvailable, err = decoder.DecodeInt64()
		checkError(err)
		statInfo.FragmentsPushed, err = decoder.DecodeInt64()
		checkError(err)
		statInfo.FragmentBytesPushed, err = decoder.DecodeInt64()
		checkError(err)
		statInfo.OnDemandFragmentsSent, err = decoder.DecodeInt64()
		checkError(err)
		statInfo.OnDemandFragmentBytesSent, err = decoder.DecodeInt64()
		checkError(err)
		//print
		//fmt.Println(statInfo)
	}
	var duplicateTraff int16
	duplicateTraff, err = decoder.DecodeInt16()
	if duplicateTraff == int16(disservice.DSSF_DuplicateTrafficInfo) {
		dupTrafficInfo = &disservice.DuplicateTrafficInfo{PeerId: peerId}
		dupTrafficInfo.OverheardDuplicateTraffic, err = decoder.DecodeInt64()
		checkError(err)
		dupTrafficInfo.TargetedDuplicateTraffic, err = decoder.DecodeInt64()
		checkError(err)
		//print
		//fmt.Println(dupTrafficInfo)
	}
	var flag int16
	flag, err = decoder.DecodeInt16()
	checkError(err)
	for ; flag == int16(disservice.DSSF_PerClientGroupTagStats); flag, _ = decoder.DecodeInt16() {
		statPerPeer := &disservice.BasicStatisticsInfoByPeer{PeerId: peerId}
		statPerPeer.RemotePeerId, err = decoder.DecodeString()
		checkError(err)
		statPerPeer.DataMessageReceived, err = decoder.DecodeInt64()
		checkError(err)
		statPerPeer.DataBytesReceived, err = decoder.DecodeInt64()
		checkError(err)
		statPerPeer.DataFragmentReceived, err = decoder.DecodeInt64()
		checkError(err)
		statPerPeer.DataFragmentBytesReceived, err = decoder.DecodeInt64()
		checkError(err)
		statPerPeer.MissingFragmentRequestMessagesReceived, err = decoder.DecodeInt64()
		checkError(err)
		statPerPeer.MissingFragmentRequestBytesReceived, err = decoder.DecodeInt64()
		checkError(err)
		statPerPeer.KeepAliveMessagesReceived, err = decoder.DecodeInt64()
		checkError(err)
		//fmt.Println(statPerPeer)
	}
	var flagend int32
	flagend, err = decoder.DecodeInt32()
	if ds.logDebug {
		log.Debug(ds.TAG(), "Flagend ", flagend)
	}
}

func toDisServiceMeasure(basicStat *disservice.BasicStatisticsInfo) *measure.Measure {

	m := &measure.Measure{
		Subject:   measure.Subject_disservice,
		Strings:   make(map[string]string),
		Integers:  make(map[string]int64),
		Doubles:   make(map[string]float64),
		Timestamp: &timestamp.Timestamp{Seconds: time.Now().Unix()}, //TODO:time override

	}

	m.GetIntegers()[subjdisservice.Int_data_message_received.String()] = basicStat.DataMessageReceived
	m.GetIntegers()[subjdisservice.Int_data_bytes_received.String()] = basicStat.DataBytesReceived
	m.GetIntegers()[subjdisservice.Int_data_fragments_received.String()] = basicStat.DataFragmentsReceived
	m.GetIntegers()[subjdisservice.Int_data_fragment_bytes_received.String()] = basicStat.DataFragmentBytesReceived
	m.GetIntegers()[subjdisservice.Int_missing_fragment_request_messages_sent.String()] = basicStat.MissingFragmentRequestMessagesSent
	m.GetIntegers()[subjdisservice.Int_missing_fragment_request_bytes_sent.String()] = basicStat.MissingFragmentRequestBytesSent
	m.GetIntegers()[subjdisservice.Int_missing_fragment_request_messages_received.String()] = basicStat.MissingFragmentRequestBytesReceived
	m.GetIntegers()[subjdisservice.Int_missing_fragment_request_bytes_received.String()] = basicStat.MissingFragmentRequestBytesReceived
	m.GetIntegers()[subjdisservice.Int_data_cache_query_messages_sent.String()] = basicStat.DataCacheQueryMessagesSent
	m.GetIntegers()[subjdisservice.Int_data_cache_query_bytes_sent.String()] = basicStat.DataCacheQueryBytesSent
	m.GetIntegers()[subjdisservice.Int_data_cache_query_messages_received.String()] = basicStat.DataCacheQueryMessagesReceived
	m.GetIntegers()[subjdisservice.Int_data_cache_query_bytes_received.String()] = basicStat.DataCacheQueryBytesReceived
	m.GetIntegers()[subjdisservice.Int_topology_state_messages_sent.String()] = basicStat.TopologyStateMessagesSent
	m.GetIntegers()[subjdisservice.Int_topology_state_bytes_sent.String()] = basicStat.TopologyStateBytesSent
	m.GetIntegers()[subjdisservice.Int_topology_state_messages_received.String()] = basicStat.TopologyStateMessagesReceived
	m.GetIntegers()[subjdisservice.Int_topology_state_bytes_received.String()] = basicStat.TopologyStateBytesReceived
	m.GetIntegers()[subjdisservice.Int_keep_alive_messages_sent.String()] = basicStat.KeepAliveMessagesSent
	m.GetIntegers()[subjdisservice.Int_keep_alive_messages_received.String()] = basicStat.KeepAliveMessagesReceived
	m.GetIntegers()[subjdisservice.Int_query_messages_sent.String()] = basicStat.QueryMessagesSent
	m.GetIntegers()[subjdisservice.Int_query_messages_received.String()] = basicStat.QueryMessagesReceived
	m.GetIntegers()[subjdisservice.Int_query_hits_messages_sent.String()] = basicStat.QueryHitsMessagesSent
	m.GetIntegers()[subjdisservice.Int_query_hits_messages_received.String()] = basicStat.QueryHitsMessagesReceived

	return m
}


func (ds *DisServiceSensor) TAG() string {
	return ds.tag + " [" + util.GetGIDString() + "] "
}
