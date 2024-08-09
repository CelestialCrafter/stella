package server

import (
	"net/http"

	"github.com/CelestialCrafter/stella/db"
	"github.com/labstack/echo/v4"
)

func Models(c echo.Context) error {
	return c.String(http.StatusOK, "Up and working!")
}

func GetPlanet(c echo.Context) error {
	if !db.CheckHashExistance(c.Param("hash")) {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "Could not find the specefic planet",
		})
	}

	planet, err := db.GetPlanetByHash(c.Param("hash"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Error querying the database",
		})
	}

	if planet == nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "Could not find the specefic planet",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "successfuly fetched planet",
		"planet":  planet,
	})
}
