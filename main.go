package main

import (
	"github.com/CelestialCrafter/stella/db"
	"github.com/CelestialCrafter/stella/server"
)

func main() {
	db.InitDB()

	server.SetupServer()
}
