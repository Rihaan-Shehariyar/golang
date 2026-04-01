package main

import (
	"context"
	pb "grpc/user/proto"
	"log"
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

// 	log.Println("Response:", res.Message)
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

	res, err := client.SayHello(ctx, &pb.HelloRequest{
		Name: "Abu sala",
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Response : ", res.Message)

}
