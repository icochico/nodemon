package sensors

import (
	log "github.com/Sirupsen/logrus"
	"net"
	"strconv"
)

type Sensor interface {
	Start()
	Stop()
}

func createUDPConn(port int, tag string) (conn *net.UDPConn, err error) {

	addr := ":" + strconv.Itoa(port)
	udpAddr, err := net.ResolveUDPAddr("udp4", addr)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	conn, err = net.ListenUDP("udp4", udpAddr)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	log.Info(tag, "Listening to: ", udpAddr.String())

	return conn, nil
}

func checkError(err error) {
	if err != nil {
		log.Error("Error: ", err.Error())
	}
}
