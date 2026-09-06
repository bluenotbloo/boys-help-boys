package main

import (
	"github.com/bluenotbloo/boys-help-boys/common/config"
	"github.com/bluenotbloo/boys-help-boys/common/logger"
)

func main() {
	// 读取配置文件
	config.LoadConfig()
	// 初始化日志器
	logger.Init()
	defer logger.Sync()
	defer panicCatch()
}

func panicCatch() {
	if r := recover(); r != nil {
		logger.Errorf("panic catch %v", r)
	}
}
