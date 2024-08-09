package utils

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"

	"github.com/charmbracelet/log"
)

func GenerateInterfaceSHA256(i interface{}) string {
	// Serialize the struct to JSON
	jsonData, err := json.Marshal(i)
	if err != nil {
		log.Error("Unable to marshal json")
		return ""
	}

	hash := sha256.Sum256(jsonData)

	return fmt.Sprintf("%x", hash)
}
