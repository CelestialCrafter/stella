package planets

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
)

const (
	blenderExe          = "org.blender.Blender"
	scriptErrorExitCode = 73
	scriptPath          = "blender/app.py"
	stellaPrefix        = "[stella]"
	modelPath           = "models/"
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
		Size:          3.25,
		Color:         [3]float32{255, 150, 255},
		BlotchDensity: 3.25,
	}

	return Planet{
		Hash:     hash,
		FilePath: modelPath,
		Features: features,
		Values:   values,
	}
}

func (p Planet) CreateModel() error {
	marshalled, err := json.Marshal(p)
	if err != nil {
		return err
	}

	// @PERF turn the blender script into a daemon so startup isnt required on every model creation
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	absoluteScriptPath := path.Join(cwd, scriptPath)

	blenderCmd := exec.Command(
		blenderExe,
		"--background",
		"--python-use-system-env",
		"--python-exit-code", fmt.Sprint(scriptErrorExitCode),
		"--python", absoluteScriptPath,
		"--", string(marshalled),
	)

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	blenderCmd.Stdout = &stdout
	blenderCmd.Stderr = &stderr

	err = blenderCmd.Run()
	exitErr, ok := err.(*exec.ExitError)
	if ok {
		return fmt.Errorf("%s: %s", exitErr, stderr.String())
	}
	if err != nil {
		return err
	}

	output := ""
	for _, line := range strings.Split(stdout.String(), "\n") {
		if !strings.HasPrefix(line, stellaPrefix) {
			continue
		}

		output += fmt.Sprint(strings.TrimPrefix(line, stellaPrefix), "\n")
	}

	fmt.Println(output)

	return nil
}
