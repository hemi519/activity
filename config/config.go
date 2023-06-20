package config

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

type Config struct {
	FetchInterval string `json:"fetch_interval"`
	Output        string `json:"output"`
	BoredAPI      string `json:"bored_api"`
}

func NewConfig(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	err = json.Unmarshal(data, cfg)
	if err != nil {
		return nil, err
	}

	cfg.ParseFetchInterval()

	return cfg, nil
}

func (cfg *Config) ParseFetchInterval() {
	duration, err := time.ParseDuration(cfg.FetchInterval)
	if err == nil {
		cfg.FetchInterval = duration.String()
	}
}
