package cmd

import (
	"docker-proxy/module/docker"

	"github.com/spf13/cobra"
)

var dockerMasterProxyCmd = &cobra.Command{
	Use:   "docker",
	Short: "proxy docker pull images with master",
	Run:   docker.MasterProxy,
}

func init() {
	masterProxyCmd.AddCommand(dockerMasterProxyCmd)
}
