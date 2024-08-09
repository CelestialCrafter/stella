package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Models(c echo.Context) error {
	return c.String(http.StatusOK, "Up and working!")
}
