package http

import (
	"net/http"
	"github.com/labstack/echo"
	"io/ioutil"
	"strconv"
)

const (
	DefaultPort = 1323
)

type Server struct {
	e        *echo.Echo
	port     uint16
	logDebug bool
}

func NewServer(port uint16, logDebug bool) (*Server) {

	return &Server{
		e:        echo.New(),
		port:     port,
		logDebug: logDebug,
	}
}

func (s*Server) Start() {
	go func() {
		s.addRoutes()
		if !s.logDebug {
			s.e.Logger.SetOutput(ioutil.Discard) //disabled logger
		}
		s.e.Logger.Fatal(s.e.Start(":" + strconv.Itoa(DefaultPort)))
	}()
}

func (s*Server) addRoutes() {

	s.e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "NodeMon Web Interface")
	})
}
