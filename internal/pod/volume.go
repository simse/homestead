package pod

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/volume"
	"github.com/docker/docker/client"
	"github.com/simse/homestead/internal/runtime"
)

// Volume represents a Docker volume used by Homestead to store user files
type Volume struct {
	ID           string
	FriendlyName string
}

// InUse checks whether a given volume is in use by Docker
func (v *Volume) InUse() bool {
	return false
}

// CreateVolume creates a Docker volume
func CreateVolume(client *client.Client, name string) string {
	vol, err := client.VolumeCreate(context.Background(), volume.VolumeCreateBody{
		Driver:     "local",
		DriverOpts: map[string]string{},
		Labels: map[string]string{
			"createdBy": "homestead-" + runtime.Environment,
		},
		Name: name,
	})

	if err != nil {
		fmt.Println(err)
		return ""
	}

	return vol.Name
}

// DeleteVolume deletes a Docker volume
func DeleteVolume(client *client.Client, name string) error {
	return client.VolumeRemove(context.Background(), name, true)
}

// VolumeName returns the Docker name of a volume
func VolumeName(name string) string {
	return name + "-homestead"
}
