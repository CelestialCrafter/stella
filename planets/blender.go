package planets

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"

	"github.com/charmbracelet/log"
)

const (
	scriptErrorExitCode = 73
	stellaPrefix        = "[stella]"
)

func (p Planet) CreateModel() error {
	start := time.Now()
	marshalled, err := json.Marshal(p)
	if err != nil {
		return err
	}

	// @PERF turn the blender script into a daemon so startup isnt required on every model creation
	blenderCmd := exec.Command(
		os.Getenv("BLENDER_EXE"),
		"--background",
		"--python-use-system-env",
		"--python-exit-code", fmt.Sprint(scriptErrorExitCode),
		"--python", path.Join(os.Getenv("BLENDER_DATA_PATH"), "app.py"),
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

	for _, line := range strings.Split(stdout.String(), "\n") {
		if !strings.HasPrefix(line, stellaPrefix) {
			continue
		}

		line = fmt.Sprint(strings.TrimPrefix(line, stellaPrefix), "\n")
		log.Infof("blender -> stella: %s", line)
	}

	log.Info("rendered new model", "hash", p.Hash, "duration", time.Since(start))
	return nil
}
