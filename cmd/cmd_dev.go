package cmd

import (
	"fmt"

	"github.com/simse/homestead/internal/config"
	"github.com/simse/homestead/internal/pod"
	"github.com/urfave/cli/v2"
)

func CmdDev(c *cli.Context) error {
	//client := pod.CreateClient()

	//pod.CreateVolume(client, "simons_test")
	db := config.CreateDB(config.ConfigFileLocation())
	fmt.Println(pod.PodNameExists("test", db))

	var pods []pod.Pod
	db.All(&pods)
	fmt.Println(pods)

	return nil
}
