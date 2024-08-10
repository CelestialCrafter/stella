package server

import (
	"errors"
	"net/http"

	"github.com/CelestialCrafter/stella/db"
	"github.com/CelestialCrafter/stella/planets"
	"github.com/golang-jwt/jwt/v5"
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
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*userClaims)
	id := claims.ID

	dbUser, err := db.GetUser(id)
	if err != nil {
		return jsonError(c, http.StatusBadRequest, err)
	}

	if dbUser.Coins <= 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "not enough coins"})
	}

	// @TODO restrict features to random/modifiers (once demo over)
	features := new(planets.PlanetFeatures)
	err = c.Bind(features)
	if err != nil {
		return jsonError(c, http.StatusBadRequest, err)
	}

	planet := planets.NewPlanet(*features, nil)
	err = planet.CreateModel()
	if err != nil {
		return jsonError(c, http.StatusInternalServerError, err)
	}

	err = db.CreatePlanet(planet.Hash, *features, id)
	if err != nil {
		return jsonError(c, http.StatusInternalServerError, err)
	}

	dbUser.Coins -= 1
	db.UpdateUser(dbUser)

	return c.JSON(http.StatusOK, planet)
}

func DeletePlanet(c echo.Context) error {
	hash := c.Param("hash")
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*userClaims)
	id := claims.ID

	planet, err := db.RemovePlanet(hash, id)
	if err != nil {
		if errors.Is(err, db.NotFoundError) {
			return jsonError(c, http.StatusNotFound, err)
		}

		return jsonError(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, planet)
}
