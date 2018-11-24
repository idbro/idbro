package server

import (
	"net/http"

	"github.com/idbro/idbro/server/middleware"
	"github.com/idbro/idbro/server/status"
	"github.com/labstack/echo"
	"go.uber.org/zap"
)

// addRouting to echo
func addRouting(e *echo.Echo, l *zap.Logger) {
	// Middleware
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger(l))

	v1R := e.Group("/v1")

	// ping endpoint, just return OK for Get and Head
	v1R.Match([]string{http.MethodGet, http.MethodHead}, "/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "")
	})

	// status endpoint
	v1R.GET("/status", status.Status)
}
