package server

import (
	"github.com/charmbracelet/log"
	"github.com/labstack/echo/v4"
)

const (
	bindAddress      = ":8000"
	svelteDevAddress = "http://localhost:5173"
)

func SetupServer() {
	e := echo.New()
	e.HideBanner = true

	logging(e)
	setupRoutes(e)

	err := e.Start(bindAddress)
	if err != nil {
		log.Fatal("error starting server", "error", err)
	}
}
