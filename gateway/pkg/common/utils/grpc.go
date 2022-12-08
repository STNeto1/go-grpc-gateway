package utils

import (
	userpb "__user/gen/pb/user/v1"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitGrpcConn() *grpc.ClientConn {

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.Dial("localhost:3000", opts...)
	if err != nil {
		log.Fatalln("Failed to dial user service: ", err)
	}

	return conn
}

func InitGrpcUserClient(conn *grpc.ClientConn) userpb.UserServiceClient {
	client := userpb.NewUserServiceClient(conn)

	return client
}
