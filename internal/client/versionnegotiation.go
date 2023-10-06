package client

import (
	"github.com/lucas-clemente/quic-go"
)

func VersionNegotiationRequest() {
	qconf := &quic.Config{}
	qconf.Versions = append(qconf.Versions, 0xff00001d)

	requestFilesWithinOneConnection(qconf, getDefaultTLSConf())
}
