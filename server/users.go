package server

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/CelestialCrafter/stella/db"
	"github.com/labstack/echo/v4"
)

func GetUser(c echo.Context) error {
	id := c.Param("id")

	user, err := db.GetUser(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return jsonError(c, http.StatusNotFound, errors.New("user not found"))
		}

		return jsonError(c, http.StatusInternalServerError, err)
	}

	planets, err := db.GetPlanets(id)
	if err != nil {
		return jsonError(c, http.StatusInternalServerError, err)
	}

	user.Planets = planets

	return c.JSON(http.StatusOK, user)
}
