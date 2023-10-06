package utils

import (
	"github.com/spf13/viper"
)

var ServerEnvMap map[string]string
var ClientEnvMap map[string]string

func BuildEnvMap() {
	ServerEnvMap = map[string]string{
		"role":          viper.GetString("role"),
		"sslKeyLogFile": viper.GetString("sslKeyLogFile"),
		"qLogDir":       viper.GetString("qLogDir"),
		"logs":          viper.GetString("logs"),
		"testcase":      viper.GetString("testcase"),
		"www":           viper.GetString("www"),
		"certs":         viper.GetString("certs"),
		"ip":            viper.GetString("ip"),
		"port":          viper.GetString("port"),
		"serverName":    viper.GetString("serverName"),
	}

	ClientEnvMap = map[string]string{
		"role":          viper.GetString("role"),
		"sslKeyLogFile": viper.GetString("sslKeyLogFile"),
		"qLogDir":       viper.GetString("qLogDir"),
		"logs":          viper.GetString("logs"),
		"testcase":      viper.GetString("testcase"),
		"downloads":     viper.GetString("downloads"),
		"requests":      viper.GetString("requests"),
	}
}
