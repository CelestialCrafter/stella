package server

import (
	"github.com/labstack/echo/v4"
)

func SetupServer() {
	e := echo.New()
	e.GET("/api/check", CheckRoute)
	e.Logger.Fatal(e.Start(":1323"))
}
