package client

import (
	"bytes"
	"crypto/tls"
	"io"
	"liangrun/internal/qlogtracer"
	"liangrun/utils"
	"net/http"
	"os"
	"strings"

	"github.com/lucas-clemente/quic-go"
	"github.com/lucas-clemente/quic-go/http3"
	"github.com/sirupsen/logrus"
)

func requestOneFile(hclient *http.Client, addr string) {
	logrus.Infof("request %v", addr)
	rsp, err := hclient.Get(addr)
	if err != nil {
		logrus.Error(err)
	}
	parts := strings.Split(addr, "/")
	filePath := utils.ClientEnvMap["downloads"] + parts[len(parts)-1]

	f, createErr := os.Create(filePath)
	if createErr != nil{
		logrus.Error(createErr)
	}
	_, err = io.Copy(f, rsp.Body)
	if err != nil {
		logrus.Error(err)
	}

	logrus.Infof("download to %v", filePath)
}

func requestFirstFile(roundTripper *http3.RoundTripper, qconf *quic.Config, tlsConf *tls.Config) {
	qlogtracer.AddTracer(false, qconf)
	requests := strings.Split(utils.ClientEnvMap["requests"], " ")

	defer roundTripper.Close()

	hclient := &http.Client{
		Transport: roundTripper,
	}

	addr := requests[0]
	rsp, err := hclient.Get(addr)
	if err != nil {
		logrus.Error(err)
	}

	body := &bytes.Buffer{}
	_, err = io.Copy(body, rsp.Body)
	if err != nil {
		logrus.Error(err)
	}
	parts := strings.Split(addr, "/")
	filePath := utils.ClientEnvMap["downloads"] + parts[len(parts)-1]
	utils.WriteFile(filePath, body.Bytes())

}

func requestRemainingFiles(roundTripper *http3.RoundTripper, qconf *quic.Config, tlsConf *tls.Config) {
	qlogtracer.AddTracer(false, qconf)
	requests := strings.Split(utils.ClientEnvMap["requests"], " ")

	defer roundTripper.Close()
	hclient := &http.Client{
		Transport: roundTripper,
	}

	for _, addr := range requests[1:] {
		requestOneFile(hclient, addr)
	}

}

func requestFilesWithinOneConnectionWithQlogOption(enableQlog bool, qconf *quic.Config, tlsConf *tls.Config) {

	if enableQlog {
		qlogtracer.AddTracer(false, qconf)
	}

	requests := strings.Split(utils.ClientEnvMap["requests"], " ")

	roundTripper := &http3.RoundTripper{
		TLSClientConfig: tlsConf,
		QuicConfig:      qconf,
	}
	defer roundTripper.Close()
	hclient := &http.Client{
		Transport: roundTripper,
	}
	for _, addr := range requests {
		requestOneFile(hclient, addr)
	}
}

func requestFilesWithinOneConnection(qconf *quic.Config, tlsConf *tls.Config) {

	requestFilesWithinOneConnectionWithQlogOption(true, qconf, tlsConf)
}

func requestFilesWithinMultipleConnection(qconf *quic.Config, tlsConf *tls.Config) {
	qlogtracer.AddTracer(false, qconf)
	requests := strings.Split(utils.ClientEnvMap["requests"], " ")

	for _, addr := range requests {
		roundTripper := &http3.RoundTripper{
			TLSClientConfig: tlsConf,
			QuicConfig:      qconf,
		}
		defer roundTripper.Close()
		hclient := &http.Client{
			Transport: roundTripper,
		}
		requestOneFile(hclient, addr)
	}
}

func getDefaultTLSConf() *tls.Config {
	keyLog := utils.ClientEnvMap["sslKeyLogFile"]
	f, err := os.OpenFile(keyLog, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	return &tls.Config{
		InsecureSkipVerify: true,
		KeyLogWriter:       f,
	}
}
