package config

import (
	"fmt"
	"github.com/bluenotbloo/boys-help-boys/common/config"
	"gopkg.in/yaml.v3"
)

var cfg *Config

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Logger   LoggerConfig   `yaml:"logger"`
	JWT      JWTConfig      `yaml:"jwt"`
	Jaeger   JaegerConfig   `yaml:"jaeger"`
	Upstream UpstreamConfig `yaml:"upstream"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
	Host string `yaml:"host"`
}

type LoggerConfig struct {
	Level      string `yaml:"level"`
	Encoding   string `yaml:"encoding"`
	Filename   string `yaml:"filename"`
	MaxSize    int    `yaml:"max-size"`
	MaxBackups int    `yaml:"max-backups"`
	MaxAge     int    `yaml:"max-age"`
	Compress   bool   `yaml:"compress"`
	Console    bool   `yaml:"console"`
}

type JWTConfig struct {
	Secret    string `yaml:"secret"`
	Issuer    string `yaml:"issuer"`
	ExpiresIn string `yaml:"expires_in"`
}

type JaegerConfig struct {
	AgentHost    string  `yaml:"agent_host"`
	AgentPort    int     `yaml:"agent_port"`
	SamplerParam float64 `yaml:"sampler_param"`
	ServiceName  string  `yaml:"service_name"`
}

type UpstreamConfig struct {
	UserService string `yaml:"user_service"`
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