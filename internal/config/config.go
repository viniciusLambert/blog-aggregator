package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	dbURL           string `json:"db_url"`
	currentUserName string `json:"current_user_name"`
}

const configFileName = ".gatorcofing.json"

func readConfig() (Config, error) {
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

func (cfg *Config) setUser(currentUserName string) error {
	cfg.currentUserName = currentUserName
	return nil
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	filePath := home + configFileName
	return filePath, nil
}

func writeConfig(cfg Config) error {
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
