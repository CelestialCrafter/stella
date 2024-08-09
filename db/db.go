package db

import (
	"database/sql"

	"os"

	"github.com/charmbracelet/log"
	_ "github.com/mattn/go-sqlite3"
)

const dbPath = "stella.db"

func InitDB() {
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		file, err := os.Create(dbPath)
		if err != nil {
			log.Fatal(err)
		}
		file.Close()
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createTableSQL := `CREATE TABLE IF NOT EXISTS planets (
        hash TEXT PRIMARY KEY,
        features INTEGER NOT NULL
    );`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
}
