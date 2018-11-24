package middleware

import (
	"strconv"
	"time"

	"github.com/labstack/echo"
	"go.uber.org/zap"
)

const (
	// LoggerField to indicate the field in context
	LoggerField = "logger"
)

// Logger middleware to generate a sub logger for each request and log information after every request.
func Logger(logger *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			req := c.Request()
			res := c.Response()

			start := time.Now()
			requestID := c.Get(IDBroRequstIDField).(string)

			// create a new sub-logger
			l := logger.With(zap.String(IDBroRequstIDField, requestID))
			// set into context
			c.Set(LoggerField, l)

			if err = next(c); err != nil {
				c.Error(err)
			}
			stop := time.Now()

			reqPath := req.URL.Path
			if reqPath == "" {
				reqPath = "/"
			}

			bytesIn, parseErr := strconv.Atoi(req.Header.Get(echo.HeaderContentLength))
			if parseErr != nil {
				bytesIn = 0
			}

			l.Info("request",
				zap.String("path", reqPath),
				zap.String("method", req.Method),
				zap.String("real-ip", c.RealIP()),
				zap.String("host", req.Host),
				zap.Int("bytes-in", bytesIn),
				zap.Int("status", res.Status),
				zap.Int64("latency", stop.Sub(start).Nanoseconds()),
				zap.Int64("bytes-out", res.Size),
			)

			return
		}
	}
}
