package main

import (
	"github.com/CelestialCrafter/stella/planets"
	"github.com/CelestialCrafter/stella/server"
)

func main() {
	planet := planets.Planet{
		Features: planets.PlanetFeatures{
			Blotches: false,
		},
		Values: planets.PlanetValues{
			Size:          3.25,
			Color:         [3]float32{255, 150, 255},
			BlotchDensity: 3.25,
		},
		Hash: "bab78t3893irwo0",
	}

	err := planet.CreateModel()
	if err != nil {
		panic(err)
	}

	server.SetupServer()
}
