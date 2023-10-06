package client

import (
	"github.com/lucas-clemente/quic-go"
)

func MultiHandshakeRequest() {
	requestFilesWithinMultipleConnection(&quic.Config{}, getDefaultTLSConf())

}
