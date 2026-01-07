package main

import (
	"jwt/database"
	"jwt/handlers"
	"jwt/middlewares"
	"jwt/models"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	database.DB.AutoMigrate(&models.User{}, &models.Profile{})

	r := gin.Default()

	r.POST("/signup", handlers.Signup)
	r.POST("/login", handlers.Login)
	r.POST("/refresh", handlers.Refresh)

	protected := r.Group("/")
	protected.Use(middlewares.AuthMiddleware())
	protected.GET("/profile", handlers.GetProfile)

	r.Run(":8080")
}
