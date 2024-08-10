package db

import (
	"database/sql"
	"errors"
	"os"

	"github.com/charmbracelet/log"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

const dbPath = "stella.db"

type Planet struct {
	Hash     string `db:"hash"`
	Features int    `db:"features"`
}

var db *sqlx.DB

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
		features INTEGER NOT NULL,
		user_id	TEXT NOT NULL,
		FOREIGN KEY (user_id) REFERENCES users (id)
	);`)
	if err != nil {
		log.Fatal("could not create planets table", "error", err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id TEXT PRIMARY KEY,
		coins INTEGER NOT NULL DEFAULT 10
	)`)
	if err != nil {
		log.Fatal("could not create users table", "error", err)
	}

	log.Info("initialized database")
}

func GetPlanetByHash(hash string) (*Planet, error) {
	var planet Planet

	err := db.Get(&planet, "SELECT hash, features FROM planets WHERE hash = ?", hash)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// No planet found with this hash
			return nil, nil
		}
		return nil, err
	}

	return &planet, nil
}

func CreatePlanet(hash string, features string, userId string) error {
	_, err := db.Exec("INSERT INTO planets (hash, features, user_id) VALUES (?, ?, ?)", hash, features, userId)

	return err
}

func RemovePlanet(hash string) (*Planet, error) {
	planet, err := GetPlanetByHash(hash)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("DELETE FROM planets WHERE hash = ?", hash)
	if err != nil {
		return nil, err
	}

	return planet, nil
}

func CreateUser(id string) error {
	_, err := db.Exec("INSERT INTO users (id) VALUES (?) ON CONFLICT DO NOTHING", id)

	return err
}
