package server

import (
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

	featuresBin, err := strconv.Atoi(c.QueryParam("features"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	planetsFeatures, err := planets.ExtractFeaturesFromBin(featuresBin)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	planet := planets.NewPlanet(*planetsFeatures)

	err = planet.CreateModel()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	err = db.CreatePlanet(planet.Hash, featuresBin, "") // @TODO FIX THIS
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"hash":     planet.Hash,
		"features": featuresBin,
	})
}
