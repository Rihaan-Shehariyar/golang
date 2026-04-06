package main

import (
	"context"
	"log"
	"net"
	pb "test/proto/user"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedHelloServiceServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {

	log.Println("Recieved:", req.Name)

	return &pb.HelloResponse{
		Name: req.Name,
	}, nil

}

func main() {

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err.Error())
	}

	grpcServer := grpc.NewServer()

 pb.RegisterHelloServiceServer(grpcServer,&server{})
 log.Println("Server Running on port : 50051")
 if err := grpcServer.Serve(lis);err!=nil{
  log.Fatal(err)
}

}
