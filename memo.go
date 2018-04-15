package main

import (
	"log"
	"os"

	"github.com/ayatk/memo/cmd"
	"github.com/urfave/cli"
)

const (
	name        = "memo"
	description = "Git based idea note management tool."
	version     = "0.1.1"
)

func main() {
	app := cli.NewApp()
	app.Name = name
	app.Usage = description
	app.Version = version

	app.Commands = []cli.Command{
		cmd.NewCmd,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
