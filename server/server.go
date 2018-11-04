package server

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/idbro/idbro/logger"
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"gopkg.in/urfave/cli.v1"
)

const (
	// timeout for grace shutdown
	graceTimeout = 30 * time.Second
)

// Start the idbro server.
func Start(c *cli.Context, port uint16) error {
	e := echo.New()

	// disable the Echo welcome print
	e.HideBanner = true

	// add routers
	routing(e)

	go func() {
		// Start server
		port := fmt.Sprintf(":%d", port)
		if err := e.Start(port); err != nil {
			logger.L.Fatal("Error start idbro server.", zap.Error(err))
		}
	}()

	logger.L.Info("idbro server started.", zap.Uint16("port", port))

	// Wait for interrupt signal to gracefully shutdown the server.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// set a timeout for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), graceTimeout)
	defer cancel()

	// gracefully shutdown
	if err := e.Shutdown(ctx); err != nil {
		return err
	}
	logger.L.Info("idbro server successfully shutdown.")

	// cleanup other resources
	if err := cleanup(); err != nil {
		return err
	}

	return nil
}

func cleanup() error {
	logger.L.Info("idbro server cleaning up...")

	return nil
}
