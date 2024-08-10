package server

import (
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

func setupRoutes(e *echo.Echo) {
	svelte(e.Group("/app"))
	e.Static("/models", "models")

	e.GET("/api/planet/:hash", GetPlanet)
	e.POST("/api/planet/new", PostPlanet)
	e.DELETE("/api/planet/delete", DeletePlanet)

	e.GET("/api/auth/login", Login)
	e.GET("/api/auth/callback", Callback)

}
