package server

import (
	"github.com/charmbracelet/log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	bindAddress      = ":8000"
	svelteDevAddress = "http://localhost:5173"
)

func jsonError(c echo.Context, status int, err error) error {
	return c.JSON(status, echo.Map{
		"message": err.Error(),
	})
}

func SetupServer() {
	e := echo.New()
	e.HideBanner = true

	e.Use(middleware.RequestID())
	logging(e)
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		LogErrorFunc: logPanicRecover,
	}))
	e.Use(middleware.CORS())

	setupRoutes(e)

	err := e.Start(bindAddress)
	if err != nil {
		log.Fatal("error starting server", "error", err)
	}
}
