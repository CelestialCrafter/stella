package planets

import (
	"math/rand"
	"os"
	"path"
	"time"
)

const (
	modelPath = "models/"
)

type PlanetType string

const (
	NormalPlanet = "normal"
	StarPlanet   = "star"
)

type PlanetFeatures struct {
	Type        PlanetType `json:"type"`
	StarNeutron bool       `json:"star_neutron"`
}

type PlanetValues struct {
	NormalSize       float32    `json:"normal_size"`
	NormalColor      [3]float32 `json:"normal_color"`
	StarBrightness   float32    `json:"star_brightness"`
	StarSize         float32    `json:"star_size"`
	StarNeutronColor [3]float32 `json:"star_neutron_color"`
}

type Planet struct {
	Features  PlanetFeatures `json:"features"`
	Values    PlanetValues   `json:"values"`
	Hash      string         `json:"hash"`
	Directory string         `json:"directory"`
}

func NewPlanet(features PlanetFeatures) Planet {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// @FIX ........? My own baabbauwb version (:
	hash := "khrbakchhhh"

	// @TODO set this to the hash instead of time.now
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	values := PlanetValues{
		// 10.0 to 20.0
		NormalSize: (r.Float32() * 10) + 10,
		// [0.0 to 255.0] * 3
		NormalColor: [3]float32{r.Float32() * 255, r.Float32() * 255, r.Float32() * 255},
		// 20.0 to 30.0
		StarSize: (r.Float32() * 10) + 20,
		// 0.0 to 5.0
		StarBrightness: (r.Float32() * 5),
		// [0.0 to 85.0] * 2 + [115.0 to 255]
		StarNeutronColor: [3]float32{r.Float32() * 85, r.Float32() * 85, (r.Float32() * 140) + 115},
	}

	return Planet{
		Hash:      hash,
		Directory: path.Join(cwd, modelPath),
		Features:  features,
		Values:    values,
	}
}
