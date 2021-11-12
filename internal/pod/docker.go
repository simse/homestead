package pod

import (
	"github.com/docker/docker/client"
)

// CreateClient will create a Docker daemon client
func CreateClient() *client.Client {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	return cli
}
