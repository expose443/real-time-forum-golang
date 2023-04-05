package config

import (
	"bytes"
	"encoding/json"
	"os"
)

type Config struct {
	DB `json:"database"`
}

type DB struct {
	Dbname string `json:"database_name"`
}

func Init(path string) (*Config, error) {
	body, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var config *Config
	err = json.NewDecoder(bytes.NewBuffer(body)).Decode(&config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
