package server

import (
	"net"
	"time"

	"github.com/lucas-clemente/quic-go"
	"github.com/sirupsen/logrus"
)

func ResumptionListen() {

	tlsConfig := getDefaultTLSConfig()
	server := getWWWServer(&quic.Config{
		MaxTokenAge: time.Hour * 1,
		Allow0RTT: func(a net.Addr) bool {
			return true
		},
	}, tlsConfig)

	err := server.ListenAndServe()
	if err != nil {
		logrus.Error(err)
	}
}
