package utils

import (
	"os"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05.000",
	})
}
func PrintEnv() {
	if ServerEnvMap["role"] == "server" {
		for k, v := range ServerEnvMap {
			logrus.Debugf("%s %s\n", k, v)
		}
	} else {
		for k, v := range ClientEnvMap {
			logrus.Debugf("%s %s\n", k, v)
		}
	}
}
