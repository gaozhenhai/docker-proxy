package cmd

import (
	"github.com/spf13/cobra"
)

var masterProxyCmd = &cobra.Command{
	Use:   "master",
	Short: "master proxy",
}

func init() {
	rootCmd.AddCommand(masterProxyCmd)
}
