package main

import (
	"log"

	"github.com/CelestialCrafter/stella/db"
	"github.com/CelestialCrafter/stella/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("could not load .env file", "error", err)
	}

	db.InitDB()
	server.SetupServer()
}
