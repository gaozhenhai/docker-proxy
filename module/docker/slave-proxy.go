package docker

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/golang/glog"
	"github.com/spf13/cobra"
)

func sendToMasterDockerProxy(w http.ResponseWriter, dockerRequest *DockerRequest) {
	dockerRequestByte, _ := json.Marshal(dockerRequest)
	dockerRequestBody := strings.NewReader(string(dockerRequestByte))
	masterProxyUrl := fmt.Sprintf("http://%s/docker/proxy", MasterDockerProxyAddr)

	req, err := http.NewRequest("POST", masterProxyUrl, dockerRequestBody)
	if err != nil {
		glog.Errorln(err)
		io.WriteString(w, err.Error())
		return
	}

	dockerProxyResponse(w, req)
}

func slaveDockerHandle(w http.ResponseWriter, r *http.Request) {
	dockerRequest := &DockerRequest{
		Host:   r.Host,
		Url:    r.RequestURI,
		Method: r.Method,
		RequestMap: map[string]string{
			"User-Agent":      r.Header.Get("User-Agent"),
			"Authorization":   r.Header.Get("Authorization"),
			"Accept-Encoding": r.Header.Get("Accept-Encoding"),
			"Connection":      r.Header.Get("Connection"),
		},
	}

	glog.Infof("[%s] host: %s, url: %s", r.Method, r.Host, r.RequestURI)
	sendToMasterDockerProxy(w, dockerRequest)
}

func SlaveDockerProxy(cmd *cobra.Command, args []string) {
	listenAddress, _ := cmd.Flags().GetString("listen")
	MasterDockerProxyAddr, _ = cmd.Flags().GetString("master-proxy")

	glog.Infof("start listen address %s", listenAddress)
	http.HandleFunc("/", slaveDockerHandle)
	http.ListenAndServe(listenAddress, nil)
}
