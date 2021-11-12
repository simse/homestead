package pod

import (
	"context"
	"io"
	"io/ioutil"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// Image represents a Docker image that serves as a Homestead dev environment
type Image struct {
	Name    string
	URL     string
	Options []ImageOption
}

// ImageOption is a options for a Docker image such as selecting a Python version
type ImageOption struct {
	Name string
	Tag  string
}

// DefaultImages contains a list of known images Homestead can use
var DefaultImages []Image = []Image{
	{
		Name: "Go",
		URL:  "ghcr.io/simse/homestead/go",
		Options: []ImageOption{
			{
				Name: "1.17",
				Tag:  "latest",
			},
		},
	},
}

// GetPullURL generates a pull URL from an image and an option index
func GetPullURL(image Image, optionsIndex int) string {
	return image.URL + ":" + image.Options[optionsIndex].Tag
}

// DefaultImageOptions returns the options for an image
func DefaultImageOptions(imageName string) []ImageOption {
	for _, image := range DefaultImages {
		if image.Name == imageName {
			return image.Options
		}
	}
	return nil
}

// PullImage pulls a Homestead image from a registry given full url
func PullImage(client *client.Client, imageURL string) error {
	reader, err := client.ImagePull(context.Background(), imageURL, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	io.Copy(ioutil.Discard, reader)

	// TODO: decode the output and check for progress

	return nil
}
