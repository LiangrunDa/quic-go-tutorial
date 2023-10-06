package main

import (
	"liangrun/internal/client"
	"liangrun/utils"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var ClientRootCmd = &cobra.Command{
	Use: "client",
}

var handshakeClient = &cobra.Command{
	Use: "handshake",
	Run: func(cmd *cobra.Command, args []string) {
		client.HandshakeRequest()
	},
}

var transferClient = &cobra.Command{
	Use: "transfer",
	Run: func(cmd *cobra.Command, args []string) {
		client.TransferRequest()
	},
}

var multihandShakeClient = &cobra.Command{
	Use: "multihandshake",
	Run: func(cmd *cobra.Command, args []string) {
		client.MultiHandshakeRequest()
	},
}

var versionNegotiationClient = &cobra.Command{
	Use: "versionnegotiation",
	Run: func(cmd *cobra.Command, args []string) {
		client.VersionNegotiationRequest()
	},
}

var retryClient = &cobra.Command{
	Use: "retry",
	Run: func(cmd *cobra.Command, args []string) {
		client.RetryRequest()
	},
}

var chacha20Client = &cobra.Command{
	Use: "chacha20",
	Run: func(cmd *cobra.Command, args []string) {
		client.Chacha20Request()
	},
}

var transportParameterClient = &cobra.Command{
	Use: "transportparameter",
	Run: func(cmd *cobra.Command, args []string) {
		client.TransportParameterRequest()
	},
}

var resumptionClient = &cobra.Command{
	Use: "resumption",
	Run: func(cmd *cobra.Command, args []string) {
		client.ResumptionRequest()
	},
}

var zeroRTTClient = &cobra.Command{
	Use: "zerortt",
	Run: func(cmd *cobra.Command, args []string) {
		client.ZeroRTTRequest()
	},
}

var goodputClient = &cobra.Command{
	Use: "goodput",
	Run: func(cmd *cobra.Command, args []string) {
		client.GoodputRequest()
	},
}

func init() {
	ClientRootCmd.AddCommand(handshakeClient)
	ClientRootCmd.AddCommand(transferClient)
	ClientRootCmd.AddCommand(multihandShakeClient)
	ClientRootCmd.AddCommand(versionNegotiationClient)
	ClientRootCmd.AddCommand(retryClient)
	ClientRootCmd.AddCommand(chacha20Client)
	ClientRootCmd.AddCommand(transportParameterClient)
	ClientRootCmd.AddCommand(resumptionClient)
	ClientRootCmd.AddCommand(zeroRTTClient)
	ClientRootCmd.AddCommand(goodputClient)

	viper.Set("role", "client")
	viper.BindEnv("sslKeyLogFile", "SSLKEYLOGFILE")
	viper.BindEnv("qLogDir", "QLOGDIR")
	viper.BindEnv("logs", "LOGS")
	viper.BindEnv("testcase", "TESTCASE")
	viper.BindEnv("downloads", "DOWNLOADS")
	viper.BindEnv("requests", "REQUESTS")
	utils.BuildEnvMap()
}

func ClientExecute() {
	var cmdFound bool
	cmds := ClientRootCmd.Commands()

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
	if err := ClientRootCmd.Execute(); err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}

func main() {
	ClientExecute()
}
