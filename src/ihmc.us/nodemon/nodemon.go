package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/CrowdSurge/banner"
	"ihmc.us/nodemon/measure"
	"os"
	"ihmc.us/nodemon/util"
	"ihmc.us/nodemon/sensors"
	"ihmc.us/nodemon/broker"
	"ihmc.us/nodemon/sensors/netsensor"
	"ihmc.us/nodemon/sensors/mockets"
	"fmt"
)

const (
	TAG                   = "NodeMon"
	TrafficChannelMaxSize = 1000
	HostChannelMaxSize    = 1000
	NetworkChannelMaxSize = 1000
	MocketsChannelMaxSize = 1000
)

//variables for versioning
var (
	Version   string
	BuildTime string
	GitHash   string
)

type NodeMon struct {
	traffic chan *measure.Measure // channel for traffic measures
	host    chan *measure.Measure // channel for host related information measures
	network chan *measure.Measure // channel for network related information measures
	mockets chan *measure.Measure // channel for Mockets related information measures
	stats   chan int              // channel for NodeMon related statistics
	quit    chan bool
}

func NewNodeMon() *NodeMon {
	return &NodeMon{
		traffic: make(chan *measure.Measure, TrafficChannelMaxSize),
		host:    make(chan *measure.Measure, HostChannelMaxSize),
		network: make(chan *measure.Measure, NetworkChannelMaxSize),
		mockets: make(chan *measure.Measure, MocketsChannelMaxSize),
		quit:    make(chan bool),
	}
}

func (n *NodeMon) Start() {
	//config log
	//log.SetFormatter(&log.JSONFormatter{})
	//log.SetLevel(log.InfoLevel)
	fmt.Println(banner.PrintS("nodemon"))
	log.SetLevel(log.DebugLevel)
	log.Info(getTAG() + "NodeMon started!")
	log.Info(getTAG() + "Version: " + Version)
	log.Info(getTAG() + "Build Time: " + BuildTime)
	log.Info(getTAG() + "Git Hash: " + GitHash)

	//initialize NATS
	nats, err := broker.NewNATS(broker.NATSPort, n.traffic, n.host, n.network, n.mockets, true)
	if err != nil {
		log.Warn(getTAG() + err.Error())
	} else {
		nats.Start()
	}

	//initialize sensors
	// NetSensor
	ns := sensors.NewNetSensor(netsensor.DefaultPort, n.traffic, n.host, n.network, false)
	ns.Start()

	//MocketsSensor
	ms := sensors.NewMocketSensor(mockets.DefaultPort, n.mockets, false)
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
}

func (n *NodeMon) Stop() {
	if n == nil {
		return
	}
	log.Info(getTAG() + "Stopping NodeMon...")
	n.quit <- true
}

func getTAG() string {
	return TAG + " [" + util.GetGIDString() + "] "
}

func main() {
	n := NewNodeMon()
	n.Start()
}
