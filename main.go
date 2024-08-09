package main

import (
	"github.com/CelestialCrafter/stella/planets"
	"github.com/CelestialCrafter/stella/server"
)

func main() {
	planet := planets.NewPlanet(planets.PlanetFeatures{
		Type: planets.StarPlanet,
	})

	err := planet.CreateModel()
	if err != nil {
		panic(err)
	}

	server.SetupServer()
}
