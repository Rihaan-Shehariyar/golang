package main

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
)



var session = map[string]string{}

func SessionGen() string {
	b := make([]byte, 32)

	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}

	return base64.URLEncoding.EncodeToString(b)
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {

	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Json"})
		return
	}

	if req.Email != "admin" || req.Password != "1234" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Credential"})
		return
	}

	sessionId := SessionGen()
	session[sessionId] = req.Email

	c.SetCookie("session_id", sessionId, 3600, "/", "/", false, true)
	c.JSON(200, gin.H{"message": "Login Successfull"})

}

func Profile(c *gin.Context) {

	sessionId, err := c.Cookie("session_id")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Please Log In"})
		return
	}

	username, ok := session[sessionId]
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Session Expired"})
		return
	}

	c.JSON(200, gin.H{"message": "Welecome",
		"username": username})

}

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sessionId, err := ctx.Cookie("sesssion_id")
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Please Log In"})
			ctx.Abort()
			return
		}
		username, ok := session[sessionId]
		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Session Expired"})
			ctx.Abort()
			return
		}

		ctx.Set("username", username)
		ctx.Next()
	}
}

func main() {

	r := gin.Default()

	r.POST("/login", Login)
	protected := r.Group("/profile")

	protected.Use(AuthMiddleware())

	protected.GET("/profile", Profile)
}


func SessionGene()string{
  b := make([]byte,32)

 _,err := rand.Read(b)
if err!=nil{
 panic("error")
}
 
  return base64.URLEncoding.EncodeToString(b)
}