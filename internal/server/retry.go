package server

import (
	"net"

	"github.com/lucas-clemente/quic-go"
	"github.com/sirupsen/logrus"
)

func RetryListen() {
	quicConf := &quic.Config{}
	quicConf.RequireAddressValidation = func(net.Addr) bool {
		return true
	}

	server := getWWWServer(quicConf, getDefaultTLSConfig())

	err := server.ListenAndServe()
	if err != nil {
		logrus.Error(err)
	}
}
