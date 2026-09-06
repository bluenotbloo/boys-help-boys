package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/bluenotbloo/boys-help-boys/common/nacos"
	"github.com/spf13/viper"
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

func readNacosConfig() (*config, error) {
	port, err := strconv.ParseUint(getEnv("NACOS_PORT", "8848"), 10, 64) // 解析 NACOS_PORT 环境变量为 uint64
	if err != nil {
		return nil, fmt.Errorf("invalid NACOS_PORT: %w", err)
	}

	client, err := nacos.NewClient(
		getEnv("NACOS_ADDR", "127.0.0.1"),
		port,
		os.Getenv("NACOS_NAMESPACE"),
	)
	if err != nil {
		return nil, err
	}

	content, err := client.GetConfig(
		getEnv("NACOS_DATA_ID", "config.yaml"),
		getEnv("NACOS_GROUP", "DEFAULT_GROUP"),
	)
	if err != nil {
		return nil, err
	}
	if strings.TrimSpace(content) == "" {
		return nil, fmt.Errorf("nacos config is empty")
	}

	var loaded config
	if err := yaml.Unmarshal([]byte(content), &loaded); err != nil {
		return nil, fmt.Errorf("unmarshal nacos config: %w", err)
	}
	return &loaded, nil
}

func readLocalConfig() (*config, error) {
	local := viper.New()
	local.SetConfigFile(getEnv("LOCAL_CONFIG_FILE", "config.yaml"))
	if err := local.ReadInConfig(); err != nil {
		return nil, err
	}

	var loaded config
	if err := local.Unmarshal(&loaded); err != nil {
		return nil, fmt.Errorf("unmarshal local config: %w", err)
	}
	return &loaded, nil
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func GetConfig() *config {
	return cfg
}

// 加载配置文件
func LoadConfig() {
	loaded, nacosErr := readNacosConfig()
	if nacosErr != nil {
		fmt.Printf("nacos config unavailable, fallback to local config: %v\n", nacosErr)
		localConfig, err := readLocalConfig()
		if err != nil {
			panic(fmt.Errorf("load config from nacos and local file: nacos: %v; local: %w", nacosErr, err))
		}
		loaded = localConfig
	}
	cfg = loaded
	fmt.Println("config load success")
}
