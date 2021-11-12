package cmd

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/gernest/wow"
	"github.com/gernest/wow/spin"
	"github.com/gosimple/slug"
	"github.com/manifoldco/promptui"
	"github.com/simse/homestead/internal/config"
	"github.com/simse/homestead/internal/pod"
	"github.com/urfave/cli/v2"
)

type NewPodChoices struct {
	Name        string
	Image       string
	ImageOption int
}

func CmdNew(c *cli.Context) error {
	// Welcome user to CLI
	fmt.Println("Homestead v0.1.0")
	fmt.Println("")

	client := pod.CreateClient()
	db := config.CreateDB(config.ConfigFileLocation())
	defer db.Close()
	newPodChoices := NewPodChoices{}

	// Ask user for project name
	validate := func(input string) error {
		if pod.PodNameExists(slug.Make(input), db) {
			//lint:ignore ST1005 The message is printed to screen
			return errors.New("A pod with that name already exists")
		}

		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Name",
		Validate: validate,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return nil
	}

	newPodChoices.Name = slug.Make(result)

	// Ask user for image
	imagePrompt := promptui.Select{
		Label: "Image",
		Items: pod.DefaultImages,
		Templates: &promptui.SelectTemplates{
			Label:    "{{ .Name }}?",
			Active:   "> {{ .Name }}",
			Inactive: "  {{ .Name }}",
			Selected: "{{ \"Image:\" | faint }} {{ .Name }}",
		},
	}

	imageIndex, _, err := imagePrompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return nil
	}

	newPodChoices.Image = pod.DefaultImages[imageIndex].Name

	// Ask user for image option
	imageOptionPrompt := promptui.Select{
		Label: "Image Variant",
		Items: pod.DefaultImageOptions(newPodChoices.Image),
		Templates: &promptui.SelectTemplates{
			Label:    "{{ .Name }}?",
			Active:   "> {{ .Name }}",
			Inactive: "  {{ .Name }}",
			Selected: "{{ \"Image variant:\" | faint }} {{ .Name }}",
		},
	}

	imageOptionResult, _, err := imageOptionPrompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return nil
	}
	newPodChoices.ImageOption = imageOptionResult

	// Create space
	fmt.Println("")
	tick := []string{color.GreenString("✔")}

	// Create volume
	w := wow.New(os.Stdout, spin.Get(spin.Dots), " Creating a volume for your files...")
	w.Start()

	vol := pod.CreateVolume(client, pod.VolumeName(newPodChoices.Name))

	time.Sleep(2 * time.Second)
	w.PersistWith(spin.Spinner{Frames: tick}, " Volume created")

	// Register in database
	podDefinition := pod.Pod{
		FriendlyName: newPodChoices.Name,
		Image:        newPodChoices.Image,
		ImageOption:  newPodChoices.ImageOption,
		Volume:       vol,
	}
	db.Save(&podDefinition)
	//fmt.Println(err)

	// Pull image
	w2 := wow.New(os.Stdout, spin.Get(spin.Dots), " Pulling Homestead image (0.00%)...")
	w2.Start()

	podImageURL := pod.GetPullURL(pod.DefaultImages[imageIndex], newPodChoices.ImageOption)

	pullProgess := make(chan float64)

	go pod.PullImage(client, podImageURL, pullProgess)

	for progress := range pullProgess {
		w2.Text(" Pulling Homestead image (" + fmt.Sprintf("%.2f", progress) + "%)...")
	}

	w2.PersistWith(spin.Spinner{Frames: tick}, " Homestead image pulled")

	// Create container
	w3 := wow.New(os.Stdout, spin.Get(spin.Dots), " Creating container...")
	w3.Start()

	pod.CreateContainer(newPodChoices.Name, podImageURL, vol, client)
	pod.StartContainer(newPodChoices.Name, client)

	w3.PersistWith(spin.Spinner{Frames: tick}, " Container created")

	fmt.Println("")

	// Print instructions on how to use the newly created Homestead pod
	fmt.Println("We're all done buddy.")

	return nil
}
