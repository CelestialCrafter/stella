package main

import (
	"github.com/CelestialCrafter/stella/db"
	"github.com/CelestialCrafter/stella/planets"
	"github.com/CelestialCrafter/stella/server"
	"github.com/charmbracelet/log"
)

func main() {
	planet := planets.NewPlanet(planets.PlanetFeatures{
		Type:        planets.StarPlanet,
		StarNeutron: true,
	})

	err := planet.CreateModel()
	if err != nil {
		log.Fatal("could not create model", "error", err)
	}

	db.InitDB()
	server.SetupServer()
}
