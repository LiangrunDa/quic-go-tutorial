package optimizeclient

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/lucas-clemente/quic-go"
	"github.com/lucas-clemente/quic-go/http3"
)

type DefaultClient struct {
}

func (client *DefaultClient) getTconf() *tls.Config {
	keyLog := ClientEnvs["sslKeyLogFile"]
	f, err := os.OpenFile(keyLog, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	return &tls.Config{
		InsecureSkipVerify: true,
		KeyLogWriter:       f,
	}
}

func (client *DefaultClient) getQconf() *quic.Config {
	return &quic.Config{}
}

func (client *DefaultClient) requestFile() {
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
