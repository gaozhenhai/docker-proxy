package docker

var (
	MasterDockerProxyAddr string
	DockerResponseHeaders = []string{
		"Content-Type",
		"Docker-Distribution-Api-Version",
		"Www-Authenticate",
		"Content-Encoding",
	}
)

type DockerRequest struct {
	Host       string            `json:"host"`
	Url        string            `json:"url"`
	Method     string            `json:"Method"`
	RequestMap map[string]string `json:"requestMap"`
}
