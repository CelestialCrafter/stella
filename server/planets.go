package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/CelestialCrafter/stella/db"
	"github.com/CelestialCrafter/stella/planets"
	"github.com/labstack/echo/v4"
)

func GetPlanet(c echo.Context) error {
	if !db.CheckPlanetExistance(c.Param("hash")) {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "Could not find the specefic planet",
		})
	}

	planet, err := db.GetPlanetByHash(c.Param("hash"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "error querying the database",
		})
	}

	if planet == nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "could not find the specefic planet",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "successfuly fetched planet",
		"planet":  planet,
	})
}

func NewPlanet(c echo.Context) error {
	if c.QueryParam("features") == "" {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "no features bitfield was provided",
		})
	}

	featuresString := c.QueryParam("features")
	planetFeatures := planets.PlanetFeatures{}
	err := json.Unmarshal([]byte(featuresString), &planetFeatures)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	planet := planets.NewPlanet(planetFeatures)
	err = planet.CreateModel()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	// @TODO add user id
	err = db.CreatePlanet(planet.Hash, featuresString, "")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"hash":     planet.Hash,
		"features": planetFeatures,
	})
}
