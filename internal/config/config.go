package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = "/.gatorconfig.json"

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func ReadConfig() (Config, error) {
	filePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, err
	}

	var config Config
	if err := json.Unmarshal(jsonData, &config); err != nil {
		return Config{}, err
	}

	return config, nil
}

func (cfg *Config) SetUser(currentUserName string) error {
	cfg.CurrentUserName = currentUserName
	return nil
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	fullPath := filepath.Join(home + configFileName)
	return fullPath, nil
}

func (cfg *Config) WriteFile() error {
	filePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	byteConfig, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, byteConfig, 0o644)
	return err
}
