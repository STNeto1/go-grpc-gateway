package gimpl

import (
	"__user/ent"
	"__user/ent/user"
	userpb "__user/gen/pb/v1"
	"context"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type userService struct {
	userpb.UnimplementedUserServiceServer

	ent *ent.Client
}

func NewUserService(e *ent.Client) *userService {
	return &userService{
		ent: e,
	}
}

func (u *userService) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid id")
	}

	usr, err := u.ent.User.Query().Where(user.ID(id)).First(ctx)
	if err != nil {
		return nil, status.Error(codes.NotFound, "user not found")
	}

	return &userpb.GetUserResponse{
		User: &userpb.User{
			Id:    usr.ID.String(),
			Name:  usr.Name,
			Email: usr.Email,
		},
	}, nil
}

func (u *userService) Login(ctx context.Context, req *userpb.LoginRequest) (*userpb.LoginResponse, error) {
	usr, err := u.ent.User.Query().Where(user.Email(req.GetEmail())).First(ctx)
	if err != nil {
		return nil, status.Error(codes.NotFound, "user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(req.Password))
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "invalid credentials")
	}

	return &userpb.LoginResponse{
		User: &userpb.User{
			Id:    usr.ID.String(),
			Name:  usr.Name,
			Email: usr.Email,
		},
	}, nil
}

func (u *userService) Register(ctx context.Context, req *userpb.RegisterRequest) (*userpb.RegisterResponse, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(req.GetPassword()), bcrypt.DefaultCost)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to hash password")
	}

	usr, err := u.ent.User.Create().
		SetName(req.GetName()).
		SetEmail(req.GetEmail()).
		SetPassword(string(hash)).
		Save(ctx)

	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create user")
	}

	return &userpb.RegisterResponse{
		User: &userpb.User{
			Id:    usr.ID.String(),
			Name:  usr.Name,
			Email: usr.Email,
		},
	}, nil
}
