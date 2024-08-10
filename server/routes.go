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
		g.Use(middleware.StaticWithConfig(middleware.StaticConfig{
			HTML5: true,
			Root:  "server/web/dist",
		}))
	}
}

func setupRoutes(e *echo.Echo) {
	svelte(e.Group("/app"))

	r := e.Group("/")
	r.Use(jwtMiddleware())

	m := e.Group("/models")
	m.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 6,
	}))

	m.Static("/", "models")

	r.GET("/api/planet/:hash", GetPlanet)
	r.POST("/api/planet/new", NewPlanet)
	r.DELETE("/api/planet/delete", DeletePlanet)

	e.GET("/api/auth/login", Login)
	e.GET("/api/auth/callback", Callback)

}
