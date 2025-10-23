// Package config for all the internal configurations
package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUsername string `json:"current_user_name"`
}

const configFileName string = ".gatorconfig.json"

// Read the configuration file from the home directory
// returns a structure with the configuration options
func Read() (Config, error) {
	var result Config

	filePath, err := getFilePath(configFileName)
	if err != nil {
		return result, err
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(data, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

// SetUser to a new username and store in config
func (c *Config) SetUser(username string) error {
	c.CurrentUsername = username
	return write(c)
}

func write(config *Config) error {
	bytes, err := json.Marshal(config)
	if err != nil {
		return err
	}

	filePath, err := getFilePath(configFileName)
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, bytes, 0o644)
}

// Constructs the file path for a specified file name
func getFilePath(fileName string) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(homeDir, fileName), nil
}
