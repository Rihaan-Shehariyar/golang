package utility

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// import (
// 	"os"
// 	"time"

// 	"github.com/golang-jwt/jwt/v5"
// )

var  JwtSecret = []byte(os.Getenv("secret"))

// type JwtClaims struct {
// 	UserId uint
// 	Role   string
// 	jwt.RegisteredClaims
// }

// func AcesssToken(userID uint,role string)(string,error){

//   claims:= JwtClaims{

//  UserId: userID,
//  Role: role,
// RegisteredClaims: jwt.RegisteredClaims{
//   ExpiresAt: jwt.NewNumericDate(time.Now().Add(7*time.Minute)),

// },

// }

//  token:=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
//  return token.SignedString(JwtSecret)

// }

// func RefreshToken(userId uint,role string)(string,error){

//   claims:=JwtClaims{
//  UserId: userId,
//  Role: role,
//  RegisteredClaims: jwt.RegisteredClaims{
//   ExpiresAt: jwt.NewNumericDate(time.Now().Add(7*24*time.Hour)),
// },

// }

//  token := jwt.NewWithClaims(jwt.SigningMethodHS256 ,claims)
//  return token.SignedString(JwtSecret)

// }




type JwtClaims struct{

  UserId uint 
  Role string
  jwt.RegisteredClaims
}


func AccessToken(userId uint,role string)(string,error){

  claims := JwtClaims{
 
 UserId: userId,
 Role: role,
 RegisteredClaims: jwt.RegisteredClaims{
 ExpiresAt: jwt.NewNumericDate(time.Now().Add(7*time.Minute)),
},

}

 token:=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
 return token.SignedString(JwtSecret)

}








func RefreshToken(userId uint , role string)(string,error){
 
 claims:= JwtClaims{
 UserId: userId,
 Role: role,
 RegisteredClaims: jwt.RegisteredClaims{
 ExpiresAt: jwt.NewNumericDate(time.Now().Add(7*time.Minute)),
},
}

 token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
 return token.SignedString(JwtSecret)

}







