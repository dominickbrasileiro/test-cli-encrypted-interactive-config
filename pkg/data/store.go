package data

import (
	"errors"
	"os"
	"test-cli/pkg/encrypter"
)

type store struct {
	dataDir   string
	encrypter encrypter.Encrypter
}

var s *store

func ConfigStore(dataDir string, encrypter encrypter.Encrypter) error {
	if _, err := os.Stat(dataDir); err != nil {
		if !os.IsNotExist(err) {
			return err
		}

		if err := os.MkdirAll(dataDir, 0775); err != nil {
			return err
		}
	}

	s = &store{
		dataDir:   dataDir,
		encrypter: encrypter,
	}

	return nil
}

func getStore() (*store, error) {
	if s == nil {
		return nil, errors.New("store is not configured")
	}

	return s, nil
}
