package main

import (
	"os"
	"time"

	"github.com/charmbracelet/log"
	"github.com/grafana/pyroscope-go"
)

func startPyroscope() {
	address, ok := os.LookupEnv("PYROSCOPE_ADDRESS")
	if !ok {
		return
	}

	_, err := pyroscope.Start(pyroscope.Config{
		ApplicationName: "crawler",
		ServerAddress:   address,
		Logger:          pyroscope.StandardLogger,
		UploadRate:      5 * time.Second,
	})
	if err != nil {
		log.Fatal("unable to start pyroscope", "error", err)
	}
}
