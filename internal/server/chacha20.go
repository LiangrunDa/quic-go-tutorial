package server

import (
	"crypto/tls"

	"github.com/lucas-clemente/quic-go"
	"github.com/sirupsen/logrus"
)

func Chacha20Listen() {
	tlsConf := getDefaultTLSConfig()
	tlsConf.CipherSuites = []uint16{tls.TLS_CHACHA20_POLY1305_SHA256}
	server := getWWWServer(&quic.Config{}, tlsConf)

	err := server.ListenAndServe()
	if err != nil {
		logrus.Error(err)
	}
}
