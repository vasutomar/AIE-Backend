package main

import (
	"aie/internal/commands"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := &cli.App{}
	app.Name = "AIE backend service"
	app.EnableBashCompletion = true
	app.Commands = []cli.Command{
		commands.StartCommand,
	}

	if err := app.Run(os.Args); err != nil {
		log.Print(err)
	}
}
