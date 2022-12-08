package main

import (
	productpb "__product/gen/pb/v1"
	"__product/pkg/common/database"
	"__product/pkg/common/gimpl"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	e := database.InitDB()

	lis, err := net.Listen("tcp", "localhost:4000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	productService := gimpl.NewProductService(e)

	grpcServer := grpc.NewServer()
	productpb.RegisterProductServiceServer(grpcServer, productService)

	log.Println("product service listening on localhost:4000")
	grpcServer.Serve(lis)

}
