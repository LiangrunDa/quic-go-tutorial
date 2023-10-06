package optimizeserver

import (
	"crypto/tls"
	"net/http"
	"os"

	"github.com/lucas-clemente/quic-go"
	"github.com/lucas-clemente/quic-go/http3"
)

type DefaultServer struct {
}

func (server *DefaultServer) getTconf() *tls.Config {
	certsDir := ServerEnvs["certs"]
	keyFile := certsDir + "priv.key"
	certFile := certsDir + "cert.pem"
	certs := make([]tls.Certificate, 1)
	certs[0], _ = tls.LoadX509KeyPair(certFile, keyFile)
	keyLog := ServerEnvs["sslKeyLogFile"]
	f, _ := os.OpenFile(keyLog, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	return &tls.Config{
		InsecureSkipVerify: true,
		KeyLogWriter:       f,
		Certificates:       certs,
	}
}

func (server *DefaultServer) getQconf() *quic.Config {
	return &quic.Config{}
}

func (server *DefaultServer) getWWWHandler() *http.ServeMux {
	www := ServerEnvs["www"]
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(www)))
	return mux

}

func (server *DefaultServer) listen() {
	addr := ServerEnvs["ip"] + ":" + ServerEnvs["port"]

	http3server := http3.Server{
		Handler:    server.getWWWHandler(),
		Addr:       addr,
		QuicConfig: server.getQconf(),
		TLSConfig:  server.getTconf(),
	}

	http3server.ListenAndServe()
}
