package utils

import (
	"hash"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// import (

// 	"time"

// 	"github.com/golang-jwt/jwt/v5"
// )

//  var jwtSecret = []byte("super-secret-key")

// func Accestoken(userId uint) (string, error) {

// 	claims := jwt.MapClaims{
//    "user_id" : userId,
//    "exp" : time.Now().Add(15*time.Minute).Unix(),

// }

//  token:=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
//  return  token.SignedString(jwtSecret)

// }

// func RefreshToken(userId uint)(string ,error){

//   claims:= jwt.MapClaims{
//   "user_id" : userId,
//    "exp" : time.Now().Add(7 * 24 * time.Hour).Unix(),
//  }

//  token:=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
//  return  token.SignedString(jwtSecret)

// }

// func ParseToken(tokenStr string)(jwt.MapClaims,error){
//   token,err:=jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
//    return jwtSecret,nil

// })

//  if err!=nil || token.Valid{
//   return nil,err
// }

//  return token.Claims.(jwt.MapClaims),nil

// }

// var jwtSecret = []byte("super-secret-key")

// func Accestoken(userId uint)(string,error){

// claims:=jwt.MapClaims{

//   "user_id" : userId,
//   "exp" : time.Now().Add(15*time.Minute).Unix(),

// }

// token:=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
// return token.SignedString(jwtSecret)

// }

// func RefreshToken(userId uint)(string,error){

//   claioms:=jwt.MapClaims{

//  "user_id" : userId,
//  "exp" : time.Now().Add(7*24 *time.Hour).Unix(),

// }

//  token:=jwt.NewWithClaims(jwt.SigningMethodHS256,claioms)
//  return token.SignedString(jwtSecret)

// }

// func ParseToken(tokenstr string)(jwt.MapClaims,error){

//  token,err := jwt.Parse(tokenstr,func(t *jwt.Token) (any, error) {

//     if _,ok := t.Method.(*jwt.SigningMethodHMAC);!ok{
//    return nil,errors.New("unexpected signing method")
// }

//  return jwtSecret,nil

// })

//  if err!=nil || !token.Valid{
//  return nil,errors.New("Invalid or expired token")

// }

//  claims,ok := token.Claims.(jwt.MapClaims)

//  if !ok{
//  return nil,errors.New("Invalid token claims")
// }

//  return claims,nil

// }



type Claims struct{
 UserId uint
 email string
 role string
 jwt.RegisteredClaims
}
var jwtSecret = []byte("super-secret-key")
func GenereateAcessToken(userID uint,email,role string)(string,error){



 claims:=Claims{
  UserId: userID,
  email: email,
  role: role ,
  RegisteredClaims: jwt.RegisteredClaims{
   ExpiresAt: jwt.NewNumericDate(time.Now().Add(15*time.Minute)),
},
}

 token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
 return token.SignedString([]byte(jwtSecret))

}


func JwtAuth()gin.HandlerFunc{
 return func(ctx *gin.Context) {
 
  auth := ctx.GetHeader("Authorization")
  if auth == ""{
   ctx.JSON(400,gin.H{"error":"Token Missing"})
   ctx.Abort()
   return 
}

 tokenStr := strings.TrimPrefix(auth,"Bearer ")
 claims := &Claims{}
 
 token,err:= jwt.ParseWithClaims(tokenStr,claims,func(t *jwt.Token) (interface{}, error) {
      return jwtSecret,nil
})

 if err!=nil{
   ctx.JSON(400,gin.H{"error":"Invalid token"})
  ctx.Abort()
  return 
}

 

 

}
}

