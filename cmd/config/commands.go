package config

import (
	"test-cli/pkg/cliutils"

	"github.com/urfave/cli"
)

func RegisterCommands(app *cli.App, name string, aliases []string) {
	app.Commands = append(app.Commands, cli.Command{
		Name:    name,
		Aliases: aliases,
		Usage:   "",
		Action: func(ctx *cli.Context) error {
			config := &Config{}

			hostname, err := cliutils.InputText("Hostname", "0.0.0.0")
			if err != nil {
				return err
			}

			port, err := cliutils.InputInteger("Port", 2545, 0, 65536)
			if err != nil {
				return err
			}

			url, err := cliutils.InputText("URL", "http://0.0.0.0:8545")
			if err != nil {
				return err
			}

			config.Hostname = hostname
			config.Port = port
			config.URL = url

			result, err := cliutils.InputPassword("Encryption Key")
			if err != nil {
				return err
			}

			return save(ctx, config, result)
		},
	})
}
