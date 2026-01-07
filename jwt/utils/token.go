package utils

import (
	
	"time"

	"github.com/golang-jwt/jwt/v5"
)

 var jwtSecret = []byte("super-secret-key")

func Accestoken(userId uint) (string, error) {

	claims := jwt.MapClaims{
   "user_id" : userId,
   "exp" : time.Now().Add(15*time.Minute).Unix(), 

}

 token:=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
 return  token.SignedString(jwtSecret)

}

func RefreshToken(userId uint)(string ,error){

  claims:= jwt.MapClaims{
  "user_id" : userId,
   "exp" : time.Now().Add(7 * 24 * time.Hour).Unix(),
 } 

 token:=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
 return  token.SignedString(jwtSecret)

}


func ParseToken(tokenStr string)(jwt.MapClaims,error){
  token,err:=jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
   return jwtSecret,nil 

})

 if err!=nil || token.Valid{
  return nil,err
}

 return token.Claims.(jwt.MapClaims),nil

}


