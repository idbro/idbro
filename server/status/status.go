package status

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

var (
	startTime time.Time
)

func init() {
	startTime = time.Now()
}

// Status endpoint
func Status(c echo.Context) error {
	now := time.Now()
	duration := now.Sub(startTime).Seconds()

	resp := fmt.Sprintf("{\"uptime\":%d}", int(duration))

	return c.JSONBlob(http.StatusOK, []byte(resp))
}
