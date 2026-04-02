package main

import (
	"log"
	"net"
	"user_system/internal/db"
	"user_system/internal/handler"
	"user_system/internal/repository"
	"user_system/internal/service"
	pb "user_system/user/proto"

	"google.golang.org/grpc"
)

func main() {

	database, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewUserRepository(database)
	svc := service.NewUserRepository(repo)
	h := handler.NewUserHandler(svc)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, h)

	log.Println("Server Running on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}

}
