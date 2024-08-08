package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func CheckRoute(c echo.Context) error {
	return c.String(http.StatusOK, "Up and working!")
}
