package server

import (
	"github.com/lucas-clemente/quic-go"
	"github.com/sirupsen/logrus"
)

func TransferListen() {
	server := getWWWServer(&quic.Config{}, getDefaultTLSConfig())

	err := server.ListenAndServe()
	if err != nil {
		logrus.Error(err)
	}
}
