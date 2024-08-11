package server

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func svelte(e *echo.Echo) {
	g := e.Group("/app")
	svelteDevUrl, err := url.Parse(svelteDevAddress)
	if err != nil {
		panic(err)
	}

	_, err = http.Get(svelteDevAddress)
	if err == nil {
		g.Use(middleware.Proxy(middleware.NewRoundRobinBalancer([]*middleware.ProxyTarget{{
			URL: svelteDevUrl,
		}})))
		e.Static("/public", "server/web/public/")
	} else {
		g.Use(middleware.StaticWithConfig(middleware.StaticConfig{
			HTML5: true,
			Root:  "server/web/dist/",
		}))
	}
}

func setupRoutes(e *echo.Echo) {
	svelte(e)

	m := e.Group("/models")
	m.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 6,
	}))

	m.Static("/", "models")

	a := e.Group("/api")
	r := a.Group("")
	r.Use(jwtMiddleware())

	a.GET("/planet/:hash", GetPlanet).Name = "get-planet"
	r.DELETE("/planet/:hash", DeletePlanet).Name = "delete-planet"
	r.POST("/planet/new", NewPlanet).Name = "create-planet"

	a.GET("/user/:id", GetUser).Name = "get-user"

	r.PUT("/key/new", NewApiKey).Name = "create-api-key"

	e.GET("/auth/login", Login).Name = "oauth-login"
	e.GET("/auth/callback", Callback).Name = "oauth-callback"

	echoRoutes := e.Routes()
	routes := make([]interface{}, 0)
	for _, route := range echoRoutes {
		if route.Method == "echo_route_not_found" || strings.Contains(route.Name, "StaticDirectoryHandler") {
			continue
		}

		routes = append(routes, route.Name, fmt.Sprintf("%s %s", route.Method, route.Path))
	}
	log.Info("registered routes", routes...)
}
