package client

import (
	"github.com/lucas-clemente/quic-go"
)

func TransferRequest() {
	requestFilesWithinOneConnection(&quic.Config{}, getDefaultTLSConf())

}
