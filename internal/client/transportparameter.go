package client

import (
	"github.com/lucas-clemente/quic-go"
)

func TransportParameterRequest() {
	requestFilesWithinOneConnection(&quic.Config{}, getDefaultTLSConf())
}
