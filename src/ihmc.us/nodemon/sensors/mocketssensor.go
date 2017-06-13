package sensors

import (
	"ihmc.us/nodemon/measure"
	log "github.com/Sirupsen/logrus"
	"strconv"
	"net"
	"time"
	"bytes"
	"encoding/binary"
	"io"
	"ihmc.us/nodemon/util"
	"ihmc.us/nodemon/sensors/mockets"
	subjmockets "ihmc.us/nodemon/subjects/mockets"
	"github.com/golang/protobuf/ptypes/timestamp"
)

// Reads statistics from IHMC's Mockets library
type MocketsSensor struct {
	port     uint16                  // the port of the sensor
	mockets  chan<- *measure.Measure // a channel to return mockets data from the sensor
	logDebug bool                    // decides whether enabling debug log from this specific sensor
	tag      string                  // a tag for the log
}

func NewMocketSensor(port uint16, mockets chan<- *measure.Measure, logDebug bool) *MocketsSensor {

	ms := new(MocketsSensor)
	ms.port = port
	ms.mockets = mockets
	ms.tag = "Mockets"
	ms.logDebug = logDebug
	//ns.wg = wg
	return ms
}

func (ms *MocketsSensor) Start() (err error) {

	conn, err := createUDPConn(int(ms.port), ms.TAG())
	if err != nil {
		return err
	}
	go ms.handleConn(conn)

	return nil
}

func (ms *MocketsSensor) handleConn(conn *net.UDPConn) {

	//counter := 0
	for {
		var buf [mockets.UDPPacketMaxSize]byte
		size, addr, err := conn.ReadFromUDP(buf[0:])
		if err != nil {
			log.Error(err.Error())
			continue
		}

		log.Info(ms.TAG(), "received packet, size ", strconv.Itoa(size), " src: ", addr.String())
		go ms.parsePacket(addr.IP, buf[:size])
		//log.Debug(TAG(MocketsTag) + "handleConn() generated so far: " + strconv.Itoa(counter) + " measures")
	}

}

func (ms *MocketsSensor) parsePacket(sensorIP net.IP, buf []byte) {
	start := time.Now()
	if ms.logDebug {
		defer log.Debug(ms.TAG()+" parsePacket() ex time: ", time.Since(start))
	}

	buffer := bytes.NewBuffer(buf[0:])

	/// END POINT INFO ///
	epi := &mockets.EndPointsInfo{}

	// msg type
	msgType, err := buffer.ReadByte()
	checkError(err)
	if ms.logDebug {
		log.Debug(ms.TAG(), "msgType: ", msgType)
	}

	var PID uint32
	err = binary.Read(buffer, binary.LittleEndian, &PID)
	checkError(err)
	epi.PID = int64(PID)
	if ms.logDebug {
		log.Debug(ms.TAG(), "PID: ", PID)
	}
	var idLen uint16
	err = binary.Read(buffer, binary.LittleEndian, &idLen)
	checkError(err)
	if ms.logDebug {
		log.Debug(ms.TAG(), "idLength: ", idLen)
	}
	if idLen > 0 {
		id := make([]byte, idLen)
		io.ReadFull(buffer, id)
		identifier := string(id[:idLen])
		epi.Identifier = identifier
		log.Debug(ms.TAG(), "identifier: "+identifier)
	}
	buffer.ReadByte() //string terminator

	//local address (use binary.BigEndian for ip addresses for network order)
	var localAddr uint32
	err = binary.Read(buffer, binary.BigEndian, &localAddr)
	checkError(err)
	epi.LocalAddr = int64(localAddr)
	if ms.logDebug {
		log.Debug(ms.TAG(), "localAddr: ", util.UInt32ToIP(localAddr))
	}
	//local port
	var localPort uint16
	err = binary.Read(buffer, binary.LittleEndian, &localPort)
	checkError(err)
	epi.LocalPort = int64(localPort)
	if ms.logDebug {
		log.Debug(ms.TAG(), "localPort: ", localPort)
	}
	//remote address (use binary.BigEndian for ip addresses for network order)
	var remoteAddr uint32
	err = binary.Read(buffer, binary.BigEndian, &remoteAddr)
	checkError(err)
	epi.RemoteAddr = int64(remoteAddr)
	if ms.logDebug {
		log.Debug(ms.TAG(), "remoteAddr: ", util.UInt32ToIP(remoteAddr))
	}
	//remote port
	var remotePort uint16
	err = binary.Read(buffer, binary.LittleEndian, &remotePort)
	checkError(err)
	epi.RemotePort = int64(remotePort)
	if ms.logDebug {
		log.Debug(ms.TAG(), "remotePort: ", remotePort)
	}

	if msgType == uint8(mockets.MSNT_Stats) {
		if ms.logDebug {
			log.Debug("MsgType is MSNT_Stats (" + strconv.Itoa(int(mockets.MSNT_Stats)) + ")")
		}

		/// STATISTICS INFO ///
		si := &mockets.StatisticsInfo{}
		var lastContactTime int64
		err = binary.Read(buffer, binary.LittleEndian, &lastContactTime)
		checkError(err)
		if ms.logDebug {
			log.Debug(ms.TAG(), "lastContactTime: ", time.Unix(lastContactTime, 0))
		}

		//TODO check why this time doesn't seem correct
		si.LastContactTime = lastContactTime

		var sentBytes uint32
		err = binary.Read(buffer, binary.LittleEndian, &sentBytes)
		checkError(err)
		if ms.logDebug {
			log.Debug(ms.TAG(), "sentBytes: ", sentBytes)
		}
		si.SentBytes = int64(sentBytes)

		var sentPackets uint32
		err = binary.Read(buffer, binary.LittleEndian, &sentPackets)
		checkError(err)
		if ms.logDebug {
			log.Debug(ms.TAG(), "sentPackets: ", sentPackets)
		}
		si.SentPackets = int64(sentPackets)

		var retransmits uint32
		err = binary.Read(buffer, binary.LittleEndian, &retransmits)
		checkError(err)
		if ms.logDebug {
			log.Debug(ms.TAG(), "retransmits: ", retransmits)
		}
		si.Retransmits = int64(retransmits)

		var receivedBytes uint32
		err = binary.Read(buffer, binary.LittleEndian, &receivedBytes)
		checkError(err)
		if ms.logDebug {
			log.Debug(ms.TAG(), "receivedBytes: ", receivedBytes)
		}
		si.ReceivedBytes = int64(receivedBytes)

		var receivedPackets uint32
		err = binary.Read(buffer, binary.LittleEndian, &receivedPackets)
		checkError(err)
		if ms.logDebug {
			log.Debug(ms.TAG(), "receivedPackets: ", receivedPackets)
		}
		si.ReceivedPackets = int64(receivedPackets)

		var duplicatedDiscardedPackets uint32
		err = binary.Read(buffer, binary.LittleEndian, &duplicatedDiscardedPackets)
		checkError(err)
		if ms.logDebug {
			log.Debug(ms.TAG(), "duplicatedDiscardedPackets: ", duplicatedDiscardedPackets)
		}
		si.DuplicatedDiscardedPackets = int64(duplicatedDiscardedPackets)

		var noRoomDiscardedPackets uint32
		err = binary.Read(buffer, binary.LittleEndian, &noRoomDiscardedPackets)
		checkError(err)
		if ms.logDebug {
			log.Debug(ms.TAG(), "noRoomDiscardedPackets: ", noRoomDiscardedPackets)
		}
		si.NoRoomDiscardedPackets = int64(noRoomDiscardedPackets)

		var reassemblySkippedDiscardedPackets uint32
		err = binary.Read(buffer, binary.LittleEndian, &reassemblySkippedDiscardedPackets)
		checkError(err)
		if ms.logDebug {
			log.Debug(ms.TAG(), "reassemblySkippedDiscardedPackets: ", reassemblySkippedDiscardedPackets)
		}
		si.ReassemblySkippedDiscardedPackets = int64(reassemblySkippedDiscardedPackets)

		var estimatedRTT float32
		err = binary.Read(buffer, binary.LittleEndian, &estimatedRTT)
		checkError(err)
		if ms.logDebug {
			log.Debug(ms.TAG(), "estimatedRTT: ", estimatedRTT)
		}
		si.EstimatedRTT = estimatedRTT

		var unacknowledgedDataSize uint32
		err = binary.Read(buffer, binary.LittleEndian, &unacknowledgedDataSize)
		checkError(err)
		if ms.logDebug {
			log.Debug(ms.TAG(), "unacknowledgedDataSize: ", unacknowledgedDataSize)
		}
		si.UnacknowledgedDataSize = int64(unacknowledgedDataSize)

		var pendingDataSize uint32
		err = binary.Read(buffer, binary.LittleEndian, &pendingDataSize)
		checkError(err)
		if ms.logDebug {
			log.Debug(ms.TAG(), "pendingDataSize: ", pendingDataSize)
		}
		si.PendingDataSize = int64(pendingDataSize)

		var pendingPacketQueueSize uint32
		err = binary.Read(buffer, binary.LittleEndian, &pendingPacketQueueSize)
		checkError(err)
		if ms.logDebug {
			log.Debug(ms.TAG(), "pendingPacketQueueSize: ", pendingPacketQueueSize)
		}
		si.PendingPacketQueueSize = int64(pendingPacketQueueSize)

		var reliableSequencedPacketQueueSize uint32
		err = binary.Read(buffer, binary.LittleEndian, &reliableSequencedPacketQueueSize)
		checkError(err)
		if ms.logDebug {
			log.Debug(ms.TAG(), "reliableSequencedPacketQueueSize: ", reliableSequencedPacketQueueSize)
		}
		si.ReliableSequencedPacketQueueSize = int64(reliableSequencedPacketQueueSize)

		var reliableUnsequencedDataSize uint32
		err = binary.Read(buffer, binary.LittleEndian, &reliableUnsequencedDataSize)
		checkError(err)
		if ms.logDebug {
			log.Debug(ms.TAG(), "reliableUnsequencedDataSize: ", reliableUnsequencedDataSize)
		}
		si.ReliableUnsequencedDataSize = int64(reliableUnsequencedDataSize)

		var reliableUnsequencedPacketQueueSize uint32
		err = binary.Read(buffer, binary.LittleEndian, &reliableUnsequencedPacketQueueSize)
		checkError(err)
		if ms.logDebug {
			log.Debug(ms.TAG(), "reliableUnsequencedPacketQueueSize: ", reliableUnsequencedPacketQueueSize)
		}
		si.ReliableUnsequencedPacketQueueSize = int64(reliableUnsequencedPacketQueueSize)

		// sending measure back
		m := toMeasure(sensorIP, epi, si)
		if ms.logDebug {
			log.Debug(ms.TAG(), "Measure: ", m.String())
		}

		ms.mockets <- m

		flags, err := buffer.ReadByte()
		checkError(err)

		if uint8(flags) != uint8(mockets.MSF_OverallMessageStatistics) {
			if ms.logDebug {
				log.Debug(ms.TAG(), "Warning. Flags byte set to: ", flags)
			}
		}

		// MESSAGE INFO //
		mi := &mockets.MessageInfo{}
		var miMsgType uint16
		err = binary.Read(buffer, binary.LittleEndian, &miMsgType)
		checkError(err)
		if ms.logDebug {
			log.Debug(ms.TAG(), "MessageInfo Type: ", miMsgType)
		}
		mi.MsgType = int32(miMsgType)

		for mi.MsgType == int32(mockets.MSNT_Stats) {
			log.Debug("MessageInfo MsgType is MSNT_Stats (" + strconv.Itoa(int(mockets.MSNT_Stats)) + ")")

			var sentReliableSequencedMsgs uint32
			err = binary.Read(buffer, binary.LittleEndian, &sentReliableSequencedMsgs)
			checkError(err)
			if ms.logDebug {
				log.Debug(ms.TAG(), "sentReliableSequencedMsgs: ", sentReliableSequencedMsgs)
			}
			mi.SentReliableSequencedMsgs = int64(sentReliableSequencedMsgs)

			var sentReliableUnsequencedMsgs uint32
			err = binary.Read(buffer, binary.LittleEndian, &sentReliableUnsequencedMsgs)
			checkError(err)
			if ms.logDebug {
				log.Debug(ms.TAG(), "sentReliableUnsequencedMsgs: ", sentReliableUnsequencedMsgs)
			}
			mi.SentReliableUnsequencedMsgs = int64(sentReliableUnsequencedMsgs)

			var sentUnreliableSequencedMsgs uint32
			err = binary.Read(buffer, binary.LittleEndian, &sentUnreliableSequencedMsgs)
			checkError(err)
			if ms.logDebug {
				log.Debug(ms.TAG(), "sentUnreliableSequencedMsgs: ", sentUnreliableSequencedMsgs)
			}
			mi.SentUnreliableSequencedMsgs = int64(sentUnreliableSequencedMsgs)

			var sentUnreliableUnsequencedMsgs uint32
			err = binary.Read(buffer, binary.LittleEndian, &sentUnreliableUnsequencedMsgs)
			checkError(err)
			if ms.logDebug {
				log.Debug(ms.TAG(), "sentUnreliableUnsequencedMsgs: ", sentUnreliableUnsequencedMsgs)
			}
			mi.SentUnreliableUnsequencedMsgs = int64(sentUnreliableUnsequencedMsgs)

			var receivedReliableSequencedMsgs uint32
			err = binary.Read(buffer, binary.LittleEndian, &receivedReliableSequencedMsgs)
			checkError(err)
			if ms.logDebug {
				log.Debug(ms.TAG(), "receivedReliableSequencedMsgs: ", receivedReliableSequencedMsgs)
			}
			mi.ReceivedReliableSequencedMsgs = int64(receivedReliableSequencedMsgs)

			var receivedReliableUnsequencedMsgs uint32
			err = binary.Read(buffer, binary.LittleEndian, &receivedReliableUnsequencedMsgs)
			checkError(err)
			if ms.logDebug {
				log.Debug(ms.TAG(), "receivedReliableUnsequencedMsgs: ", receivedReliableUnsequencedMsgs)
			}
			mi.ReceivedReliableUnsequencedMsgs = int64(receivedReliableUnsequencedMsgs)

			var receivedUnreliableSequencedMsgs uint32
			err = binary.Read(buffer, binary.LittleEndian, &receivedUnreliableSequencedMsgs)
			checkError(err)
			if ms.logDebug {
				log.Debug(ms.TAG(), "receivedUnreliableSequencedMsgs: ", receivedUnreliableSequencedMsgs)
			}
			mi.ReceivedUnreliableSequencedMsgs = int64(receivedUnreliableSequencedMsgs)

			var receivedUnreliableUnsequencedMsgs uint32
			err = binary.Read(buffer, binary.LittleEndian, &receivedUnreliableUnsequencedMsgs)
			checkError(err)
			if ms.logDebug {
				log.Debug(ms.TAG(), "receivedUnreliableUnsequencedMsgs: ", receivedUnreliableUnsequencedMsgs)
			}
			mi.ReceivedUnreliableUnsequencedMsgs = int64(receivedUnreliableUnsequencedMsgs)

			var cancelledPackets uint32
			err = binary.Read(buffer, binary.LittleEndian, &cancelledPackets)
			checkError(err)
			if ms.logDebug {
				log.Debug(ms.TAG(), "cancelledPackets: ", cancelledPackets)
			}
			mi.CancelledPackets = int64(cancelledPackets)

			///////////////////////////////
			//TODO send mi to channel before reading next msg type
			///////////////////////////////

			var miMsgType uint16
			err = binary.Read(buffer, binary.LittleEndian, &miMsgType)
			mi.MsgType = int32(miMsgType)
		}
	}
}

func toMeasure(sensorIP net.IP, epi *mockets.EndPointsInfo, si *mockets.StatisticsInfo) *measure.Measure {

	m := &measure.Measure{
		Subject:   measure.Subject_mockets,
		Strings:   make(map[string]string),
		Integers:  make(map[string]int64),
		Doubles:   make(map[string]float64),
		Timestamp: &timestamp.Timestamp{Seconds: time.Now().Unix()}, //TODO:time override

	}

	m.GetStrings()[subjmockets.Str_pid.String()] = strconv.Itoa(int(epi.PID))
	m.GetStrings()[subjmockets.Str_identifier.String()] = epi.Identifier
	m.GetStrings()[subjmockets.Str_sensor_ip.String()] = sensorIP.String() // duplicated
	m.GetStrings()[subjmockets.Str_src_ip.String()] = util.UInt32ToIP(uint32(epi.LocalAddr)).String()
	m.GetStrings()[subjmockets.Str_src_port.String()] = strconv.Itoa(int(epi.LocalPort))
	m.GetStrings()[subjmockets.Str_dest_ip.String()] = util.UInt32ToIP(uint32(epi.RemoteAddr)).String()
	m.GetStrings()[subjmockets.Str_dest_port.String()] = strconv.Itoa(int(epi.RemotePort))

	m.GetIntegers()[subjmockets.Int_last_contact_time.String()] = si.LastContactTime
	m.GetIntegers()[subjmockets.Int_sent_bytes.String()] = si.SentBytes
	m.GetIntegers()[subjmockets.Int_sent_packets.String()] = si.SentPackets
	m.GetIntegers()[subjmockets.Int_retransmits.String()] = si.Retransmits
	m.GetIntegers()[subjmockets.Int_received_bytes.String()] = si.ReceivedBytes
	m.GetIntegers()[subjmockets.Int_received_packets.String()] = si.ReceivedPackets
	m.GetIntegers()[subjmockets.Int_duplicated_discarded_packets.String()] = si.DuplicatedDiscardedPackets
	m.GetIntegers()[subjmockets.Int_no_room_discarded_packets.String()] = si.NoRoomDiscardedPackets
	m.GetIntegers()[subjmockets.Int_reassembly_skipped_discarded_packets.String()] = si.ReassemblySkippedDiscardedPackets
	//RTT
	m.GetDoubles()[subjmockets.Double_estimated_rtt.String()] = float64(si.EstimatedRTT)
	m.GetIntegers()[subjmockets.Int_unacknowledged_data_size.String()] = si.UnacknowledgedDataSize
	m.GetIntegers()[subjmockets.Int_unacknowledged_queue_size.String()] = si.UnacknowledgedQueueSize
	m.GetIntegers()[subjmockets.Int_pending_data_size.String()] = si.PendingDataSize
	m.GetIntegers()[subjmockets.Int_pending_packet_queue_size.String()] = si.PendingPacketQueueSize
	m.GetIntegers()[subjmockets.Int_reliable_sequenced_data_size.String()] = si.ReliableSequencedDataSize
	m.GetIntegers()[subjmockets.Int_reliable_sequenced_packet_queue_size.String()] = si.ReliableSequencedPacketQueueSize
	m.GetIntegers()[subjmockets.Int_reliable_unsequenced_data_size.String()] = si.ReliableUnsequencedDataSize
	m.GetIntegers()[subjmockets.Int_reliable_unsequenced_packet_queue_size.String()] = si.ReliableUnsequencedPacketQueueSize

	return m
}

func (ms *MocketsSensor) TAG() string {
	return ms.tag + " [" + util.GetGIDString() + "] "
}
