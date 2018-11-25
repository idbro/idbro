package server

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/idbro/idbro/logger"
	"github.com/idbro/idbro/storage"
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"gopkg.in/urfave/cli.v1"
)

type (
	// App contains the information of the service
	App struct {
		// the echo server
		e *echo.Echo
		// the logger to use
		logger *zap.Logger
		// the storage layer used
		storage storage.Storage
	}
)

const (
	// timeout for grace shutdown
	graceTimeout = 30 * time.Second
)

// Start the idbro server.
func Start(c *cli.Context, port uint16) error {
	l, err := logger.New("server")
	if err != nil {
		return err
	}

	var app App
	app.e = createServer(l)
	app.logger = l

	return app.start(port)
}

// Create the echo server
func createServer(l *zap.Logger) *echo.Echo {
	e := echo.New()

	// disable the Echo welcome print
	e.HideBanner = true
	// add routers
	addRouting(e, l)

	return e
}

// Start the echo server
func (app *App) start(port uint16) error {
	go func() {
		// Start server
		port := fmt.Sprintf(":%d", port)
		if err := app.e.Start(port); err != nil {
			app.logger.Fatal("Error start idbro server.", zap.Error(err))
		}
	}()

	app.logger.Info("idbro server started.", zap.Uint16("port", port))

	// Wait for interrupt signal to gracefully shutdown the server.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// set a timeout for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), graceTimeout)
	defer cancel()

	// gracefully shutdown
	if err := app.e.Shutdown(ctx); err != nil {
		return err
	}
	app.logger.Info("idbro server successfully shutdown.")

	// cleanup other resources
	if err := app.cleanup(); err != nil {
		return err
	}

	return nil
}

// Cleanup the app
func (app *App) cleanup() error {
	app.logger.Info("idbro server cleaning up...")

	return nil
}
