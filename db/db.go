package db

import (
	"database/sql"
	"os"

	"github.com/charmbracelet/log"
	_ "github.com/mattn/go-sqlite3"
)

const dbPath = "stella.db"

type Planet struct {
	Hash     string
	Features int
}

var db *sql.DB

func InitDB() {
	var err error

	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		file, err := os.Create(dbPath)
		if err != nil {
			log.Fatal("Failed to create database file:", err)
		}
		file.Close()
	}

	db, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}

	createTableSQL := `CREATE TABLE IF NOT EXISTS planets (
        hash TEXT PRIMARY KEY,
        features INTEGER NOT NULL
    );`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}

	log.Info("Database initialized successfully")
}

func CloseDB() {
	if db != nil {
		db.Close()
		log.Info("Database connection closed")
	}
}

func CheckHashExistance(hash string) bool {
	var exists bool
	db.QueryRow("SELECT 1 FROM planets WHERE hash = ? LIMIT 1", hash).Scan(&exists)

	return exists
}

func GetPlanetByHash(hash string) (*Planet, error) {
	var planet Planet

	err := db.QueryRow("SELECT hash, features FROM planets WHERE hash = ?", hash).Scan(&planet.Hash, &planet.Features)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No planet found with this hash
		}
		return nil, err // An error occurred during the query
	}

	return &planet, nil
}
