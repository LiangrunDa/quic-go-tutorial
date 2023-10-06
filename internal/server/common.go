package server

import (
	"crypto/tls"
	"liangrun/internal/qlogtracer"
	"liangrun/utils"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/lucas-clemente/quic-go"
	"github.com/lucas-clemente/quic-go/http3"
)

func setupWWWHandler() *gin.Engine {
	www := utils.ServerEnvMap["www"]
	engine := gin.New()
	engine.StaticFS("/", http.Dir(www))
	return engine
}

func getWWWServerWithQlogOption(enableQlog bool, quicConf *quic.Config, tlsConf *tls.Config) http3.Server {
	handler := setupWWWHandler()
	if enableQlog{
		qlogtracer.AddTracer(true, quicConf)
	}
	addr := utils.ServerEnvMap["ip"] + ":" + utils.ServerEnvMap["port"]

	certsDir := utils.ServerEnvMap["certs"]
	keyFile := certsDir + "priv.key"
	certFile := certsDir + "cert.pem"
	certs := make([]tls.Certificate, 1)
	certs[0], _ = tls.LoadX509KeyPair(certFile, keyFile)
	server := http3.Server{
		Handler:    handler,
		Addr:       addr,
		QuicConfig: quicConf,
		TLSConfig:  tlsConf,
	}
	return server
}

func getWWWServer(quicConf *quic.Config, tlsConf *tls.Config) http3.Server {
	return getWWWServerWithQlogOption(true, quicConf, tlsConf)
}

func getDefaultTLSConfig() *tls.Config {
	certsDir := utils.ServerEnvMap["certs"]
	keyFile := certsDir + "priv.key"
	certFile := certsDir + "cert.pem"
	certs := make([]tls.Certificate, 1)
	certs[0], _ = tls.LoadX509KeyPair(certFile, keyFile)
	keyLog := utils.ServerEnvMap["sslKeyLogFile"]
	f, _ := os.OpenFile(keyLog, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	return &tls.Config{
		InsecureSkipVerify: true,
		KeyLogWriter:       f,
		Certificates:       certs,
	}
}
