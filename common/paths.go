package common

import (
	"os"
	"path"
)

var ModelPath = func() string {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return path.Join(cwd, "models/")
}()
