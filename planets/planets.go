package planets

import (
	"os"
	"path"
)

const (
	modelPath = "models/"
)

type PlanetValues struct {
	Size          float32    `json:"size"`
	Color         [3]float32 `json:"color"`
	BlotchDensity float32    `json:"blotchDensity"`
}

type PlanetFeatures struct {
	Blotches bool `json:"blotches"`
}

type Planet struct {
	Features PlanetFeatures `json:"features"`
	Values   PlanetValues   `json:"values"`
	Hash     string         `json:"hash"`
	FilePath string         `json:"filepath"`
}

func NewPlanet(features PlanetFeatures) Planet {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// @FIX ........?
	hash := "baabbauwb"

	modelPath := path.Join(cwd, modelPath, hash+".glb")

	// @FIX ...?
	values := PlanetValues{
		Size:          12,
		Color:         [3]float32{56, 74, 56},
		BlotchDensity: 3.25,
	}

	return Planet{
		Hash:     hash,
		FilePath: modelPath,
		Features: features,
		Values:   values,
	}
}
