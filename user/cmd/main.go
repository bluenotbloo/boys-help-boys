package main

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/bluenotbloo/boys-help-boys/common/logger"
	pb "github.com/bluenotbloo/boys-help-boys/proto/user/gen/user/v1"
	"github.com/bluenotbloo/boys-help-boys/user/config"
	rpcserver "github.com/bluenotbloo/boys-help-boys/user/internal/rpc_server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// 读取配置文件
	config.LoadConfig()
	// 初始化日志器
	logger.Init()
	defer logger.Sync()
	// 启动grpc服务
	runGRPC()
}

func runGRPC() {
	cfg := config.GetConfig()
	logger.Infof("user config: %+v", cfg)

	listen, err := net.Listen("tcp", cfg.Server.Host+":"+cfg.Server.GrpcPort)
	if err != nil {
		logger.Errorf("listen tcp failed, prot:%s,error:%s", cfg.Server.GrpcPort, err)
		panic(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, rpcserver.NewUserServer())
	reflection.Register(grpcServer) // 支持 grpcurl 调试

	// 优雅启停
	go func() {
		logger.Infof("user gRPC server listening on %s", cfg.Server.GrpcPort)
		err = grpcServer.Serve(listen)
		if err != nil {
			logger.Errorf("user rpc server start failed, error:%s", err)
			return
		}
	}()

	// 等待中断信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Infof("shutting down server...")

	// 优雅停止：停止接受新连接，等待进行中请求完成
	grpcServer.GracefulStop()
	logger.Infof("server exited")

}
