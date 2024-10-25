package commands

import (
	"aie/internal/server"

	"github.com/urfave/cli"
)

var StartCommand = cli.Command{
	Name:    "start",
	Aliases: []string{"start"},
	Usage:   "Starts AIE backend service",
	Action:  startAction,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:   "log-level",
			Usage:  "Set the logging level - debug, info, warn, error, trace. Defaults to info.",
			EnvVar: "LOG_LEVEL",
			Value:  "info",
		},
	},
}

func startAction(c *cli.Context) error {
	logLevel := c.String("log-level")
	// Start the AIE backend service
	server.Start(logLevel)
	return nil
}
