package cmd

import (
	"github.com/spf13/cobra"
)

var slaveProxyCmd = &cobra.Command{
	Use:   "slave",
	Short: "slave proxy",
}

func init() {
	rootCmd.AddCommand(slaveProxyCmd)
}
