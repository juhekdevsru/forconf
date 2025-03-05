package config

import (
	"encoding/json"
	"os"
)

// Config - структура конфигурации
type Config struct {
	DatabaseURL string `json:"database_url"`
}

// LoadConfig загружает конфиг из config.json
func LoadConfig() (*Config, error) {
	file, err := os.Open("config/config.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	config := &Config{}
	if err := json.NewDecoder(file).Decode(config); err != nil {
		return nil, err
	}

	return config, nil
}
