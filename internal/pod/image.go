package pod

import (
	"context"
	"encoding/json"
	"io"
	"os/exec"

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

// Struct representing events returned from image pulling
type pullEvent struct {
	ID             string `json:"id"`
	Status         string `json:"status"`
	Error          string `json:"error,omitempty"`
	Progress       string `json:"progress,omitempty"`
	ProgressDetail struct {
		Current int `json:"current"`
		Total   int `json:"total"`
	} `json:"progressDetail"`
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

// GetImageSize will return the image size given full url and platform
func GetImageSize(imageURL string, platform string) (int, error) {
	// TODO: do this part programmatically
	// Call Docker CLI to get image manifest
	manifestBytes, err := exec.Command("docker", "manifest", "inspect", "-v", imageURL).Output()
	if err != nil {
		panic(err)
	}

	// Parse manifest
	var manifests []Manifest
	err = json.Unmarshal(manifestBytes, &manifests)
	if err != nil {
		panic(err)
	}

	for _, manifest := range manifests {
		if manifest.Architecture() == platform {
			return manifest.Size(), nil
		}
	}

	return 0, nil
}

// PullImage pulls a Homestead image from a registry given full url
func PullImage(client *client.Client, imageURL string, progress chan float64) error {
	// Get full image size
	totalSize, _ := GetImageSize(imageURL, "arm64")
	totalSize = int(float64(totalSize) * 0.95) // TODO: investigate discrepency

	// Tell Docker daemon to pull the image
	reader, err := client.ImagePull(context.Background(), imageURL, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}

	// Start reading progress
	layers := map[string]int{}
	var event *pullEvent
	decoder := json.NewDecoder(reader)

	for {
		// Docker is done pulling image
		if err := decoder.Decode(&event); err != nil {
			if err == io.EOF {
				break
			}

			panic(err)
		}

		// Register layer
		if event.Status == "Pulling fs layer" {
			// Init layerProgress struct
			layers[event.ID] = 0
		}

		// Register progress
		if event.Status == "Downloading" {
			layers[event.ID] = event.ProgressDetail.Current
		}

		// Print progress
		totalProgress := 0.0
		for _, layerDownloaded := range layers {
			totalProgress += float64(layerDownloaded) / float64(totalSize) * float64(100)
		}

		if totalProgress > 100 {
			totalProgress = 100
		}

		progress <- totalProgress
	}

	close(progress)

	return nil
}
