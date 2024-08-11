package db

import (
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"errors"
	"os"

	"github.com/CelestialCrafter/stella/planets"
	"github.com/charmbracelet/log"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

const dbPath = "stella.db"

type User struct {
	UserId  string           `json:"id" db:"user_id"`
	Admin   bool             `json:"admin" db:"admin"`
	Coins   uint             `json:"coins" db:"coins"`
	Planets []planets.Planet `json:"planets"`
}

type dbPlanet struct {
	Hash     string `db:"hash"`
	Features string `db:"features"`
}

var db *sqlx.DB
var NotFoundError = errors.New("object not found")

func InitDB() {
	var err error

	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		file, err := os.Create(dbPath)
		if err != nil {
			log.Fatal("could not create database file", "error", err)
		}
		file.Close()
	}

	db, err = sqlx.Connect("sqlite3", dbPath)
	if err != nil {
		log.Fatal("could not open database", "error", err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS planets (
		hash TEXT PRIMARY KEY,
		features TEXT NOT NULL,
		owner_id TEXT NOT NULL,
		FOREIGN KEY (owner_id) REFERENCES users (user_id)
	);`)
	if err != nil {
		log.Fatal("could not create planets table", "error", err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		user_id TEXT PRIMARY KEY,
		admin BOOL NOT NULL DEFAULT FALSE,
		coins INTEGER NOT NULL DEFAULT 10
	)`)
	if err != nil {
		log.Fatal("could not create users table", "error", err)
	}

	log.Info("initialized database")
}

func dbPlanetToPlanet(dbPlanet dbPlanet) (planets.Planet, error) {
	features := planets.PlanetFeatures{}
	err := json.Unmarshal([]byte(dbPlanet.Features), &features)
	if err != nil {
		return planets.Planet{}, err
	}

	decodedHash, err := hex.DecodeString(dbPlanet.Hash)
	if err != nil {
		return planets.Planet{}, err
	}

	return planets.NewPlanet(features, decodedHash), nil
}

func GetPlanet(hash string) (planets.Planet, error) {
	var dbPlanet dbPlanet
	err := db.Get(&dbPlanet, "SELECT hash, features FROM planets WHERE hash = ?", hash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return planets.Planet{}, NotFoundError
		}
		return planets.Planet{}, err
	}

	return dbPlanetToPlanet(dbPlanet)
}

func GetPlanets(userId string) ([]planets.Planet, error) {
	var dbPlanets []dbPlanet

	err := db.Select(&dbPlanets, "SELECT hash, features FROM planets WHERE owner_id = ?", userId)
	if err != nil {
		return nil, err
	}

	planets := make([]planets.Planet, len(dbPlanets))
	for i, dbPlanet := range dbPlanets {
		planets[i], err = dbPlanetToPlanet(dbPlanet)
		if err != nil {
			return nil, err
		}
	}

	return planets, nil
}

func CreatePlanet(hash string, features planets.PlanetFeatures, userId string) error {
	featuresBytes, err := json.Marshal(features)
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO planets (hash, features, owner_id) VALUES (?, ?, ?)", hash, string(featuresBytes), userId)

	return err
}

func RemovePlanet(hash string, owner string) (planets.Planet, error) {
	dbPlanet := dbPlanet{}
	err := db.Get(&dbPlanet, "SELECT hash, features FROM planets WHERE hash = ? AND owner_id = ?", hash, owner)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return planets.Planet{}, NotFoundError
		}
		return planets.Planet{}, err
	}

	_, err = db.Exec("DELETE FROM planets WHERE hash = ?", hash)
	if err != nil {
		return planets.Planet{}, err
	}

	return dbPlanetToPlanet(dbPlanet)
}

func CreateUser(id string) (User, error) {
	user := User{}
	_, err := db.Exec("INSERT INTO users (user_id) VALUES (?) ON CONFLICT DO NOTHING", id)
	if err != nil {
		return User{}, err
	}

	err = db.Get(&user, "SELECT user_id, admin, coins FROM users WHERE user_id = ?", id)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func GetUser(id string) (User, error) {
	user := User{}

	err := db.Get(&user, "SELECT user_id, admin, coins FROM users WHERE user_id = ?", id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return User{}, NotFoundError
		}
		return User{}, err
	}

	return user, nil
}

func UpdateUser(user User) error {
	_, err := db.Exec("UPDATE users SET admin = ?, coins = ? WHERE user_id = ?", user.Admin, user.Coins, user.UserId)

	return err
}

func UpdatePlanet(hash string, userId string) (planets.Planet, error) {
	_, err := db.Exec("UPDATE planets SET owner_id = ? WHERE hash = ?", userId, hash)
	if err != nil {
		return planets.Planet{}, err
	}

	return GetPlanet(hash)
}
