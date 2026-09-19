package main

import (
	"github.com/bluenotbloo/boys-help-boys/common/logger"
	"github.com/bluenotbloo/boys-help-boys/gateway/config"
	"github.com/bluenotbloo/boys-help-boys/gateway/internal/router"
	"github.com/gin-gonic/gin"
)

func main() {
	// 读取配置文件
	config.LoadConfig()
	// 初始化日志器
	logger.Init()
	defer logger.Sync()
	// 启动服务
	run()
	defer panicCatch()
}

func run() {
	cfg := config.GetConfig()
	logger.Infof("config: %+v", cfg)
	gin := gin.Default()
	// 初始化路由
	router.InitRouter(gin)
	// 启动服务，监听 8081 端口
	err := gin.Run(cfg.Server.Host + ":" + cfg.Server.Port)
	if err != nil {
		logger.Errorf("failed to run server: %v", err)
	}
	logger.Infof("server is running on %s:%s", cfg.Server.Host, cfg.Server.Port)
}

func panicCatch() {
	if r := recover(); r != nil {
		logger.Errorf("panic catch %v", r)
	}
}
