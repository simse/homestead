package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

func main() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	// Create volume
	/* cli.VolumeCreate(context.Background(), volume.VolumeCreateBody{
		Driver:     "local",
		DriverOpts: map[string]string{},
		Labels: map[string]string{
			"createdBy": "homestead",
		},
		Name: "homestead_test",
	}) */

	// Create container
	response, err := cli.ContainerCreate(context.Background(), &container.Config{
		Image:     "simsemand/homestead/go:latest",
		Tty:       false,
		OpenStdin: true,
	}, &container.HostConfig{
		Binds: []string{"homestead_test:/home/homestead/workspace"},
	}, &network.NetworkingConfig{}, &v1.Platform{
		Architecture: "arm64",
		OS:           "linux",
	}, "homestead_test")

	fmt.Println(response)
	fmt.Println(err)
}
