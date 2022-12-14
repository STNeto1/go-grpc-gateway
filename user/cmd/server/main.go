package main

import (
	userpb "__user/gen/pb/v1"
	"__user/pkg/common/database"
	"__user/pkg/common/gimpl"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	e := database.InitDB()

	lis, err := net.Listen("tcp", "localhost:3000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	userService := gimpl.NewUserService(e)

	grpcServer := grpc.NewServer()
	userpb.RegisterUserServiceServer(grpcServer, userService)

	log.Println("user service listening on localhost:3000")
	grpcServer.Serve(lis)

}
