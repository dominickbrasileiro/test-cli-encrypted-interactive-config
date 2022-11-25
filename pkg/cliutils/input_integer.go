package cliutils

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/manifoldco/promptui"
)

func InputInteger(label string, defaultVal int, min int, max int) (int, error) {
	prompt := promptui.Prompt{}

	if defaultVal != 0 {
		prompt.Label = fmt.Sprintf("%s (%d)", label, defaultVal)
	} else {
		prompt.Label = label
	}

	input, err := prompt.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Prompt failed %v\n", err)
		os.Exit(1)
	}

	if input == "" && defaultVal != 0 {
		return defaultVal, nil
	}

	result, err := strconv.Atoi(input)
	if err != nil {
		return 0, errors.New("invalid number")
	}

	if result < min {
		return 0, fmt.Errorf("number must be higher or equal than %d", min)
	}

	if result > max {
		return 0, fmt.Errorf("number must be lower or equal than %d", max)
	}

	return result, nil
}
