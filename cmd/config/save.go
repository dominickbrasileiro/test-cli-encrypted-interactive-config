package config

import (
	"test-cli/pkg/data"

	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

func save(ctx *cli.Context, c *Config, encryptionKey string) error {
	yaml, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	return data.StoreEncrypted(
		encryptionKey,
		"config",
		string(yaml),
		true,
	)
}
