package server

import (
	"net/http"
	"strconv"

	"github.com/CelestialCrafter/stella/db"
	"github.com/CelestialCrafter/stella/planets"
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

	featuresBin, err := strconv.Atoi(c.QueryParam("features"))
	if err != nil {
		return err
	}

	planetsFeatures, err := planets.ExtractFeaturesFromBin(featuresBin)
	if err != nil {
		return err
	}

	planet := planets.NewPlanet(*planetsFeatures)

	err = planet.CreateModel()
	if err != nil {
		return err
	}

	err = db.CreatePlanet(planet.Hash, featuresBin)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"hash":     planet.Hash,
		"features": featuresBin,
	})
}
