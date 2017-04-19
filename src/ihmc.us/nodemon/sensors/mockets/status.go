package mockets

const (
	DefaultPort      = 1400
	UDPPacketMaxSize = 1500 //in bytes
)

//enum for NoticeType
type NoticeType uint8

const (
	MSNT_Undefined             NoticeType = iota
	MSNT_ConnectionFailed
	MSNT_ConnectionEstablished
	MSNT_ConnectionReceived
	MSNT_Stats
	MSNT_Disconnected
	MSNT_ConnectionRestored
)

//enum for StatusFlag
type StatusFlags uint8

const (
	MSF_Undefined                StatusFlags = iota
	MSF_End
	MSF_OverallMessageStatistics
	MSF_PerTypeMessageStatistics
)

type EndPointsInfo struct {
	PID        int64
	Identifier string
	LocalAddr  int64
	LocalPort  int64
	RemoteAddr int64
	RemotePort int64
}

type StatisticsInfo struct {
	LastContactTime                    int64
	SentBytes                          int64
	SentPackets                        int64
	Retransmits                        int64
	ReceivedBytes                      int64
	ReceivedPackets                    int64
	DuplicatedDiscardedPackets         int64
	NoRoomDiscardedPackets             int64
	ReassemblySkippedDiscardedPackets  int64
	EstimatedRTT                       float32
	UnacknowledgedDataSize             int64
	UnacknowledgedQueueSize            int64 // Update missing
	PendingDataSize                    int64
	PendingPacketQueueSize             int64
	ReliableSequencedDataSize          int64
	ReliableSequencedPacketQueueSize   int64
	ReliableUnsequencedDataSize        int64
	ReliableUnsequencedPacketQueueSize int64
}
type MessageInfo struct {
	MsgType                           int32
	SentReliableSequencedMsgs         int64
	SentReliableUnsequencedMsgs       int64
	SentUnreliableSequencedMsgs       int64
	SentUnreliableUnsequencedMsgs     int64
	ReceivedReliableSequencedMsgs     int64
	ReceivedReliableUnsequencedMsgs   int64
	ReceivedUnreliableSequencedMsgs   int64
	ReceivedUnreliableUnsequencedMsgs int64
	CancelledPackets                  int64
}
