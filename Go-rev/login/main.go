package main

import (
	"crypto/rand"
	"encoding/base64"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GenerateSessionId() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.RawURLEncoding.EncodeToString(b)
}

type Register struct {
	Email    string `json:"email"`
	Password string `json:"email"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var Users = map[string]string{}
var session = map[string]string{}

func RegisterRequest(c *gin.Context) {

	var req Register

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(401, gin.H{"error": "Invalid Json"})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	Users[req.Email] = string(hashedPassword)

}

func LoginRequest(c *gin.Context) {

	var req Login

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(401, gin.H{"error": "Invalid Json"})
		return
	}

	hashpassword, exist := Users[req.Email]
	if !exist {
		c.JSON(400, gin.H{"error": "Invalid emai;"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(hashpassword), []byte(req.Password))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid password"})
		return
	}

	sessionId := GenerateSessionId()
	session[sessionId] = req.Email

	c.SetCookie("session_id", sessionId, 3600, "/", "", false, true)

}
