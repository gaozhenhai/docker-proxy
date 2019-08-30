package cmd

import (
	"docker-proxy/module/docker"

	"github.com/spf13/cobra"
)

var dockerMasterProxyCmd = &cobra.Command{
	Use:   "docker",
	Short: "proxy docker pull images with master",
	Run:   docker.MasterDockerProxy,
}

func init() {
	dockerMasterProxyCmd.Flags().StringP("listen", "l", "0.0.0.0:7000", "proxy listen address")

	masterProxyCmd.AddCommand(dockerMasterProxyCmd)
	masterProxyCmd.MarkFlagRequired("listen")
}
