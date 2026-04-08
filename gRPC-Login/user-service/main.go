package main

import (
	pb "auth/user/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", "50051")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(AuthInterceptor))

	pb.RegisterLoginServiceServer(grpcServer, server{})

	grpcServer.Serve(lis)

}
