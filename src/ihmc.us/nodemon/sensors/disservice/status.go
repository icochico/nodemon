package disservice

//Constants fort DisServiceStatus
const (
	DIS_SERVICE_STATUS_HEADER_BYTE int = 0x00D1
	DefaultPort                    int = 2400
	UDPPacketMaxSize	       int = 2048
)

//enum for StatusNoticeType
type StatusNoticeType int16

const (
	DSSNT_Undefined          StatusNoticeType = iota
	DSSNT_ClientConnected
	DSSNT_ClientDisconnected
	DSSNT_SummaryStats
	DSSNT_DetailedStats
	DSSNT_TopologyStatus
)

//enum for StatusFlags
type StatusFlags int16

const (
	DSSF_Undefined              StatusFlags = iota
	DSSF_End
	DSSF_OverallStats
	DSSF_PerClientGroupTagStats
	DSSF_DuplicateTrafficInfo
)

type BasicStatisticsInfo struct {
	PeerId                                 string
	DataMessageReceived                    int64
	DataBytesReceived                      int64
	DataFragmentsReceived                  int64
	DataFragmentBytesReceived              int64
	MissingFragmentRequestMessagesSent     int64
	MissingFragmentRequestBytesSent        int64
	MissingFragmentRequestMessagesReceived int64
	MissingFragmentRequestBytesReceived    int64
	DataCacheQueryMessagesSent             int64
	DataCacheQueryBytesSent                int64
	DataCacheQueryMessagesReceived         int64
	DataCacheQueryBytesReceived            int64
	TopologyStateMessagesSent              int64
	TopologyStateBytesSent                 int64
	TopologyStateMessagesReceived          int64
	TopologyStateBytesReceived             int64
	KeepAliveMessagesSent                  int64
	KeepAliveMessagesReceived              int64
	QueryMessagesSent                      int64
	QueryMessagesReceived                  int64
	QueryHitsMessagesSent                  int64
	QueryHitsMessagesReceived              int64
}

type StatsInfo struct {
	PeerId                      string
	ClientMessagesPushed        int64
	ClientBytesPushed           int64
	ClientMessagesMadeAvailable int64
	ClientBytesMadeAvailable    int64
	FragmentsPushed             int64
	FragmentBytesPushed         int64
	OnDemandFragmentsSent       int64
	OnDemandFragmentBytesSent   int64
}

type BasicStatisticsInfoByPeer struct {
	PeerId                                 string
	RemotePeerId                           string
	DataMessageReceived                    int64
	DataBytesReceived                      int64
	DataFragmentReceived                   int64
	DataFragmentBytesReceived              int64
	MissingFragmentRequestMessagesReceived int64
	MissingFragmentRequestBytesReceived    int64
	KeepAliveMessagesReceived              int64
}

type DuplicateTrafficInfo struct {
	PeerId                    string
	TargetedDuplicateTraffic  int64
	OverheardDuplicateTraffic int64
}

type ClientGroupTagStatsInfoHeader struct {
	PeerId          string
	ClientId        int32
	Tag             int32
	GroupNameLength int32
}
