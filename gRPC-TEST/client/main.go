package main

import (
	"context"
	"log"
	pb "test/proto/user"
	"time"

	"google.golang.org/grpc"
)


// func main() {

// 	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer conn.Close()

// 	client := pb.NewHelloServiceClient(conn)

// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

// 	defer cancel()

// 	res, err := client.SayHello(ctx, &pb.HelloRequest{
// 		Name: "Rihaan",
// 	})

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	log.Println("Response:", res.Name)

// }

func main() {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	client := pb.NewHelloServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	CreateUser, _ := client.CreateUser(ctx, &pb.CreateUserRequest{
		Name:  "Rihaan",
		Email: "rihaan@gmail.com",
	})

	log.Println("Created:", CreateUser)

	getUser, _ := client.GetUser(ctx, &pb.GetUserRequest{
		Id: CreateUser.User.Id,
	})

	log.Println("Fetched :", getUser.User)

	listUser, _ := client.ListUsers(ctx, &pb.Empty{})

	log.Println("All Users :", listUser)
}

// func main(){
 
//  conn,err := grpc.Dial("localhost:50051",grpc.WithInsecure())
//  if err!=nil {
// 	log.Fatal(err)
//  }

//  defer conn.Close()

//  client := pb.NewHelloServiceClient(conn)

//  ctx,cancel := context.WithTimeout(context.Background(),time.Second)
//  defer cancel()

// CreateUser,_ := client.CreateUser(ctx,&pb.CreateUserRequest{
//  Name: "",
// })
 
 
// }