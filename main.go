package main

import (
	"os"

	"github.com/idbro/idbro/cmd"
	"github.com/idbro/idbro/logger"
	"go.uber.org/zap"
	"gopkg.in/urfave/cli.v1"
)

func main() {
	l, err := logger.New("command line")
	if err != nil {
		panic(err)
	}

	app := cli.NewApp()
	app.Name = "idbro"
	app.Usage = "A command line tool for managing idbro easily"
	app.Version = "0.1.0"
	app.Description = "idbro command line tool"

	app.Commands = cmd.Commands()

	if err := app.Run(os.Args); err != nil {
		l.Fatal("run failed", zap.Error(err))
	}
}
