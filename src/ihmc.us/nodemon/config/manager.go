package config

// handles NodeMon configuration
type Manager struct {
	NATSHost             string // hostname or ip address of the gNATSd server
	NATSPort             uint16 // port used by the gNATSs server
	NATSLogDebug         bool   // enables log debug for NATS Server
	NetSensorPort        uint16 // port used by NetSensor
	NetSensorLogDebug    bool   // enables debug log for NetSensor
	MocketsSensorPort    uint16 // port used by MocketsSensor
	MocketsLogDebug      bool   // enables debug log for MocketsSensor
	DisServiceSensorPort uint16 // port used by DisServiceSensor
	DisServiceLogDebug   bool   // enables debug log for DisServiceSensor
	HTTPServerPort       uint16 // port used by the internal HTTP Server
	HTTPServerLogDebug   bool   // enables debug log for HTTP Server
	File                 string // config file
	LogLevel             string // log level. Possible values debug, info, warn, error, fatal, panic
}
