package cmd

import (
	"flag"
	"os"

	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var rootCmd = &cobra.Command{
	Use:   "docker-proxy",
	Short: "docker-proxy is proxy for docker pull&push images",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		glog.Errorln(err)
		os.Exit(-1)
	}
}

func init() {
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.CommandLine.Parse([]string{})
	flag.Set("logtostderr", "true")
}
