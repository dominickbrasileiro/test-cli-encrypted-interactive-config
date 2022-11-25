package cliutils

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
)

func InputPassword(label string) (string, error) {
	prompt := promptui.Prompt{
		Label: label,
		Mask:  '*',
	}

	input, err := prompt.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Prompt failed %v\n", err)
		os.Exit(1)
	}

	return input, nil
}
