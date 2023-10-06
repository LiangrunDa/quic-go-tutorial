package server

import (
	"github.com/lucas-clemente/quic-go"
	"github.com/sirupsen/logrus"
)

func TransportParameterListen() {

	quicConf := &quic.Config{}
	quicConf.MaxIncomingStreams = 10
	server := getWWWServer(quicConf, getDefaultTLSConfig())

	err := server.ListenAndServe()
	if err != nil {
		logrus.Error(err)
	}

}
