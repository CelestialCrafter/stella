package planets

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/rand"
	"os"
	"path"

	"github.com/CelestialCrafter/stella/common"
	"github.com/CelestialCrafter/stella/utils"
)

const (
	modelPath = "models/"
)
const (
	NormalPlanet = "normal"
	StarPlanet   = "star"
)

type PlanetFeatures struct {
	Type        string `json:"type"`
	StarNeutron bool   `json:"star_neutron"`
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

func frangeWrapper(r *rand.Rand) func(min float32, max float32) float32 {
	return func(min, max float32) float32 {
		return (r.Float32() * (max - min)) + min
	}
}

func NewPlanet(features PlanetFeatures) Planet {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	newHash := common.Hash()
	newHashInt := int64(binary.BigEndian.Uint32(newHash))
	r := rand.New(rand.NewSource(newHashInt))
	frange := frangeWrapper(r)

	values := PlanetValues{
		NormalSize:       frange(10, 20),
		NormalColor:      [3]float32{frange(0, 255), frange(0, 255), frange(0, 255)},
		StarSize:         frange(20, 30),
		StarBrightness:   frange(0, 5),
		StarNeutronColor: [3]float32{frange(0, 255), frange(130, 255), 255},
	}

	return Planet{
		Hash:      hex.EncodeToString(newHash),
		Directory: path.Join(cwd, modelPath),
		Features:  features,
		Values:    values,
	}
}

func ExtractFeaturesFromBin(featuresBin int) (*PlanetFeatures, error) {
	// featuresBinSlice := utils.SplitInt(featuresBin)
	featuresBinSlice := []int{}

	var planetType string
	switch featuresBinSlice[0] {
	case 0:
		planetType = NormalPlanet
	case 1:
		planetType = StarPlanet

	default:
		return nil, fmt.Errorf("unsupported value")
	}

	var StarNeutron bool
	switch featuresBinSlice[1] {
	case 0:
		StarNeutron = false
	case 1:
		StarNeutron = true

	default:
		return nil, fmt.Errorf("unsupported value")
	}

	return &PlanetFeatures{Type: planetType, StarNeutron: StarNeutron}, nil
}
