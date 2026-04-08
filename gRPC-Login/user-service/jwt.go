package main

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtSecret = []byte("secret")

func GenerateToken(username string) (string, error) {

	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)


}

func ValidateToken(tokenstr string)(*jwt.Token,error){
 return jwt.Parse(tokenstr,func(t *jwt.Token) (interface{}, error) {
  return jwtSecret,nil
})
}
