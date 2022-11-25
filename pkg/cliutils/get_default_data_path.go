package cliutils

import (
	"os"
	"path/filepath"
)

func GetDefaultDataPath() (string, error) {
	base, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(base, ".test"), nil
}
