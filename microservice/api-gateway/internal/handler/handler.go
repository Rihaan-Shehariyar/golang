package handler

import (
	"context"
	orderpb "shared-proto/order"
	userpb "shared-proto/user"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var secret = []byte("mySecret")

func generateToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{

		"email": email,
		"exp":   time.Now().Add(time.Hour).Unix(),
	})

	return token.SignedString(secret)

}

func CreateUser(c *gin.Context, client userpb.UserServiceClient) {

	var req struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	res, err := client.CreateUser(ctx, &userpb.CreateUserRequest{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, res)

}

func Login(c *gin.Context, client userpb.UserServiceClient) {

	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	c.ShouldBindJSON(&req)

	_, err := client.Login(context.Background(), &userpb.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		c.JSON(401, gin.H{"error": "invalid credentials"})
		return
	}

	token, err := generateToken(req.Email)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"token": token})

}

func CreateOrder(c *gin.Context, client orderpb.OrderServiceClient) {
	var req struct {
		UserID  int32  `json:"user_id" binding:"required"`
		Product string `json:"product" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	res, err := client.CreateOrder(ctx, &orderpb.CreateOrderRequest{
		UserId:  req.UserID,
		Product: req.Product,
	})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, res)
}

func GetOrder(c *gin.Context, client orderpb.OrderServiceClient) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	res, err := client.GetOrder(ctx, &orderpb.GetOrderRequest{
		Id: int32(id),
	})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, res)
}
