package main

import (
	pb "auth/auth/proto"
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	client := pb.NewAuthServiceClient(conn)

	r := gin.Default()

	r.POST("/login", func(ctx *gin.Context) {
		var req struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		ctx.BindJSON(&req)

		res, err := client.Login(context.Background(), &pb.LoginRequest{
			Username: req.Username,
			Password: req.Password,
		})
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, res)

	})

	r.GET("/profile", func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		println(token)

		c := metadataContext(token)

		res, err := client.GetProfile(c, &pb.ProfileRequest{})
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, res)

	})
	r.Run(":8080")

}

func metadataContext(token string) context.Context {
	md := metadata.Pairs("authorization", token)
	return metadata.NewOutgoingContext(context.Background(), md)
}