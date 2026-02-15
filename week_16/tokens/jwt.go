package tokens

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)


  var JwtSecretKey = []byte(os.Getenv("Secret Key"))

type JwtClaims struct {
	UserId uint
	email  string
	jwt.RegisteredClaims
}

func GenerateAcessTokens(userId uint, email string) (string, error) {

	claims:=JwtClaims{
   
  UserId: userId,
  email: email,
  RegisteredClaims: jwt.RegisteredClaims{
  ExpiresAt: jwt.NewNumericDate(time.Now().Add(7*time.Minute)),
},

  

}

 tokens := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
 return tokens.SignedString(JwtSecretKey)


}


func GenerateRefereshToken(userId uint,email string)(string,error){

  claims:= JwtClaims{
 UserId: userId,
 email: email,
 RegisteredClaims: jwt.RegisteredClaims{
   ExpiresAt: jwt.NewNumericDate(time.Now().Add(7*24*time.Hour)),
},
}

 token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
 return token.SignedString(JwtSecretKey)

}