package client

import (
	"github.com/lucas-clemente/quic-go"
)

func RetryRequest() {
	requestFilesWithinOneConnection(&quic.Config{}, getDefaultTLSConf())
}
