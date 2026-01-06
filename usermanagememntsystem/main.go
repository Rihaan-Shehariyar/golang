package main

import (
	"user_system/database"
	"user_system/models"
	"user_system/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	database.Connect()

  database.DB.AutoMigrate(
   &models.User{},
  &models.Profile{},
)


  r := gin.Default()
 
  routes.SetUpRoutes(r)

 r.Run(":9090")
}
