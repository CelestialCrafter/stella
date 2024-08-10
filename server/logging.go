package server

import (
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
					"uri", v.URI,
					"code", v.Status,
					"error", v.Error,
					"ip", c.RealIP(),
					"id", c.Response().Header().Get(echo.HeaderXRequestID),
				)
				return nil
			}

			log.Info(
				"request",
				"uri", v.URI,
				"code", v.Status,
				"ip", c.RealIP(),
				"id", c.Response().Header().Get(echo.HeaderXRequestID),
			)
			return nil
		},
	}))
}
func logPanicRecover(c echo.Context, err error, stack []byte) error {
	log.Error(
		"request panic",
		"error",
		err,
		"id",
		c.Response().Header().Get(echo.HeaderXRequestID),
		"stack",
		string(stack),
	)
	return err
}
