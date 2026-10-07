package rpc_server

import (
	"context"

	"github.com/bluenotbloo/boys-help-boys/common/logger"
	pb "github.com/bluenotbloo/boys-help-boys/proto/user/gen/user/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
}

func NewUserServer() *UserServer {
	return &UserServer{}
}

func (s *UserServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	logger.Infof("login request: username=%s", req.Username)

	// TODO: 真实业务逻辑
	if req.Username == "admin" && req.Password == "123456" {
		return &pb.LoginResponse{UserId: 1}, nil
	}
	return nil, status.Errorf(codes.Unauthenticated, "invalid credentials")
}

// 后续补充 Register、GetUserInfo 等方法
