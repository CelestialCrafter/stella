package server

import (
	"errors"
	"net/http"
	"os"
	"sync"

	"github.com/CelestialCrafter/stella/db"
	"github.com/CelestialCrafter/stella/planets"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

var newPlanetLocks = map[string]*sync.Mutex{}

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

func NewPlanet(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*userClaims)
	id := claims.ID

	_, ok := newPlanetLocks[id]
	if !ok {
		newPlanetLocks[id] = &sync.Mutex{}
	}
	lock := newPlanetLocks[id]

	lock.Lock()
	defer lock.Unlock()

	user, err := db.GetUser(id)
	if err != nil {
		return jsonError(c, http.StatusBadRequest, err)
	}

	if !user.Admin && user.Coins <= 0 {
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

	user.Coins -= 1
	err = db.UpdateUser(user)
	if err != nil {
		return jsonError(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, planet)
}

func DeletePlanet(c echo.Context) error {
	hash := c.Param("hash")
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*userClaims)
	id := claims.ID

	planet, err := db.RemovePlanet(hash, id)
	if err != nil {
		if errors.Is(err, db.NotFoundError) {
			return jsonError(c, http.StatusNotFound, err)
		}

		return jsonError(c, http.StatusInternalServerError, err)
	}

	err = os.Remove("models/" + hash + "glb")
	if err != nil {
		return jsonError(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, planet)
}

func UpdatePlanet(c echo.Context) error {
	hash := c.Param("hash")
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*userClaims)
	id := claims.ID

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

	// Comment this following line if you want to keep all the planets that have been updated and that are not in the db
	err = os.Remove("models/" + hash + "glb")
	if err != nil {
		return jsonError(c, http.StatusInternalServerError, err)
	}

	err = db.UpdatePlanet(hash, planet.Hash, features, id)
	if err != nil {
		if errors.Is(err, db.NotFoundError) {
			return jsonError(c, http.StatusNotFound, err)
		}

		return jsonError(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, planet)
}
