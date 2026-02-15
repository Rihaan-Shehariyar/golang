package handlers

import (
	"admin/database"
	"admin/models"

	"github.com/gin-gonic/gin"
)

func GetAllUser(c *gin.Context){
 
 var user []models.User

 if err:=database.DB.Find(&user).Error;err!=nil{
  c.JSON(400,gin.H{"error":"Failed to Fetch users"})
 return
}

 c.JSON(200,user)

}



