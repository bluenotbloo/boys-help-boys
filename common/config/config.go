package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/bluenotbloo/boys-help-boys/common/nacos"
)

// 从 Nacos 获取配置文件
func readNacosConfig() ([]byte, error) {
	// 从环境变量获取 Nacos 配置
	port, err := strconv.ParseUint(getEnv("NACOS_PORT", "8848"), 10, 64) // 解析 NACOS_PORT 环境变量为 uint64
	if err != nil {
		return nil, fmt.Errorf("invalid NACOS_PORT: %w", err)
	}
	addr := getEnv("NACOS_ADDR", "127.0.0.1")
	namespace := getEnv("NACOS_NAMESPACE", "")
	dataId := getEnv("NACOS_DATA_ID", "")
	group := getEnv("NACOS_GROUP", "")

	// 创建 Nacos 客户端
	client, err := nacos.NewClient(
		addr,
		port,
		namespace,
	)
	if err != nil {
		return nil, err
	}
	// 从 Nacos 获取配置内容
	content, err := client.GetConfig(
		dataId,
		group,
	)
	if err != nil {
		return nil, err
	}

	return []byte(content), nil
}

// 从本地文件读取配置
func readLocalConfig() ([]byte, error) {
	local_file_path := getEnv("LOCAL_CONFIG_FILE", "config.yaml")

	loaded_config, err := os.ReadFile(local_file_path)
	if err != nil {
		return nil, fmt.Errorf("failed to read local config file: %w", err)
	}

	return loaded_config, nil
}

// getEnv 获取环境变量，如果不存在则返回默认值
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

// 加载配置文件，优先从 Nacos 获取，如果失败则从本地文件读取
func LoadConfig() ([]byte, error) {
	loaded, nacosErr := readNacosConfig()
	if nacosErr != nil {
		fmt.Printf("nacos config unavailable, fallback to local config: %v\n", nacosErr)
		localConfig, err := readLocalConfig()
		if err != nil {
			return nil, fmt.Errorf("load config from nacos and local file: nacos: %v; local: %w", nacosErr, err)
		}
		loaded = localConfig
	}
	fmt.Println("config load success")
	return loaded, nil
}
