package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	BotToken string `json:"botToken"`
}

var config Config

// Read reads the configuration from a JSON file.
// If the path is nil, it assumes the default path "projectRoot/tellyfin.json".
// It returns the Config struct with the configuration data.
func Read(path *string) Config {
	if path == nil {
		// Get the current working directory
		cwd, err := os.Getwd()
		if err != nil {
			log.Fatal(fmt.Errorf("fail to get current working directory: %w", err))
		}
		// Construct the absolute path by joining the current working directory and the relative path
		absPath := filepath.Join(cwd, "./tellyfin.json")
		path = &absPath

	}
	// Read the JSON file
	file, err := os.ReadFile(*path)
	if err != nil {
		log.Fatal(fmt.Errorf("fail to read config file: %w", err))
	}
	// Unmarshal the JSON data into the Config struct
	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Fatal(fmt.Errorf("fail to unmarshal config file: %w", err))
	}

	return config
}
func Get() Config {
	return config
}
