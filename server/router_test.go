package server

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

type routerFindResult struct {
	method       string
	path         string
	expectResult int
}

const (
	ableToFindResult   = 0
	isNotFoundResult   = 1
	isNotAllowedResult = 2
)

func notFoundInTest(c echo.Context) error {
	return nil
}

func notAllowedInTest(c echo.Context) error {
	return nil
}

func TestRouting(t *testing.T) {
	echo.NotFoundHandler = notFoundInTest
	echo.MethodNotAllowedHandler = notAllowedInTest

	e := echo.New()
	routing(e)
	r := e.Router()

	pathes := []routerFindResult{
		// Get, Head is handled
		{method: http.MethodGet, path: "/v1/ping", expectResult: ableToFindResult},
		{method: http.MethodHead, path: "/v1/ping", expectResult: ableToFindResult},

		{method: http.MethodGet, path: "/v2/ping", expectResult: isNotFoundResult},
		{method: http.MethodPost, path: "/v1/ping", expectResult: isNotAllowedResult},
		{method: http.MethodPut, path: "/v1/ping", expectResult: isNotAllowedResult},

		{method: http.MethodGet, path: "/v1/status", expectResult: ableToFindResult},
		{method: http.MethodPost, path: "/v1/status", expectResult: isNotAllowedResult},
	}

	for _, one := range pathes {
		c := e.NewContext(nil, nil)
		assert.Equal(t, one.expectResult, findResult(r, &one, c))
	}
}

func findResult(r *echo.Router, oneFind *routerFindResult, c echo.Context) int {
	r.Find(oneFind.method, oneFind.path, c)
	handler := reflect.ValueOf(c.Handler()).Pointer()

	if handler == reflect.ValueOf(notFoundInTest).Pointer() {
		return isNotFoundResult
	}
	if handler == reflect.ValueOf(notAllowedInTest).Pointer() {
		return isNotAllowedResult
	}
	return ableToFindResult
}
