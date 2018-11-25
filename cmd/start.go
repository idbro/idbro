package cmd

import (
	"errors"

	"github.com/idbro/idbro/server"
	"gopkg.in/urfave/cli.v1"
)

func startCommand() cli.Command {
	return cli.Command{
		Name:  "start",
		Usage: "start the idbro server",
		Action: func(c *cli.Context) error {
			port := c.Int("port")
			if port <= 0 || port > 65535 {
				return errors.New("server port out of range")
			}
			return server.Start(c, uint16(port))
		},
		Flags: []cli.Flag{
			cli.IntFlag{
				Name:  "port, p",
				Usage: "the server opening port",
				Value: 8080,
			},
		},
	}
}
