package client

import (
	"crypto/tls"

	"github.com/lucas-clemente/quic-go"
)

func Chacha20Request() {
	tlsConf := getDefaultTLSConf()
	tlsConf.CipherSuites = []uint16{tls.TLS_CHACHA20_POLY1305_SHA256}
	requestFilesWithinOneConnection(&quic.Config{}, tlsConf)

}
