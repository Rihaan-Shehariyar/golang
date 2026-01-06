package main

import (


	"github.com/gin-gonic/gin"
)

// func authMiddleware()gin.HandlerFunc{
//  return func(ctx *gin.Context) {
//    if user.name ==""{
//    ctx.JSON(http.StatusBadRequest,gin.H{
//    "error" : "empty box",
// })
// ctx.Abort()
// return 
// }
// }
 
// }


func main(){
 
 r := gin.Default()

 r.POST("/login",func(ctx *gin.Context) {
   ctx.JSON(200,gin.H{
  "message" : "Hello",
})
})

 r.Run(":8080")
}