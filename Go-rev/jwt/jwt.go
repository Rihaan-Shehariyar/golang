package jwt

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserId uint
	email  string
	role   string
	jwt.RegisteredClaims
}

var jwtsecret = []byte("secret-key")

func AccessToken(userId uint, email string, role string) (string, error) {

	claims := Claims{
		UserId: userId,
		email:  email,
		role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtsecret)

}

func JwtAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.GetHeader("Authorization")
		if auth == "" {

		}

		tokenstr := strings.TrimPrefix(auth,"Bearer ")
		claims := Claims{}

		token, err := jwt.ParseWithClaims(tokenstr, claims, func(t *jwt.Token) (interface{}, error) {
			return jwtsecret, nil
		})

		if err != nil || !token.Valid {

		}

	}

}
