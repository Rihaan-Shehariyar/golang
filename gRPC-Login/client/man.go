package client

import (
	pb "auth/user/proto"
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func main() {

	conn, err := grpc.Dial("localhost:50051",grpc.WithTransportCredentials(insecure.NewCredentials()))
   if err!=nil {
	log.Fatal(err)
   }

 defer conn.Close()

 client := pb.NewLoginServiceClient(conn)

 r := gin.Default()

 r.POST("/login",func(ctx *gin.Context) {
  var req struct{
 
 Username string `json:"username"`
 Password string `json:"password"`

}
 
ctx.ShouldBindJSON(&req)

 res,err := client.Login(ctx,&pb.LoginRequest{
 Username: req.Username,
 Password: req.Password,
})

 if err!=nil {
	ctx.JSON(401,gin.H{"error" : "error"})
  return
 }

 ctx.JSON(200,res)
  
})

r.GET("/profile",func(ctx *gin.Context) {

 token := ctx.GetHeader("Authorization")

 md := metadata.New(map[string]string{
 "authorization" : token,
})

 meta := metadata.NewOutgoingContext(context.Background(),md)

 res,err := client.GetProfile(meta,&pb.ProfileRequet{})
 if err!=nil {
	ctx.JSON(400,gin.H{"error":"error"})
 }

 ctx.JSON(200,res)

})
 

}