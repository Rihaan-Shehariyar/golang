package handlers

import (
	"gorm/database"
	"gorm/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context){

 var reg models.Users

 if err:=c.ShouldBindJSON(&reg);err!=nil{

  c.JSON(http.StatusBadGateway,gin.H{
  "error":"invalid json",
})
return
}

 if reg.Name == "" || reg.Email==""{
  c.JSON(http.StatusBadGateway,gin.H{
  "error":"invalid email",
})

 return
}

tx:= database.DB.Begin()

 tx.Create(reg)

}