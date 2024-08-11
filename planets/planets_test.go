package planets

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"os"
	"path"
	"reflect"
	"testing"
)

func TestPlanet(t *testing.T) {
	features := PlanetFeatures{
		Type:        "normal",
		StarNeutron: false,
		NormalRings: true,
	}

	hashString := "7465737420706c616e65742068617368"
	hash, err := hex.DecodeString(hashString)
	if err != nil {
		panic(err)
	}

	planet := NewPlanet(features, hash)
	planet.Directory = ""

	want := Planet{
		Features: features,
		Values: PlanetValues{
			12.9511795,
			[3]float32{56.581818, 172.14716, 198.64021},
			2,
			[][3]float32{{53.745842, 179.36815, 133.64684}, {209.60231, 254.40836, 201.55669}},
			[][3]float32{{293.34973, 123.0839, 50.589245}, {278.69116, 325.73477, 140.07175}},
			4.6119537,
			0.9455631,
			22.177917,
			[3]float32{28.532635, 23.233162, 176.6917},
		},
		Hash: hashString,
	}

	if !reflect.DeepEqual(planet, want) {
		t.Errorf("planet with hash %s did not match expected planet", hash)
	}
}

func TestModel(t *testing.T) {
	features := PlanetFeatures{
		Type:        "normal",
		StarNeutron: false,
		NormalRings: true,
	}

	hashString := "7465737420706c616e65742068617368"
	hash, err := hex.DecodeString(hashString)
	if err != nil {
		panic(err)
	}

	planet := NewPlanet(features, hash)
	err = planet.CreateModel()
	if err != nil {
		panic(fmt.Errorf("could not create model: %w", err))
	}

	generated, err := os.ReadFile(path.Join(planet.Directory, planet.Hash+".glb"))
	if err != nil {
		panic(fmt.Errorf("could not open generated model file: %w", err))
	}

	want, err := os.ReadFile("test.glb")
	if err != nil {
		panic(fmt.Errorf("could not open generated model file: %w", err))
	}

	if !bytes.Equal(generated, want) {
		t.Errorf("generated model did not match expected model")
	}
}
