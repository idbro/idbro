package cmd

import (
	"gopkg.in/urfave/cli.v1"
)

// Commands to get commands mounted
func Commands() []cli.Command {
	return []cli.Command{
		startCommand(),
	}
}
