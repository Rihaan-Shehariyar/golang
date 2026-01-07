package middlewares

import (
	"jwt/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc{
 return func(ctx *gin.Context) {
   header:=ctx.GetHeader("Authorization")
 if header==""{
  ctx.JSON(http.StatusUnauthorized,gin.H{
   "error" : "missing token",
})

ctx.Abort()
return 
}  
 
 tokenStr := strings.Split(header, "")[1]
 claims,err := utils.ParseToken(tokenStr)
 if err!=nil{
   ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			ctx.Abort()
			return
}
   
   ctx.Set("user_id",uint(claims["user_id"].(float64)))
   ctx.Next()
}

}