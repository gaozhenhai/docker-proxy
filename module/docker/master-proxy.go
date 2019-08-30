package docker

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/golang/glog"
	"github.com/spf13/cobra"
)

func sendToHarbor(w http.ResponseWriter, dockerRequest *DockerRequest) {
	req, err := http.NewRequest(dockerRequest.Method, dockerRequest.Url, nil)
	if err != nil {
		glog.Errorln(err)
		io.WriteString(w, err.Error())
		return
	}

	for k, v := range dockerRequest.RequestMap {
		if v != "" {
			req.Header.Add(k, v)
		}
	}

	dockerProxyResponse(w, req)
}

func masterDockerHandle(w http.ResponseWriter, r *http.Request) {
	dockerRequest := &DockerRequest{}

	body, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(body, dockerRequest); err != nil {
		glog.Errorln(err)
		io.WriteString(w, err.Error())
		return
	}

	glog.Infof("[%s] host: %s, url: %s", dockerRequest.Method, r.Host, dockerRequest.Url)
	sendToHarbor(w, dockerRequest)
}

func MasterDockerProxy(cmd *cobra.Command, args []string) {
	listenAddress, _ := cmd.Flags().GetString("listen")

	glog.Infof("start listen address %s", listenAddress)
	http.HandleFunc("/docker/proxy", masterDockerHandle)
	http.ListenAndServe(listenAddress, nil)
}
