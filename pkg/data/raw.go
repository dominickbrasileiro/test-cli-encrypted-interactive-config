package data

import (
	"errors"
	"os"
	"path/filepath"
)

func StoreRaw(key string, value string, overwrite bool) error {
	s, err := getStore()
	if err != nil {
		return err
	}

	path := filepath.Join(s.dataDir, filepath.Base(key))

	if !overwrite {
		if _, err := os.Stat(path); err == nil {
			return errors.New("key already stored")
		}
	}

	if err := os.WriteFile(path, []byte(value), 0664); err != nil {
		return err
	}

	return nil
}

func ReadRaw(key string) (string, error) {
	s, err := getStore()
	if err != nil {
		return "", err
	}

	path := filepath.Join(s.dataDir, filepath.Base(key))

	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(content), nil
}
