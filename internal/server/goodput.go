package server

import (
	"github.com/lucas-clemente/quic-go"
	"github.com/sirupsen/logrus"
)

func GoodputListen() {
	server := getWWWServerWithQlogOption(false, &quic.Config{}, getDefaultTLSConfig())

	err := server.ListenAndServe()
	if err != nil {
		logrus.Error(err)
	}
}
