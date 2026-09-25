package service

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (*UserService) Login(username, password string) (string, error) {
	return "1", nil
}
