package optimizeserver

import (
	"crypto/tls"
	"liangrun/utils"
	"net/http"
	"os"

	"github.com/lucas-clemente/quic-go"
	"github.com/lucas-clemente/quic-go/http3"
)

type QconfOptimizedServer struct {
}

func (server *QconfOptimizedServer) getTconf() *tls.Config {
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

func (server *QconfOptimizedServer) getQconf() *quic.Config {
	return &quic.Config{
		InitialStreamReceiveWindow:     uint64(utils.Humanize("50MB")),
		InitialConnectionReceiveWindow: uint64(utils.Humanize("50MB")),
		MaxStreamReceiveWindow:         uint64(utils.Humanize("1GB")),
		MaxConnectionReceiveWindow:     uint64(utils.Humanize("1GB")),
		AllowConnectionWindowIncrease: func(conn quic.Connection, delta uint64) bool {
			return true
		},
		MaxIncomingStreams:    1000,
		MaxIncomingUniStreams: 1000,
	}
}

func (server *QconfOptimizedServer) getWWWHandler() *http.ServeMux {
	www := ServerEnvs["www"]
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(www)))
	return mux

}

func (server *QconfOptimizedServer) listen() {
	addr := ServerEnvs["ip"] + ":" + ServerEnvs["port"]

	http3server := http3.Server{
		Handler:    server.getWWWHandler(),
		Addr:       addr,
		QuicConfig: server.getQconf(),
		TLSConfig:  server.getTconf(),
	}

	http3server.ListenAndServe()
}
