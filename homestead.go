package main

import (
	"log"
	"os"

	"github.com/simse/homestead/cmd"
	"github.com/urfave/cli/v2"
)

func main() {
	/*cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}*/

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
	/*response, err := cli.ContainerCreate(context.Background(), &container.Config{
		Image:     "ghcr.io/simse/homestead/go:latest",
		Tty:       false,
		OpenStdin: true,
	}, &container.HostConfig{
		Binds: []string{"homestead_test:/home/homestead/workspace"},
	}, &network.NetworkingConfig{}, &v1.Platform{
		Architecture: "arm64",
		OS:           "linux",
	}, "homestead_test")

	fmt.Println(response)
	fmt.Println(err)*/

	app := &cli.App{
		Name:        "Homestead",
		Description: "A tool for creating fast and disposable development environments on your machine",
		Commands: []*cli.Command{
			{
				Name:    "version",
				Aliases: []string{"v"},
				Usage:   "print Homestead version",
				Action:  cmd.CmdVersion,
			},
			{
				Name:    "dev",
				Aliases: []string{"t"},
				Usage:   "Homestead development command",
				Action:  cmd.CmdDev,
			},
			{
				Name:    "new",
				Aliases: []string{"n"},
				Usage:   "create new Homestead environment",
				Action:  cmd.CmdNew,
			},
			{
				Name:    "open",
				Aliases: []string{"o"},
				Usage:   "opens Homestead environment in VS Code",
				Action:  cmd.CmdOpen,
			},
			{
				Name:    "connect",
				Aliases: []string{"c"},
				Usage:   "connect to Homestead environment over SSH (not actually)",
				Action:  cmd.CmdConnect,
			},
			{
				Name: "delete",
				//Aliases: []string{"c"},
				Usage:  "delete Homestead environment",
				Action: cmd.CmdDelete,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
