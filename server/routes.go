package server

import (
	"net/http"
	"net/url"
	"time"

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
		e.Static("/public", "server/web/dist/")
	}
}

func setupRoutes(e *echo.Echo) {
	svelte(e)

	s := e.Group("/sounds")
	s.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 6,
	}))
	s.Static("/", "sounds")

	m := e.Group("/models")
	m.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 6,
	}))

	m.Static("/", "blender/models")
	e.GET("/", func(c echo.Context) error {
		return c.Redirect(http.StatusSeeOther, "/app")
	})

	a := e.Group("/api")
	a.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Skipper:      middleware.DefaultSkipper,
		ErrorMessage: "response timed out",
		Timeout:      30 * time.Second,
	}))

	r := a.Group("")
	r.Use(jwtMiddleware())

	// do not move this to the a group - having a timeout breaks hijacking
	e.GET("/api/planet/play", PlayPlanet).Name = "play-planet"

	a.GET("/planet/:hash", GetPlanet).Name = "get-planet"
	r.DELETE("/planet/:hash", DeletePlanet).Name = "delete-planet"
	r.POST("/planet/new", NewPlanet).Name = "create-planet"
	r.POST("/planet/transfer/:hash", TransferPlanet).Name = "give-planet"

	a.GET("/user/:id", GetUser).Name = "get-user"

	r.PUT("/key/new", NewApiKey).Name = "create-api-key"

	e.GET("/auth/login", Login).Name = "oauth-login"
	e.GET("/auth/callback", Callback).Name = "oauth-callback"

	logRoutes(e)
}
