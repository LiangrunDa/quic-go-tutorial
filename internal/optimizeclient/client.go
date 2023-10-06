package optimizeclient

import (
	"os"
)

var ClientEnvs map[string]string

func StartClient() {
	ClientEnvs = map[string]string{
		"sslKeyLogFile": os.Getenv("SSLKEYLOGFILE"),
		"downloads":     os.Getenv("DOWNLOADS"),
		"requests":      os.Getenv("REQUESTS"),
	}
	client := &QconfOptimizedClient{}
	client.requestFile()
}
