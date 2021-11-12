package cmd

import (
	"os/exec"

	"github.com/urfave/cli/v2"
)

// CmdOpen is a command to open Homestead pod in VS Code
func CmdOpen(c *cli.Context) error {
	exec.Command("code", "--folder-uri=vscode-remote://ms-vscode-remote.remote-containers/").Run()

	return nil
}
