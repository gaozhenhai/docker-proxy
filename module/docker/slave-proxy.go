package docker

import (
	"io"
	"net/http"

	"github.com/golang/glog"
	"github.com/spf13/cobra"
)

func slaveHandle(w http.ResponseWriter, r *http.Request) {
	glog.Infof("[%s] host: %s, url: %s", r.Method, r.Host, r.RequestURI)
	io.WriteString(w, "success")
}

func SlaveProxy(cmd *cobra.Command, args []string) {
	listenAddress, _ := cmd.Flags().GetString("listen")

	glog.Infof("start listen address %s", listenAddress)
	http.HandleFunc("/", slaveHandle)
	http.ListenAndServe(listenAddress, nil)
}
