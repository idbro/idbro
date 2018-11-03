package main

import (
	"errors"
	"os"

	"github.com/idbro/idbro/logger"
	"github.com/idbro/idbro/server"
	"go.uber.org/zap"
	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "idbro"
	app.Version = "0.1.0"
	app.Description = "idbro command line tool"

	app.Commands = []cli.Command{
		{
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
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		logger.L.Fatal("Command line app run failed", zap.Error(err))
	}
}
