package server

import (
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	bindAddress      = ":8000"
	svelteDevAddress = "http://localhost:5173"
)

func svelte(e *echo.Echo) {
	svelteDevUrl, err := url.Parse(svelteDevAddress)
	if err != nil {
		panic(err)
	}

	_, err = http.Get(svelteDevAddress)
	if err == nil {
		e.Use(middleware.Proxy(middleware.NewRoundRobinBalancer([]*middleware.ProxyTarget{{
			URL: svelteDevUrl,
		}})))
	} else {
		e.Static("/", "server/web/dist")
	}
}

func SetupServer() {
	e := echo.New()
	e.HideBanner = true

	e.Static("/models", "models")
	svelte(e)

	e.Logger.Fatal(e.Start(bindAddress))
}
