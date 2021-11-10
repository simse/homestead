package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func CmdVersion(c *cli.Context) error {
	fmt.Println("hello world!!")

	return nil
}
