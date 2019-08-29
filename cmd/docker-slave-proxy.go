package cmd

import (
	"docker-proxy/module/docker"

	"github.com/spf13/cobra"
)

var dockerSlaveProxyCmd = &cobra.Command{
	Use:   "docker",
	Short: "proxy docker pull images with slave",
	Run:   docker.SlaveProxy,
}

func init() {
	dockerSlaveProxyCmd.Flags().StringP("listen", "l", "127.0.0.1:8000", "proxy listen address")

	slaveProxyCmd.AddCommand(dockerSlaveProxyCmd)
	dockerSlaveProxyCmd.MarkFlagRequired("listen")
}
