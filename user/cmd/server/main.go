package main

import (
	userpb "__user/gen/pb/user/v1"
	"__user/pkg/common/gimpl"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:3000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	userService := gimpl.NewUserService()

	grpcServer := grpc.NewServer()
	userpb.RegisterUserServiceServer(grpcServer, userService)

	log.Println("user service listening on localhost:3000")
	grpcServer.Serve(lis)

}
