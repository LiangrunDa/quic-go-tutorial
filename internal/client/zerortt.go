package client

import (
	"crypto/tls"

	"github.com/lucas-clemente/quic-go"
	"github.com/lucas-clemente/quic-go/http3"
)

func ZeroRTTRequest() {
	quicConf := &quic.Config{}

	tlsConfig := getDefaultTLSConf()
	tlsConfig.ClientSessionCache = tls.NewLRUClientSessionCache(10)
	roundTripper := &http3.RoundTripper{
		TLSClientConfig: tlsConfig,
		QuicConfig:      quicConf,
	}
	requestFirstFile(roundTripper, quicConf, tlsConfig)
	requestRemainingFiles(roundTripper, quicConf, tlsConfig)

}
