package config

import (
	"fmt"

	"github.com/bluenotbloo/boys-help-boys/common/config"
	"gopkg.in/yaml.v3"
)

var cfg *Config

type Config struct {
	Server ServerConfig `yaml:"user"`
}

type ServerConfig struct {
	Port     string `yaml:"port"`
	Host     string `yaml:"host"`
	GrpcPort string `yaml:"grpc_port"`
}

func GetConfig() *Config {
	return cfg
}

func LoadConfig() {
	loaded, err := config.LoadConfig()
	if err != nil {
		panic(fmt.Errorf("failed to load config: %w", err))
	}
	cfg = &Config{}
	if err := yaml.Unmarshal(loaded, cfg); err != nil {
		panic(fmt.Errorf("failed to unmarshal config: %w", err))
	}
}
