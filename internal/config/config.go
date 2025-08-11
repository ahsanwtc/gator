package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func Read() (Config, error) {
	var _config Config

	path, err := getConfigPath()
	if err != nil {
		return _config, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return _config, err
	}

	err = json.Unmarshal(data, &_config)
	return _config, err
}

func (cfg *Config) SetUser(username string) error {
	cfg.CURRENT_USER = username
	path, err := getConfigPath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(cfg, "", "  ") 
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

func getConfigPath() (string, error) {
	var path string
	homePath, err := os.UserHomeDir()
	if err != nil {
		return path, fmt.Errorf("user's home directory couldn't be found")
	}

	path = filepath.Join(homePath, configFileName)
	return path, nil
}