package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	DB      map[string]DB `json:"db"`
	Pattern []Pattern     `json:"pattern"`
}

func Load() (*Config, error) {
	f, err := os.Open("config.json")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var cfg Config
	err = json.NewDecoder(f).Decode(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
