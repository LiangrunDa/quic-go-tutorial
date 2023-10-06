package client

import (
	"github.com/lucas-clemente/quic-go"
)

func HandshakeRequest() {
	requestFilesWithinOneConnection(&quic.Config{}, getDefaultTLSConf())
}
