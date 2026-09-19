package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	// 初始化路由
	r := gin.Default()
	router.InitRouter(r)

	// 创建 HTTP 服务器
	cfg := config.GetConfig()
	addr := cfg.Server.Host + ":" + cfg.Server.Port
	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	// 优雅关闭
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("listen: %s\n", err)
		}
	}()

	logger.Infof("server is running on %s", addr)

	// 等待中断信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Infof("shutting down server...")

	// 给正在处理的请求 5 秒完成
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatalf("server forced to shutdown: %v", err)
	}

	logger.Infof("server exiting")
}