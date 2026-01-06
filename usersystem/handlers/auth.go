package handlers

import (
	"user_system/database"
	"user_system/models"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context){

  user:= models.User{
  Name: "john",
  Email: "rihaanshehariyar@gmail.com",
  Password: "1234",
  
}
  
 db:= database.DB

 db.Create(&user)


}