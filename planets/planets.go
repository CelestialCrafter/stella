package planets

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
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

func hash() []byte {
	b := make([]byte, 4)
	timestamp := time.Now().UnixNano()
	binary.BigEndian.PutUint32(b, uint32(timestamp))
	hash := sha256.Sum256(b)
	return hash[:]
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

	newHash := hash()
	newHashInt := int64(binary.BigEndian.Uint32(newHash))
	r := rand.New(rand.NewSource(newHashInt))
	frange := frangeWrapper(r)

	values := PlanetValues{
		NormalSize:       frange(10, 20),
		NormalColor:      [3]float32{frange(0, 255), frange(0, 255), frange(0, 255)},
		StarSize:         frange(20, 30),
		StarBrightness:   frange(0, 5),
		StarNeutronColor: [3]float32{frange(0, 85), frange(0, 85), frange(115, 255)},
	}

	return Planet{
		Hash:      hex.EncodeToString(newHash),
		Directory: path.Join(cwd, modelPath),
		Features:  features,
		Values:    values,
	}
}
