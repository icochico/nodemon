package cmd

import (
	log "github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"ihmc.us/nodemon/http"
	"ihmc.us/nodemon/measure"
	"ihmc.us/nodemon/util"
	"ihmc.us/nodemon/broker"
	"ihmc.us/nodemon/sensors"
)

const (
	CPUChanSize        = 1000
	MemoryChanSize     = 1000
	TrafficChanMaxSize = 1000
	HostChanSize       = 1000
	NetworkChanSize    = 1000
	MocketsChanSize    = 1000
)

func init() {
	RootCmd.AddCommand(runCmd)
}

type NodeMon struct {
	cpu     chan *measure.Measure // channel for CPU related measures
	memory  chan *measure.Measure // channel for memory related  measures
	traffic chan *measure.Measure // channel for traffic measures
	host    chan *measure.Measure // channel for host related information measures
	network chan *measure.Measure // channel for network related information measures
	mockets chan *measure.Measure // channel for Mockets related information measures
	stats   chan int              // channel for NodeMon related statistics
	quit    chan bool
}

func NewNodeMon() *NodeMon {
	return &NodeMon{
		cpu:     make(chan *measure.Measure, CPUChanSize),
		memory:  make(chan *measure.Measure, MemoryChanSize),
		traffic: make(chan *measure.Measure, TrafficChanMaxSize),
		host:    make(chan *measure.Measure, HostChanSize),
		network: make(chan *measure.Measure, NetworkChanSize),
		mockets: make(chan *measure.Measure, MocketsChanSize),
		quit:    make(chan bool),
	}
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the NodeMon",
	Long:  "Run the NodeMon with the specified options",
	Run: func(cmd *cobra.Command, args []string) {

		n := NewNodeMon()

		//config log
		//log.SetFormatter(&log.JSONFormatter{})
		//log.SetLevel(log.InfoLevel)
		logLevel, err := log.ParseLevel(cfg.LogLevel)
		if err != nil {
			logLevel = log.DebugLevel
		}
		log.SetLevel(logLevel)

		//initialize HTTP Server
		httpServer := http.NewServer(cfg.HTTPServerPort, cfg.HTTPServerLogDebug)
		httpServer.Start()

		//initialize NATS
		nats := broker.NewNATS(cfg.NATSPort, n.traffic, n.host, n.network, n.mockets, cfg.NATSLogDebug)
		nats.Start() //decoupled because of reconnect attempts

		//initialize sensors

		// SystemSensor
		ss := sensors.NewSystemSensor(n.cpu, n.memory, false)
		ss.Start()

		// NetSensor
		ns := sensors.NewNetSensor(cfg.NetSensorPort, n.traffic, n.host, n.network, cfg.NetSensorLogDebug)
		ns.Start()

		//MocketsSensor
		ms := sensors.NewMocketSensor(cfg.MocketsSensorPort, n.mockets, cfg.NetSensorLogDebug)
		ms.Start()
		log.Info(getTAG() + "Waiting for packets...")

		for {
			//time.Sleep(10 * time.Millisecond)
			select {
			case <-n.quit:
				log.Info(getTAG() + "Quitting!")
				os.Exit(0)
			}
		}

	},
}

func getTAG() string {
	return TAG + " [" + util.GetGIDString() + "] "
}
