package handlers

import (
	"admin/database"
	"admin/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GetAllUsers(c *gin.Context){
 
 var users []models.User
   database.DB.Find(&users)
  c.JSON(200,users)

}


func CreateUserbyAdmin(c *gin.Context){

  var user models.User

if err:= c.ShouldBindJSON(&user);err!=nil{

  c.JSON(400,gin.H{"error" : "Invalid Input"})
  return

}

 hash,_:=bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost)
 user.Password = string(hash)

 database.DB.Create(&user)
 c.JSON(201,user)

}


func UpdateUser(c *gin.Context){

  id:=c.Param("id")

 var user models.User

 if err:=database.DB.First(&user,id).Error;err!=nil{
  
 c.JSON(404,gin.H{"error":"User Not Found"})

}

 c.ShouldBindJSON(&user)
 database.DB.Save(&user)
 c.JSON(200,user)

}


func DeleteUser(c *gin.Context){

 id:=c.Param("id")
 database.DB.Delete(&models.User{},id)
 c.JSON(200,gin.H{"message":"Deleted Successfully"})
}