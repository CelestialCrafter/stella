package db

import (
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
	OwnerId  string `db:"owner_id"`
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
		nickname TEXT,
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

	planet := planets.NewPlanet(features, decodedHash)
	planet.Directory = ""

	return planet, nil
}

func GetPlanet(hash string) (planets.Planet, error) {
	var dbPlanet dbPlanet
	err := db.Get(&dbPlanet, "SELECT * FROM planets WHERE hash = ?", hash)
	if err != nil {
		return planets.Planet{}, err
	}

	return dbPlanetToPlanet(dbPlanet)
}

func GetPlanets(owner string) ([]planets.Planet, error) {
	var dbPlanets []dbPlanet

	err := db.Select(&dbPlanets, "SELECT * FROM planets WHERE owner_id = ?", owner)
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

func CreatePlanet(hash string, features planets.PlanetFeatures, owner string, nickName string) (planets.Planet, error) {
	featuresBytes, err := json.Marshal(features)
	if err != nil {
		return planets.Planet{}, err
	}

	dbPlanet := dbPlanet{}
	err = db.Get(&dbPlanet, "INSERT INTO planets (hash, features, owner_id, nickname) VALUES (?, ?, ?, ?) RETURNING *", hash, string(featuresBytes), owner, nickName)
	if err != nil {
		return planets.Planet{}, err
	}

	return dbPlanetToPlanet(dbPlanet)
}

func RemovePlanet(hash string, owner string) (planets.Planet, error) {
	dbPlanet := dbPlanet{}
	err := db.Get(&dbPlanet, "DELETE FROM planets WHERE hash = ? AND owner_id = ? RETURNING *", hash, owner)
	if err != nil {
		return planets.Planet{}, err
	}

	return dbPlanetToPlanet(dbPlanet)
}

func CreateUser(id string) (User, error) {
	user := User{}
	err := db.Get(&user, "INSERT INTO users (user_id) VALUES (?) RETURNING *", id)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func GetUser(id string) (User, error) {
	user := User{}
	err := db.Get(&user, "SELECT * FROM users WHERE user_id = ?", id)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func UpdateUser(newUser User) (User, error) {
	user := User{}
	err := db.Get(&user, "UPDATE users SET admin = ?, coins = ? WHERE user_id = ? RETURNING *", newUser.Admin, newUser.Coins, newUser.UserId)
	if err != nil {
		return User{}, err
	}

	return user, err
}

func TransferPlanet(hash string, destination string, source string) (planets.Planet, error) {
	dbPlanet := dbPlanet{}
	err := db.Get(&dbPlanet, "UPDATE planets SET owner_id = ? WHERE hash = ? AND owner_id = ? RETURNING *", destination, hash, source)
	if err != nil {
		return planets.Planet{}, err
	}

	return dbPlanetToPlanet(dbPlanet)
}
