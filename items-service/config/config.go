package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Port  string `yaml:"port"`
	DBurl string `yaml:"db_url"`
}

var cfg *Config

func Load() *Config {
	if cfg != nil {
		return cfg
	}

	f, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("cannot read config: %v", err)
	}

	cfg = &Config{}
	if err := yaml.Unmarshal(f, cfg); err != nil {
		log.Fatalf("cannto parse config: %v", err)
	}

	return cfg
}

func Get() *Config {
	if cfg == nil {
		return Load()
	}

	return cfg
}
