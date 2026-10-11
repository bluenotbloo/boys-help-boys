package rpc_server

import (
	"context"
	"strconv"

	"github.com/bluenotbloo/boys-help-boys/common/logger"
	pb "github.com/bluenotbloo/boys-help-boys/proto/user/gen/user/v1"
	userService "github.com/bluenotbloo/boys-help-boys/user/internal/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
	userService *userService.UserService
}

func NewUserServer() *UserServer {
	return &UserServer{}
}

func (s *UserServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	logger.Infof("login request: username=%s", req.Username)

	// TODO: 真实业务逻辑
	if req.Username == "" || req.Password == "" {
		return nil, status.Errorf(codes.InvalidArgument, "username and password cannot be empty")
	}
	userID, err := s.userService.Login(req.Username, req.Password)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid credentials")
	}
	userId, _ := strconv.Atoi(userID)
	return &pb.LoginResponse{UserId: int64(userId)}, nil
}

// 后续补充 Register、GetUserInfo 等方法
