package main

import (
	"context"
	"log"
	"net"

	pb "grpc/user/proto"

	"google.golang.org/grpc"
)

// server struct implements the gRPC service
type server struct {
	pb.UnimplementedHelloServiceServer
}

// SayHello is the RPC method implementation
func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Println("Received:", req.Name)

	return &pb.HelloResponse{
		Message: "Hello " + req.Name,
	}, nil
}

func main() {
	// 1. Create TCP listener
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("Failed to listen:", err)
	}

	// 2. Create gRPC server
	grpcServer := grpc.NewServer()

	// 3. Register service
	pb.RegisterHelloServiceServer(grpcServer, &server{})

	log.Println("🚀 Server running on port 50051...")

	// 4. Start server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Failed to serve:", err)
	}
}