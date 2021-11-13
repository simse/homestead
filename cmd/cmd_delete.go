package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/gernest/wow"
	"github.com/gernest/wow/spin"
	"github.com/manifoldco/promptui"
	"github.com/simse/homestead/internal/config"
	"github.com/simse/homestead/internal/pod"
	"github.com/urfave/cli/v2"
)

// CmdDelete is the delete command that removes a Homestead environment
func CmdDelete(c *cli.Context) error {
	// Establish Docker client
	client := pod.CreateClient()

	// Establish database connection
	db := config.CreateDB(config.ConfigFileLocation())

	// Get pod name
	podName := c.Args().Get(0)

	// Ask for confirmation from user
	fmt.Println("This will delete the container AND the associated volume. If you have not saved your `workspace` folder elsewhere, IT WILL BE LOST.")
	fmt.Println("")

	// Create and trigger prompt
	prompt := promptui.Prompt{
		Label:     "Delete " + podName,
		IsConfirm: true,
	}
	_, err := prompt.Run()

	// User did not confirm
	if err != nil {
		fmt.Println("Nothing has been deleted.")
		return nil
	}

	// User confirmed
	tick := []string{color.GreenString("✔")}

	// Remove container
	w := wow.New(os.Stdout, spin.Get(spin.Dots), " Removing container...")
	w.Start()
	pod.RemoveContainer(podName, client)
	w.PersistWith(spin.Spinner{Frames: tick}, " Removed container")

	// Remove volume
	w2 := wow.New(os.Stdout, spin.Get(spin.Dots), " Removing volume...")
	w2.Start()
	pod.DeleteVolume(client, pod.VolumeName(podName))
	w2.PersistWith(spin.Spinner{Frames: tick}, " Removed volume")

	// Remove from database
	var podFromDb pod.Pod
	db.One("FriendlyName", podName, &podFromDb)
	db.DeleteStruct(&podFromDb)

	fmt.Println("\nRemoved " + podName)

	return nil
}
