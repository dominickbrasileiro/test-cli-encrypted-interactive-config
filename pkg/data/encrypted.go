package data

import (
	"errors"
	"os"
	"path/filepath"
)

func StoreEncrypted(secret string, key string, value string, overwrite bool) error {
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

	encrypted, err := s.encrypter.Encrypt(secret, value)
	if err != nil {
		return err
	}

	if err := os.WriteFile(path, []byte(encrypted), 0664); err != nil {
		return err
	}

	return nil
}

func ReadEncrypted(secret, key string) (string, error) {
	s, err := getStore()
	if err != nil {
		return "", err
	}

	path := filepath.Join(s.dataDir, filepath.Base(key))

	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return s.encrypter.Decrypt(secret, string(content))
}
