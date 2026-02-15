package middleware

import (
	"admin/utility"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// import (
// 	"admin/utility"
// 	"strings"

// 	"github.com/gin-gonic/gin"
// 	"github.com/golang-jwt/jwt/v5"
// )

// func JwtMiddleWare() gin.HandlerFunc{
// return func(ctx *gin.Context) {

//   auth:=ctx.GetHeader("Authorization")
//  if auth==""{
//    ctx.JSON(401,gin.H{"error":"Missing Token"})
//   ctx.Abort()
//   return
// }

//  tokenstr := strings.TrimPrefix(auth,"Bearer ")

//  token,err:=jwt.Parse(tokenstr,func(t *jwt.Token) (interface{}, error) {
//       return utility.JwtSecret,nil
// })

//  if err!=nil|| !token.Valid{
//   ctx.JSON(401,gin.H{"error":"Invalid token"})
//  ctx.Abort()
// return
// }

//  claims:=token.Claims.(*utility.JwtClaims)

//  ctx.Set("user_id",claims.UserId)
//  ctx.Set("role",claims.Role)

//  ctx.Next()

// }
// }



func JwtMiddleWare()gin.HandlerFunc{
return func(ctx *gin.Context) {

 auth:=ctx.GetHeader("Authorization") 
 if auth==""{
 ctx.JSON(401,gin.H{"error":"Invalid input"})
 ctx.Abort()
 return 
}

 tokenstr:=strings.TrimPrefix(auth,"Bearer ")
 
token,err:=jwt.Parse(tokenstr,func(t *jwt.Token) (interface{}, error) {
  return utility.JwtSecret,nil
})

 if err!=nil || !token.Valid{
 ctx.JSON(402,gin.H{"error":"error"})
 ctx.Abort()
 return 
}

 claims := token.Claims.(*utility.JwtClaims)
 ctx.Set("user_id",claims.UserId)
 ctx.Set("role",claims.Role)
 

}
}