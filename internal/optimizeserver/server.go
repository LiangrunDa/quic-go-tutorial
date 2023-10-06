package optimizeserver

import (
	"os"
)

var ServerEnvs map[string]string

func StartServer() {
	ServerEnvs = map[string]string{
		"sslKeyLogFile": os.Getenv("SSLKEYLOGFILE"),
		"www":           os.Getenv("WWW"),
		"certs":         os.Getenv("CERTS"),
		"ip":            os.Getenv("IP"),
		"port":          os.Getenv("PORT"),
		"serverName":    os.Getenv("SERVERNAME"),
	}
	server := &InMemoryServer{}
	server.listen()
}
