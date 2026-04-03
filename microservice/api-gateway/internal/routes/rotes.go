package routes

import (
	"api-gateway/internal/client"
	"api-gateway/internal/handler"
	"api-gateway/internal/middleware"

	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine, clients *client.Clients) {

	r.Use(middleware.RateLimit())
	r.POST("/user", func(ctx *gin.Context) {
		handler.CreateUser(ctx, clients.UserClient)
	})

	r.POST("/login", func(ctx *gin.Context) {
		handler.Login(ctx, clients.UserClient)
	})

	auth := r.Group("/")
	auth.Use(middleware.Auth())

	auth.POST("/order", func(ctx *gin.Context) {
		handler.CreateOrder(ctx, clients.OrderClient)
	})

	auth.GET("/order/:id", func(ctx *gin.Context) {
		handler.GetOrder(ctx, clients.OrderClient)
	})

}
