package handlers

import (
	"jwt/database"
	"jwt/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProfile(c *gin.Context){
 
 userId := c.MustGet("user_id").(uint)
 
 var user models.User
 database.DB.Preload("Profile").First(&user,userId)

 c.JSON(http.StatusOK,gin.H{
 "id" : user.ID,
 "name" : user.Name,
 "email": user.Email,
 "profile" : user.Profile,
})

}