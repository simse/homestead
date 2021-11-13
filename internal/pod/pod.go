package pod

import (
	"context"

	"github.com/asdine/storm/v3"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

// Pod represents a full Homestead environment
type Pod struct {
	ID             int `storm:"id,increment"`
	FriendlyName   string
	Image          string
	ImageOption    int
	Volume         string
	CommandOnStart string
}

// PodNameExists checks if a pod with the given name exists
func PodNameExists(name string, db *storm.DB) bool {
	var pod Pod
	err := db.One("FriendlyName", name, &pod)
	return err == nil
}

// CreateContainer will create a container given name and Docker client
func CreateContainer(name string, imageURL string, volumeName string, cli *client.Client) error {
	_, err := cli.ContainerCreate(context.Background(), &container.Config{
		Image:     imageURL,
		Tty:       false,
		OpenStdin: true,
	}, &container.HostConfig{
		Binds: []string{volumeName + ":/home/homestead/workspace"},
	}, &network.NetworkingConfig{}, &v1.Platform{
		Architecture: "arm64",
		OS:           "linux",
	}, name+"_homestead")

	return err
}

// RemoveContainer will remove a container given name
func RemoveContainer(name string, cli *client.Client) error {
	return cli.ContainerRemove(context.Background(), name+"_homestead", types.ContainerRemoveOptions{
		Force: true,
	})
}

// StartContainer will start a container given name
func StartContainer(name string, cli *client.Client) error {
	return cli.ContainerStart(context.Background(), name+"_homestead", types.ContainerStartOptions{})
}

// RunCommand will run a command in a container given name
func RunCommand(name string, command string, cli *client.Client) error {
	execId, err := cli.ContainerExecCreate(context.Background(), name+"_homestead", types.ExecConfig{
		AttachStdin:  false,
		AttachStdout: true,
		AttachStderr: true,
		Tty:          false,
		Cmd:          []string{"/bin/sh", "-c", command},
	})
	if err != nil {
		panic(err)
	}

	err = cli.ContainerExecStart(context.Background(), execId.ID, types.ExecStartCheck{})

	return err
}
