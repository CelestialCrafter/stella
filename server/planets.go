package server

import (
	"database/sql"
	"errors"
	"net/http"
	"os"
	"path"
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
		if errors.Is(err, sql.ErrNoRows) {
			return jsonError(c, http.StatusNotFound, errors.New("planet not found"))
		}
		return jsonError(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, planet)
}

func NewPlanet(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*userClaims)
	id := claims.ID

	if !claims.Admin {
		_, ok := newPlanetLocks[id]
		if !ok {
			newPlanetLocks[id] = &sync.Mutex{}
		}
		lock := newPlanetLocks[id]

		lock.Lock()
		defer lock.Unlock()
	}

	user, err := db.GetUser(id)
	if err != nil {
		return jsonError(c, http.StatusBadRequest, err)
	}

	if !user.Admin && user.Coins <= 0 {
		return jsonError(c, http.StatusBadRequest, errors.New("not enough coins"))
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

	// @FIX put create planet and update user into a tx
	planet, err = db.CreatePlanet(planet.Hash, *features, id)
	if err != nil {
		return jsonError(c, http.StatusInternalServerError, err)
	}

	if !user.Admin {
		user.Coins -= 1
		_, err = db.UpdateUser(user)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return jsonError(c, http.StatusNotFound, errors.New("user not found"))
			}
			return jsonError(c, http.StatusInternalServerError, err)
		}
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
		if errors.Is(err, sql.ErrNoRows) {
			return jsonError(c, http.StatusNotFound, errors.New("planet not found"))
		}

		return jsonError(c, http.StatusInternalServerError, err)
	}

	err = os.Remove(path.Join(os.Getenv("BLENDER_DATA_PATH"), "models/", hash+".glb"))
	if err != nil {
		return jsonError(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, planet)
}

type destination struct {
	Id string `json:"id"`
}

func TransferPlanet(c echo.Context) error {
	hash := c.Param("hash")
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*userClaims)
	source := claims.ID

	destination := new(destination)
	err := c.Bind(destination)
	if err != nil {
		return jsonError(c, http.StatusBadRequest, err)
	}

	planet, err := db.TransferPlanet(hash, destination.Id, source)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return jsonError(c, http.StatusNotFound, errors.New("planet not found"))
		}

		return jsonError(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, planet)
}
