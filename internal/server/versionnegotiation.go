package server

import (
	"github.com/lucas-clemente/quic-go"
	"github.com/sirupsen/logrus"
)

func VersionNegotiationListen() {
	quicConf := &quic.Config{}
	quicConf.Versions = append(quicConf.Versions, 0x1)
	server := getWWWServer(quicConf, getDefaultTLSConfig())

	err := server.ListenAndServe()
	if err != nil {
		logrus.Error(err)
	}
}
