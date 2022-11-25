package main

import (
	"fmt"
	"os"
	"test-cli/cmd/config"
	"test-cli/pkg/cliutils"
	"test-cli/pkg/data"
	"test-cli/pkg/encrypter"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "test"
	app.Usage = "Test CLI"
	app.Version = "0.0.1"

	config.RegisterCommands(app, "config", []string{"c"})

	dataDir, err := cliutils.GetDefaultDataPath()
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}

	if err := data.ConfigStore(
		dataDir,
		encrypter.NewDefaultEncrypter(),
	); err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
