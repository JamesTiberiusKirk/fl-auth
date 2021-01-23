package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

// Test this is for testing packages
func HelloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "helloworld\n")
}
