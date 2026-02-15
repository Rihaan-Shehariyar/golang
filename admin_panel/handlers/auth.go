package handlers

import (
	"admin/database"
	"admin/models"
	"admin/utility"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type input struct{
  
  Email string `json:"email"`
  Password string `json:"password"`

}

func Login(c *gin.Context){
  
 var log input

if err:=c.ShouldBindJSON(&log);err!=nil{
  c.JSON(400,gin.H{"error":"Invalid Json"})
 return
}

 var user models.User

 if err:=database.DB.Where("email=?",log.email).First(&user).Error;err!=nil{
  c.JSON(400,gin.H{"error":"Invalid email"})
 return
}

 err:=bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(log.password))
 if err!=nil{
  c.JSON(401,gin.H{"error":"Invalif Password"})
 return
}


 accessToken,_:=utility.AcesssToken(user.ID,user.Role)
 refreshToken,_:=utility.RefreshToken(user.ID,user.Role)

 user.RefreshToken = refreshToken
 database.DB.Save(&user)

c.SetCookie("refresh_token",refreshToken,7*24*3600,"/","",false,true)


c.JSON(200,gin.H{"acces_tokrn":accessToken})
}



func Refresh(c *gin.Context){
 
 refrehToken,err:=c.Cookie("refresh_token")
 if err!=nil{
  c.JSON(401,gin.H{"error":"Refresh Token missinng"})
 return
}

 token,err:=jwt.Parse(refrehToken,func(t *jwt.Token) (interface{}, error) {
  return []byte("refresh_token"),nil
 })

if err!=nil{
  c.JSON(401,gin.H{"error":"Invalid refresh token"})
 return
}

 var user models.User

 if err:= database.DB.Where("refresh_token= ?",user.RefreshToken).First(&user).Error;err!=nil{
  c.JSON(401,gin.H{"error":"Token revoked"})
 return

}

 

}
