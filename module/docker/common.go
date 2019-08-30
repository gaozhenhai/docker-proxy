package docker

import (
	"io"
	"net/http"

	"github.com/golang/glog"
)

func dockerProxyResponse(w http.ResponseWriter, req *http.Request) {
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		glog.Errorln(err)
		io.WriteString(w, err.Error())
		return
	}
	defer response.Body.Close()

	for _, k := range DockerResponseHeaders {
		if v := response.Header.Get(k); v != "" {
			w.Header().Add(k, v)
		}
	}
	w.WriteHeader(response.StatusCode)

	io.Copy(w, response.Body)
}
