package cmd

import (
	"fmt"

	"github.com/simse/homestead/internal/config"
	"github.com/simse/homestead/internal/pod"
	"github.com/urfave/cli/v2"
)

// CmdConnect is a command to connect to a Homestead pod
func CmdConnect(c *cli.Context) error {
	// Connect to DB
	db := config.CreateDB(config.ConfigFileLocation())

	// Get pod name
	podName := c.Args().Get(0)

	// Check it exists
	if !pod.PodNameExists(podName, db) {
		return cli.Exit("Pod does not exist", 1)
	}

	// exec.Command("docker", "exec", "-it", podName+"-homestead", "/bin/bash").Run()

	fmt.Println("run: "+"docker", "exec", "-it", podName+"_homestead", "/bin/bash")

	return nil
}
