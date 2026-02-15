package middleware

import "github.com/gin-gonic/gin"


func AdminOnly() gin.HandlerFunc{
return func(ctx *gin.Context) {

  role,exist := ctx.Get("role")
 if !exist || role != "admin"{
  ctx.JSON(401,gin.H{"error":"Admin only"})
  ctx.Abort()
 return
  }

 ctx.Next()

}

}




