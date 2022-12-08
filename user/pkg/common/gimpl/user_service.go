package gimpl

import (
	userpb "__user/gen/pb/user/v1"
	"context"
)

type userService struct {
	userpb.UnimplementedUserServiceServer
}

func NewUserService() *userService {
	return &userService{}
}

func (u *userService) GetUser(_ context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	return &userpb.GetUserResponse{
		User: &userpb.User{
			Id:    req.GetId(),
			Name:  "some name",
			Email: "some email",
		},
	}, nil
}

func (u *userService) Login(_ context.Context, req *userpb.LoginRequest) (*userpb.LoginResponse, error) {
	return &userpb.LoginResponse{
		Token: "some token",
	}, nil
}

func (u *userService) Register(_ context.Context, req *userpb.RegisterRequest) (*userpb.RegisterResponse, error) {
	return &userpb.RegisterResponse{
		Token: "some token",
	}, nil
}
