package cmd

import (
	"fmt"

	"github.com/simse/homestead/internal/pod"
	"github.com/urfave/cli/v2"
)

func CmdDev(c *cli.Context) error {
	client := pod.CreateClient()

	//pod.CreateVolume(client, "simons_test")
	/*db := config.CreateDB(config.ConfigFileLocation())
	fmt.Println(pod.PodNameExists("test", db))

	var pods []pod.Pod
	db.All(&pods)
	fmt.Println(pods)*/

	progress := make(chan float64)
	go pod.PullImage(client, "ghcr.io/simse/homestead/go", progress)

	for p := range progress {
		fmt.Printf("%.2f%%\n", p)
	}

	return nil
}
