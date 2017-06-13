package broker

import (
	log "github.com/Sirupsen/logrus"
	"ihmc.us/nodemon/measure"
	"github.com/nats-io/go-nats"
	"strconv"
	"github.com/golang/protobuf/proto"
	"ihmc.us/nodemon/util"
	"time"
)

const (
	TAG           = "NATS"
	NATSSchema    = "nats://"
	DefaultPort   = 4222
	DefaultHost   = "localhost"
	ReconnectTime = 3000 //msec
)

type NATS struct {
	conn     *nats.Conn
	port     uint16
	traffic  <-chan *measure.Measure
	host     <-chan *measure.Measure
	network  <-chan *measure.Measure
	mockets  <-chan *measure.Measure
	logDebug bool
	quit     chan bool
}

func NewNATS(port uint16, traffic, host, network, mockets <-chan *measure.Measure, logDebug bool) (*NATS) {
	return &NATS{
		port:     port,
		traffic:  traffic,
		host:     host,
		network:  network,
		mockets:  mockets,
		logDebug: logDebug,
		quit:     make(chan bool),
	}
}

func (n *NATS) Start() {

	go func() {
		url := NATSSchema + DefaultHost + ":" + strconv.Itoa(int(n.port))
		log.Info(getTAG() + "Attempting connection to:" + url)

		conn, err := nats.Connect(url)
		for err != nil {
			log.Warn(getTAG() + err.Error())
			time.Sleep(ReconnectTime * time.Millisecond)
			conn, err = nats.Connect(url)
		}

		log.Info(getTAG() + "Connection successful to: " + conn.ConnectedUrl())
		// assign conn object
		n.conn = conn
		go n.handlePub()
	}()

}

func (n *NATS) handlePub() {

	countTraffic, countHost, countNetwork, countMockets := 0, 0, 0, 0
	log.Info(getTAG() + "Waiting to publish ...")
	for {
		var m *measure.Measure
		select {
		case m = <-n.traffic:
			countTraffic++
			go n.publish(m)
		case m = <-n.host:
			countHost++
			go n.publish(m)
		case m = <-n.network:
			countNetwork++
			go n.publish(m)
		case m = <-n.mockets:
			countMockets++
			go n.publish(m)
		case <-n.quit:
			log.Info(getTAG() + "Quitting!")
			n.conn.Close() //closing NATS channel
			return         //exit goroutine
		}
		//default:
		//	sleep := 50 * time.Millisecond
		//	log.Debug("Nothing to publish, sleeping: " + sleep.String())
		//	time.Sleep(sleep)
		//}
	}
}

func (n *NATS) publish(m *measure.Measure) {
	start := time.Now()

	data, err := proto.Marshal(m)
	if err != nil {
		log.Warn(getTAG() + err.Error())
		return
	}

	err = n.conn.Publish(m.GetSubject().String(), data)
	if err != nil {
		log.Warn(getTAG() + err.Error())
		return
	}

	end := time.Since(start)

	if n.logDebug {
		log.Debug(getTAG() + "publish() subj: " + m.GetSubject().String() + " size: " +
			strconv.Itoa(len(data)) + " bytes, time: " + end.String())
	}
}

func (n *NATS) Stop() {
	go func() {
		n.quit <- true
	}()
}

func getTAG() string {
	return TAG + " [" + util.GetGIDString() + "] "
}
