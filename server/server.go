package server

import (
	"time"

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

func getRequestId(c echo.Context) string {
	id := c.Response().Header().Get(echo.HeaderXRequestID)
	if id == "" {
		return "no id"
	}

	return id
}

func SetupServer() {
	e := echo.New()
	e.HideBanner = true

	e.Use(middleware.RequestID())
	logging(e)
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		LogErrorFunc: logPanicRecover,
	}))
	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Skipper:      middleware.DefaultSkipper,
		ErrorMessage: "response timed out",
		Timeout:      30 * time.Second,
	}))
	e.Use(middleware.CORS())

	setupRoutes(e)

	err := e.Start(bindAddress)
	if err != nil {
		log.Fatal("error starting server", "error", err)
	}
}
