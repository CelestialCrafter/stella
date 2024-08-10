package server

import (
	"errors"
	"net/http"

	"github.com/CelestialCrafter/stella/db"
	"github.com/labstack/echo/v4"
)

func GetUser(c echo.Context) error {
	user, err := db.GetUser(c.Param("id"))
	if err != nil {
		if errors.Is(err, db.NotFoundError) {
			return jsonError(c, http.StatusNotFound, err)
		}
		return jsonError(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, user)
}
