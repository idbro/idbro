package middleware

import (
	"encoding/base64"
	"math/rand"
	"time"

	"github.com/labstack/echo"
)

const (
	byteLen = 24
	strLen  = 32

	// IDBroRequstIDField to indicate single request
	IDBroRequstIDField = "IDBro-Request-ID"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func getRandomString() string {
	b := make([]byte, byteLen)
	rand.Read(b)
	return base64.RawURLEncoding.EncodeToString(b)
}

// RequestID for Echo to generate a random ID for each request
func RequestID() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(IDBroRequstIDField, getRandomString())
			return next(c)
		}
	}
}
