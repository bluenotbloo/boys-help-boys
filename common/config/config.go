package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var cfg *config

type config struct {
	Logger loggerConfig `yaml:"logger"`
}

type loggerConfig struct {
	Level      string `yaml:"level"`
	Encoding   string `yaml:"encoding"`
	Filename   string `yaml:"filename"`
	MaxSize    int    `yaml:"max-size"`
	MaxBackups int    `yaml:"max-backups"`
	MaxAge     int    `yaml:"max-age"`
	Compress   bool   `yaml:"compress"`
	Console    bool   `yaml:"console"`
}

// 读取配置文件
func readConfig() (*config, error) {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		return nil, err
	}

	var cfg config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		panic(err)
	}
	return &cfg, nil
}

func GetConfig() *config {
	return cfg
}

// 加载配置文件
func LoadConfig() {
	config, err := readConfig()
	if err != nil {
		panic(err)
	}
	cfg = config
	fmt.Println("config load success")
	return
}
