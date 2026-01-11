package middleware

import (
	"admin/utility"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JwtAuthMiddleware() gin.HandlerFunc{
 
  return func(ctx *gin.Context) {

  auth := ctx.GetHeader("Authorization")
  if auth==""{
   ctx.JSON(http.StatusUnauthorized,gin.H{"error":"Token missing"})
   ctx.Abort()
    return 

}

 tokenString:=strings.TrimPrefix(auth,"Bearer ")
 
 token,err:=jwt.ParseWithClaims(tokenString,utility.JwtClaims{},func(t *jwt.Token) (interface{}, error) {

  return utility.JwtSeecret,nil
})

 if err!=nil||!token.Valid{
  
   ctx.JSON(http.StatusUnauthorized,gin.H{"error":"Invalid Token"})
   ctx.Abort()
    return 

}

claims:=token.Claims.(*utility.JwtClaims)

ctx.Set("user_id",claims.UserId)
ctx.Set("email",claims.Email)
ctx.Set("role",claims.Role)

ctx.Next()

}

}