package logger

import (
	"log"
	"os"

	"go.uber.org/zap"
)

var (
	// L is the short name for Logger
	L *zap.Logger
)

func init() {
	var config zap.Config
	if os.Getenv("IDBRO_ENV") == "production" {
		config = zap.NewProductionConfig()
	} else {
		config = zap.NewDevelopmentConfig()
	}
	// Remove caller info
	config.DisableCaller = true

	var err error
	L, err = config.Build()

	if err != nil {
		log.Fatal(err)
	}
}
