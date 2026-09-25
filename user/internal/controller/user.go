package controller

import (
	"context"

	"github.com/bluenotbloo/boys-help-boys/user/internal/dto"
	"github.com/bluenotbloo/boys-help-boys/user/internal/service"
)

type UserController struct {
	Service *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		Service: service.NewUserService(),
	}
}

func (u *UserController) Login(ctx context.Context, req *dto.UserLoginRequest) *dto.Response {
	resp := &dto.Response{
		Code:    500,
		Message: "success",
	}
	// TODO: 实现登录逻辑，例如验证用户名和密码，生成 JWT token 等
	username := req.Username
	passwd := req.Password

	userId, loginErr := u.Service.Login(username, passwd)
	if loginErr != nil {
		resp.Code = 403
		resp.Message = "user login faild"
	}
	resp.Data = &dto.UserLoginResponse{UserID: userId}
	return resp
}

func (u *UserController) Register(ctx context.Context, req *dto.UserRegisterRequest) *dto.Response {
	// TODO: 实现注册逻辑，例如验证用户信息，创建新用户等

	userResp := &dto.UserRegisterResponse{
		UserID: 12345,
	}

	resp := &dto.Response{
		Code:    200,
		Message: "Register success",
		Data:    userResp,
	}
	return resp
}
