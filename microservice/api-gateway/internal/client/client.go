package client

import (
	"log"
	orderpb "shared-proto/order"
	userpb "shared-proto/user"

	"google.golang.org/grpc"
)

type Clients struct {
	UserClient  userpb.UserServiceClient
	OrderClient orderpb.OrderServiceClient
}

func NewClients() *Clients {
	userConn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	orderConn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	return &Clients{
		UserClient:  userpb.NewUserServiceClient(userConn),
		OrderClient: orderpb.NewOrderServiceClient(orderConn),
	}

}
