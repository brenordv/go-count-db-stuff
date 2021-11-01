package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/brenordv/go-count-db-stuff/internal/models"
	"log"
	"os"
)

func FileExists(f string) bool {
	if _, err := os.Stat(f); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}

func ReadConfigFile(f string) *models.RuntimeConfig {
	if !FileExists(f) {
		log.Panicf("File '%s' not found!", f)
	}

	file, err := os.Open(f)
	ErrorHandler(err, fmt.Sprintf("Failed to open file '%s'.", f))

	var runtimeConfig models.RuntimeConfig
	parser := json.NewDecoder(file)
	err = parser.Decode(&runtimeConfig)
	ErrorHandler(err, fmt.Sprintf("Failed to parse file '%s'.", f))

	return &runtimeConfig
}
