package server

import (
	"errors"
	"net/http"

	"github.com/CelestialCrafter/stella/db"
	"github.com/CelestialCrafter/stella/planets"
	"github.com/labstack/echo/v4"
)

func GetPlanet(c echo.Context) error {
	planet, err := db.GetPlanet(c.Param("hash"))
	if err != nil {
		if errors.Is(err, db.NotFoundError) {
			return jsonError(c, http.StatusNotFound, err)
		}
		return jsonError(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, planet)
}

func GetAllPlanets(c echo.Context) error {
	planets, err := db.GetPlanets(c.Param("id"))
	if err != nil {
		return jsonError(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, planets)
}

func NewPlanet(c echo.Context) error {
	// @TODO restrict features to random/modifiers (once demo over)
	features := new(planets.PlanetFeatures)
	err := c.Bind(features)
	if err != nil {
		return jsonError(c, http.StatusBadRequest, err)
	}

	planet := planets.NewPlanet(*features, nil)
	err = planet.CreateModel()
	if err != nil {
		return jsonError(c, http.StatusInternalServerError, err)
	}

	// @TODO add user id
	err = db.CreatePlanet(planet.Hash, *features, "")
	if err != nil {
		return jsonError(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, planet)
}

func DeletePlanet(c echo.Context) error {
	hash := c.Param("hash")
	planet, err := db.RemovePlanet(hash)
	if err != nil {
		if errors.Is(err, db.NotFoundError) {
			return jsonError(c, http.StatusNotFound, err)
		}

		return jsonError(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, planet)
}
