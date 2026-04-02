package main

import (
	"database/sql"
	"log"
	"net"
	"order_service/internal/handler"
	"order_service/internal/repository"
	"order_service/internal/service"

	_ "github.com/lib/pq"

	pb "shared-proto/order"
	userpb "shared-proto/user"

	"google.golang.org/grpc"
)

func main() {

	connstr := "user=postgres password=Rihaan@123 dbname=ecommerce host=localhost port=5432 sslmode=disable"

	db, err := sql.Open("postgres", connstr)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	userClient := userpb.NewUserServiceClient(conn)

	repo := repository.NewOrderRepository(db)
	svc := service.NewOrderService(repo)
	h := handler.NewOrderHandler(svc, userClient)

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatal(err)
	}

	grpcserver := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcserver, h)

	log.Println("Order Service running on 50052")
	grpcserver.Serve(lis)

}
