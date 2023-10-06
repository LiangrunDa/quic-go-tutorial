package optimizeserver

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"liangrun/utils"
	"net/http"
	"os"
	"path/filepath"

	"github.com/lucas-clemente/quic-go"
	"github.com/lucas-clemente/quic-go/http3"
)

type inMemoryFileServer struct {
	files map[string][]byte
}

func (s *inMemoryFileServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data, ok := s.files[r.URL.Path]
	if !ok {
		http.NotFound(w, r)
		return
	}
	w.Write(data)
}

func (s *inMemoryFileServer) updateFileMap() {
	dir := ServerEnvs["www"]

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		// Read the file contents
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		// Store the file contents in the map
		relPath, _ := filepath.Rel(dir, path)
		s.files["/"+relPath] = data
		return nil
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

type InMemoryServer struct {
}

func (server *InMemoryServer) getTconf() *tls.Config {
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

func (server *InMemoryServer) getQconf() *quic.Config {
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

func (server *InMemoryServer) listen() {
	addr := ServerEnvs["ip"] + ":" + ServerEnvs["port"]

	// Create an in-memory file server
	fileServer := &inMemoryFileServer{files: make(map[string][]byte)}
	fileServer.updateFileMap()

	http3server := http3.Server{
		Handler:    fileServer,
		Addr:       addr,
		QuicConfig: server.getQconf(),
		TLSConfig:  server.getTconf(),
	}

	http3server.ListenAndServe()
}
