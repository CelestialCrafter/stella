package main

import (
	"github.com/CelestialCrafter/stella/planets"
	"github.com/CelestialCrafter/stella/server"
	"github.com/charmbracelet/log"
)

func main() {
	planet := planets.NewPlanet(planets.PlanetFeatures{
		Type: planets.StarPlanet,
	})

	err := planet.CreateModel()
	if err != nil {
		log.Error(err)
		panic(err)
	}

	server.SetupServer()
}
