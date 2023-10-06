package optimizeclient

import (
	"crypto/tls"
	"fmt"
	"io"
	"liangrun/utils"
	"net/http"
	"os"
	"strings"

	"github.com/lucas-clemente/quic-go"
	"github.com/lucas-clemente/quic-go/http3"
)

type QconfOptimizedClient struct {
}

func (client *QconfOptimizedClient) getTconf() *tls.Config {
	keyLog := ClientEnvs["sslKeyLogFile"]
	f, err := os.OpenFile(keyLog, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	return &tls.Config{
		InsecureSkipVerify: true,
		KeyLogWriter:       f,
		CipherSuites:       []uint16{tls.TLS_CHACHA20_POLY1305_SHA256},
	}
}

func (client *QconfOptimizedClient) getQconf() *quic.Config {
	return &quic.Config{
		InitialStreamReceiveWindow:     uint64(utils.Humanize("1GB")),
		InitialConnectionReceiveWindow: uint64(utils.Humanize("1GB")),
		MaxStreamReceiveWindow:         uint64(utils.Humanize("1GB")),
		MaxConnectionReceiveWindow:     uint64(utils.Humanize("1GB")),
		AllowConnectionWindowIncrease: func(conn quic.Connection, delta uint64) bool {
			return true
		},
		MaxIncomingStreams:    1000,
		MaxIncomingUniStreams: 1000,
	}
}

func (client *QconfOptimizedClient) requestFile() {
	requests := strings.Split(ClientEnvs["requests"], " ")

	roundTripper := &http3.RoundTripper{
		TLSClientConfig: client.getTconf(),
		QuicConfig:      client.getQconf(),
	}
	defer roundTripper.Close()

	hclient := &http.Client{
		Transport: roundTripper,
	}

	for _, addr := range requests {
		rsp, err := hclient.Get(addr)
		if err != nil {
			fmt.Print(err)
		}
		parts := strings.Split(addr, "/")
		filePath := ClientEnvs["downloads"] + parts[len(parts)-1]
		f, _ := os.Create(filePath)
		io.Copy(f, rsp.Body)
	}
}
