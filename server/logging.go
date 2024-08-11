package server

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

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
					"method", c.Request().Method,
					"uri", v.URI,
					"code", v.Status,
					"error", v.Error,
					"ip", c.RealIP(),
					"id", getRequestId(c),
				)
				return nil
			}

			log.Info(
				"request",
				"method", c.Request().Method,
				"uri", v.URI,
				"code", v.Status,
				"ip", c.RealIP(),
				"id", getRequestId(c),
			)
			return nil
		},
	}))
}

func logPanicRecover(c echo.Context, err error, stack []byte) error {
	log.Error(
		"request panic",
		"error", err,
		"id", getRequestId(c),
		"stack", string(stack),
	)
	return err
}

func logRoutes(e *echo.Echo) {
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
