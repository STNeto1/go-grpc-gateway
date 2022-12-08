package utils

import (
	productpb "__product/gen/pb/v1"
	userpb "__user/gen/pb/v1"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitGrpcUserClient() (*grpc.ClientConn, userpb.UserServiceClient) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.Dial("localhost:3000", opts...)
	if err != nil {
		log.Fatalln("Failed to dial user service: ", err)
	}

	client := userpb.NewUserServiceClient(conn)

	return conn, client
}

func InitGrpcProductClient() (*grpc.ClientConn, productpb.ProductServiceClient) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.Dial("localhost:4000", opts...)
	if err != nil {
		log.Fatalln("Failed to dial product service: ", err)
	}

	client := productpb.NewProductServiceClient(conn)

	return conn, client
}
