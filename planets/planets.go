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
	BLENDER_EXE            = "org.blender.Blender"
	SCRIPT_ERROR_EXIT_CODE = 73
	SCRIPT_PATH            = "blender/app.py"
	STELLA_PREFIX          = "[stella]"
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
}

func (p Planet) ModelPath() string {
	return path.Join("planet-models", p.Hash+".obj")
}

func (p Planet) CreateModel() error {
	marshalled, err := json.Marshal(p)
	if err != nil {
		return err
	}

	// @FIX turn the blender script into a daemon so startup isnt required on every model creation
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	absoluteScriptPath := path.Join(cwd, SCRIPT_PATH)

	blenderCmd := exec.Command(
		BLENDER_EXE,
		"--background",
		"--python-use-system-env",
		"--python-exit-code", fmt.Sprint(SCRIPT_ERROR_EXIT_CODE),
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
		if !strings.HasPrefix(line, STELLA_PREFIX) {
			continue
		}

		output += fmt.Sprint(strings.TrimPrefix(line, STELLA_PREFIX), "\n")
	}

	fmt.Println(output)

	return nil
}
