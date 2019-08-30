package cmd

import (
	"docker-proxy/module/docker"

	"github.com/spf13/cobra"
)

var dockerSlaveProxyCmd = &cobra.Command{
	Use:   "docker",
	Short: "proxy docker pull images with slave",
	Run:   docker.SlaveDockerProxy,
}

func init() {
	dockerSlaveProxyCmd.Flags().StringP("listen", "l", "0.0.0.0:6000", "proxy listen address")
	dockerSlaveProxyCmd.Flags().StringP("master-proxy", "m", "", "master proxy server address")

	slaveProxyCmd.AddCommand(dockerSlaveProxyCmd)
	dockerSlaveProxyCmd.MarkFlagRequired("listen")
	dockerSlaveProxyCmd.MarkFlagRequired("master-proxy")
}
