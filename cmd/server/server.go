package main

import (
	"os"

	"liangrun/internal/server"
	"liangrun/utils"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var ServerRootCmd = &cobra.Command{
	Use: "server",
}

var handshakeServer = &cobra.Command{
	Use: "handshake",
	Run: func(cmd *cobra.Command, args []string) {
		server.HandshakeListen()
	},
}

var transferServer = &cobra.Command{
	Use: "transfer",
	Run: func(cmd *cobra.Command, args []string) {
		server.TransferListen()
	},
}

var multiHandshakeServer = &cobra.Command{
	Use: "multihandshake",
	Run: func(cmd *cobra.Command, args []string) {
		server.MultiHandshakeListen()
	},
}

var versionNegotiationServer = &cobra.Command{
	Use: "versionnegotiation",
	Run: func(cmd *cobra.Command, args []string) {
		server.VersionNegotiationListen()
	},
}

var retryServer = &cobra.Command{
	Use: "retry",
	Run: func(cmd *cobra.Command, args []string) {
		server.RetryListen()
	},
}

var chacha20Server = &cobra.Command{
	Use: "chacha20",
	Run: func(cmd *cobra.Command, args []string) {
		server.Chacha20Listen()
	},
}

var transportParameterServer = &cobra.Command{
	Use: "transportparameter",
	Run: func(cmd *cobra.Command, args []string) {
		server.TransportParameterListen()
	},
}

var resumptionServer = &cobra.Command{
	Use: "resumption",
	Run: func(cmd *cobra.Command, args []string) {
		server.ResumptionListen()
	},
}

var zeroRTTServer = &cobra.Command{
	Use: "zerortt",
	Run: func(cmd *cobra.Command, args []string) {
		server.ZeroRTTListen()
	},
}

var goodputServer = &cobra.Command{
	Use: "goodput",
	Run: func(cmd *cobra.Command, args []string) {
		server.GoodputListen()
	},
}


func init() {
	ServerRootCmd.AddCommand(handshakeServer)
	ServerRootCmd.AddCommand(transferServer)
	ServerRootCmd.AddCommand(multiHandshakeServer)
	ServerRootCmd.AddCommand(versionNegotiationServer)
	ServerRootCmd.AddCommand(retryServer)
	ServerRootCmd.AddCommand(chacha20Server)
	ServerRootCmd.AddCommand(transportParameterServer)
	ServerRootCmd.AddCommand(resumptionServer)
	ServerRootCmd.AddCommand(zeroRTTServer)
	ServerRootCmd.AddCommand(goodputServer)

	viper.Set("role", "server")
	viper.BindEnv("sslKeyLogFile", "SSLKEYLOGFILE")
	viper.BindEnv("qLogDir", "QLOGDIR")
	viper.BindEnv("logs", "LOGS")
	viper.BindEnv("testcase", "TESTCASE")
	viper.BindEnv("www", "WWW")
	viper.BindEnv("certs", "CERTS")
	viper.BindEnv("ip", "IP")
	viper.BindEnv("port", "PORT")
	viper.BindEnv("serverName", "SERVERNAME")
	utils.BuildEnvMap()
}

func ServerExecute() {
	var cmdFound bool
	cmds := ServerRootCmd.Commands()

	for _, a := range cmds {
		for _, b := range os.Args[1:] {
			if a.Name() == b {
				cmdFound = true
				break
			}
		}
	}
	if !cmdFound {
		os.Exit(127)
	}
	if err := ServerRootCmd.Execute(); err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}

func main() {
	ServerExecute()
}
