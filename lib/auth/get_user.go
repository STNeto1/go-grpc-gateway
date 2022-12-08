package auth

import (
	userpb "__user/gen/pb/v1"
	"context"
)

func GetUser(id string, client userpb.UserServiceClient, ctx context.Context) (*userpb.User, error) {
	usr, err := client.GetUser(ctx, &userpb.GetUserRequest{
		Id: id,
	})
	
	if err != nil {
		return nil, err
	}

	return usr.User, nil
}
