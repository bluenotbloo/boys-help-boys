package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/bluenotbloo/boys-help-boys/common/logger"
	pb "github.com/bluenotbloo/boys-help-boys/proto/user/gen/user/v1"
	"github.com/bluenotbloo/boys-help-boys/user/config"
	rpcserver "github.com/bluenotbloo/boys-help-boys/user/internal/rpc_server"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func main() {
	// 读取配置文件
	config.LoadConfig()
	// 初始化日志器
	logger.Init()
	defer logger.Sync()

	cfg := config.GetConfig()
	grpcAddr := net.JoinHostPort(cfg.Server.Host, cfg.Server.GrpcPort)
	httpAddr := net.JoinHostPort(cfg.Server.Host, cfg.Server.Port)

	// 启动 gRPC 服务
	go runGRPC(grpcAddr)
	// 启动 grpc-gateway HTTP 服务
	go runRPCGateway(grpcAddr, httpAddr)

	// 等待中断信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Infof("shutting down server...")
}

// runGRPC 启动 gRPC 服务
func runGRPC(addr string) {
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		logger.Errorf("listen tcp failed, addr:%s, error:%s", addr, err)
		panic(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, rpcserver.NewUserServer())
	reflection.Register(grpcServer) // 支持 grpcurl 调试

	logger.Infof("user gRPC server listening on %s", addr)
	if err := grpcServer.Serve(listen); err != nil {
		logger.Errorf("user rpc server start failed, error:%s", err)
	}
}

// runRPCGateway 启动 grpc-gateway HTTP 服务，反向代理到本服务 gRPC 端口
func runRPCGateway(grpcAddr, httpAddr string) {
	ctx := context.Background()

	// 创建 grpc-gateway 的 mux（HTTP → gRPC 反向代理）
	mux := runtime.NewServeMux()

	// 注册 user 服务的 HTTP 路由，转发到本服务 gRPC 端口
	err := pb.RegisterUserServiceHandlerFromEndpoint(
		ctx,
		mux,
		grpcAddr,
		[]grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())},
	)
	if err != nil {
		logger.Fatalf("register user service handler failed: %v", err)
	}

	logger.Infof("user grpc-gateway HTTP API listening on %s", httpAddr)
	if err := http.ListenAndServe(httpAddr, mux); err != nil && err != http.ErrServerClosed {
		logger.Errorf("user grpc-gateway http server error: %v", err)
	}
}
