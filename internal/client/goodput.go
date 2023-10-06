package client

import (
	"github.com/lucas-clemente/quic-go"
)

func GoodputRequest() {
	requestFilesWithinOneConnectionWithQlogOption(false, &quic.Config{}, getDefaultTLSConf())
}