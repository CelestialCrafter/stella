package server

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/CelestialCrafter/stella/db"
	"github.com/CelestialCrafter/stella/planets"
	"github.com/labstack/echo/v4"
)

func GetPlanet(c echo.Context) error {
	planet, err := db.GetPlanetByHash(c.Param("hash"))
	if err != nil {
		return jsonError(c, http.StatusInternalServerError, err)
	}

	if planet == nil {
		return jsonError(c, http.StatusNotFound, errors.New("planet not found"))
	}

	return c.JSON(http.StatusOK, echo.Map{
		"planet": planet,
	})
}

func PostPlanet(c echo.Context) error {
	if c.QueryParam("features") == "" {
		return jsonError(c, http.StatusBadRequest, errors.New("features were not provided"))
	}

	featuresString := c.QueryParam("features")
	planetFeatures := planets.PlanetFeatures{}
	err := json.Unmarshal([]byte(featuresString), &planetFeatures)
	if err != nil {
		return jsonError(c, http.StatusInternalServerError, err)
	}

	planet := planets.NewPlanet(planetFeatures)
	err = planet.CreateModel()
	if err != nil {
		return jsonError(c, http.StatusInternalServerError, err)
	}

	// @TODO add user id
	err = db.CreatePlanet(planet.Hash, featuresString, "")
	if err != nil {
		return jsonError(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"hash":     planet.Hash,
		"features": planetFeatures,
	})
}

func DeletePlanet(c echo.Context) error {
	if c.QueryParam("hash") == "" {
		return jsonError(c, http.StatusBadRequest, errors.New("hash was not provided"))
	}

	hash := c.QueryParam("hash")

	planet, err := db.RemovePlanet(hash)
	if err != nil {
		return jsonError(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"planet": planet,
	})
}
