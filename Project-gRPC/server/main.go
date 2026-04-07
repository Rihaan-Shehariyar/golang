package main

import (
	pb "auth/auth/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(AuthInterceptor))

	pb.RegisterAuthServiceServer(grpcServer, &server{})

	log.Println("Server running on 50051:")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}

}
