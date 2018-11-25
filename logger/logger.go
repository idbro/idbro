package logger

import (
	"os"

	"go.uber.org/zap"
)

// New to create a new zap logger with the given category
func New(category string) (*zap.Logger, error) {
	var config zap.Config
	if os.Getenv("IDBRO_ENV") == "production" {
		config = zap.NewProductionConfig()
	} else {
		config = zap.NewDevelopmentConfig()
	}

	// Remove caller info
	config.DisableCaller = true

	// Set the initial category
	config.InitialFields = map[string]interface{}{
		"category": category,
	}

	return config.Build()
}
