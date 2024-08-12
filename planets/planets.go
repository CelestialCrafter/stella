package planets

import (
	"encoding/binary"
	"encoding/hex"
	"math/rand"
	"os"
	"path"

	"github.com/CelestialCrafter/stella/common"
)

type PlanetFeatures struct {
	Type           string `json:"type" validate:"required"`
	StarNeutron    bool   `json:"star_neutron"`
	NormalRings    bool   `json:"normal_rings"`
	BlackholeStyle string `json:"blackhole_style"`
}

type PlanetValues struct {
	NormalSize         float32       `json:"normal_size"`
	NormalColor        [3]float32    `json:"normal_color"`
	NormalRingAmount   int           `json:"normal_ring_amount"`
	NormalRingColors   [][3]float32  `json:"normal_ring_colors"`
	NormalRingRotation [][3]float32  `json:"normal_ring_rotation"`
	NormalRingSize     float32       `json:"normal_ring_size"`
	BlackHoleSize      float32       `json:"blackhole_size"`
	BlackHoleColors    [3][3]float32 `json:"blackhole_colors"`
	StarBrightness     float32       `json:"star_brightness"`
	StarSize           float32       `json:"star_size"`
	StarNeutronColor   [3]float32    `json:"star_neutron_color"`
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

func NewPlanet(features PlanetFeatures, newHash []byte) Planet {
	if newHash == nil {
		newHash = common.Hash()
	}
	newHashInt := int64(binary.BigEndian.Uint32(newHash))
	r := rand.New(rand.NewSource(newHashInt))
	frange := frangeWrapper(r)

	values := PlanetValues{
		NormalSize:  frange(1, 5),
		NormalColor: [3]float32{frange(0, 255), frange(0, 255), frange(0, 255)},
		NormalRingRotation: [][3]float32{
			{frange(0, 360), frange(0, 360), frange(0, 360)},
			{frange(0, 360), frange(0, 360), frange(0, 360)},
		},
		NormalRingColors: [][3]float32{
			{frange(0, 255), frange(0, 255), frange(0, 255)},
			{frange(0, 255), frange(0, 255), frange(0, 255)},
		},
		NormalRingSize:   frange(1, 2),
		NormalRingAmount: r.Intn(2) + 1,
		BlackHoleSize:    frange(1, 10),
		BlackHoleColors: [3][3]float32{
			{frange(0, 255), frange(0, 255), frange(0, 255)},
			{frange(0, 255), frange(0, 255), frange(0, 255)},
			{frange(0, 255), frange(0, 255), frange(0, 255)},
		},
		StarSize:         frange(5, 10),
		StarBrightness:   frange(0, 5),
		StarNeutronColor: [3]float32{frange(15, 30), frange(20, 40), frange(150, 255)},
	}

	return Planet{
		Hash:      hex.EncodeToString(newHash),
		Directory: path.Join(os.Getenv("BLENDER_DATA_PATH"), "models/"),
		Features:  features,
		Values:    values,
	}
}
