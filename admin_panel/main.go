package main

import (
	"admin/database"
	"admin/handlers"
	"admin/middleware"
	"admin/models"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	database.Connect()
	database.DB.AutoMigrate(&models.User{})

	r := gin.Default()

	r.POST("/login", handlers.Login)

	admin := r.Group("/admin")
	admin.Use(middleware.JwtAuthMiddleware(), middleware.AdminOnly())
	{

		admin.GET("/users", handlers.GetAllUsers)
		admin.POST("/users", handlers.CreateUserbyAdmin)
		admin.PUT("/users/:id", handlers.UpdateUser)
		admin.DELETE("/users/:id", handlers.DeleteUser)

	}

	log.Println("Server running")
	r.Run(":9090")

}
