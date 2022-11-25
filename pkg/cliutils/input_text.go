package cliutils

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
)

func InputText(label string, defaultVal string) (string, error) {
	prompt := promptui.Prompt{}

	if defaultVal != "" {
		prompt.Label = fmt.Sprintf("%s (%s)", label, defaultVal)
	} else {
		prompt.Label = label
	}

	input, err := prompt.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Prompt failed %v\n", err)
		os.Exit(1)
	}

	if input == "" && defaultVal != "" {
		return defaultVal, nil
	}

	return input, nil
}
