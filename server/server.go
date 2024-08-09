package server

import (
	"net/http"
	"net/url"

	"github.com/charmbracelet/log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	bindAddress      = ":8000"
	svelteDevAddress = "http://localhost:5173"
)

func svelte(g *echo.Group) {
	svelteDevUrl, err := url.Parse(svelteDevAddress)
	if err != nil {
		panic(err)
	}

	_, err = http.Get(svelteDevAddress)
	if err == nil {
		g.Use(middleware.Proxy(middleware.NewRoundRobinBalancer([]*middleware.ProxyTarget{{
			URL: svelteDevUrl,
		}})))
	} else {
		g.Static("/", "server/web/dist")
	}
}

func logging(e *echo.Echo) {
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:   true,
		LogURI:      true,
		LogError:    true,
		HandleError: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error != nil {
				log.Warn(
					"request errored",
					"uri", v.URI,
					"code", v.Status,
					"error", v.Error,
				)
				return nil
			}

			log.Info(
				"request",
				"uri", v.URI,
				"code", v.Status,
			)
			return nil
		},
	}))
}

func SetupServer() {
	e := echo.New()
	e.HideBanner = true

	logging(e)
	svelte(e.Group("/app"))
	e.Static("/models", "models")

	e.Logger.Fatal(e.Start(bindAddress))
}
