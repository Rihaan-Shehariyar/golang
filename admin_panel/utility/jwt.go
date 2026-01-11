package utility

import (
	
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JwtSeecret = []byte(os.Getenv("SECRET-KEY"))

type JwtClaims struct{
   
   UserId uint `json:"user_id"`
   Email string `json:"email"`
   Role string `json:"role"`
   jwt.RegisteredClaims
}

func GenerateToken(userId uint, email string, role string) (string, error) {

	claims := JwtClaims{
   
  UserId: userId,
  Email: email,
  Role: role,
 RegisteredClaims: jwt.RegisteredClaims{
   ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
},
 
}
 
 token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims) 
 return token.SignedString(JwtSeecret)

}