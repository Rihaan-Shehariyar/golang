package handlers

import (
	"admin/database"
	"admin/models"
	"admin/utility"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *gin.Context){
 
 var req AuthRequest

if err:=c.ShouldBindJSON(&req);err!=nil{
  
  c.JSON(http.StatusUnauthorized,gin.H{
   "error" : "Invalid JSON",
})
 return
}

 var user models.User

 if err:=database.DB.Where("email=?",user.Email).First(&user).Error;err!=nil{

   c.JSON(401,gin.H{"error":"Invalid Credentials"})
 return

}

 if err:=bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(req.Password))
  err!=nil{
   c.JSON(401,gin.H{"error":"Invalid Password"})
  return
}

 token,_ := utility.GenerateToken(user.ID,user.Email,user.Password)
 
 c.JSON(200,gin.H{"token":token})

}



func Signup(c *gin.Context){

  var user models.User

if err:=c.ShouldBindJSON(&user);err!=nil{
  c.JSON(400,gin.H{"error":"Invalid Json"})
   return
}

 if user.Email==""||user.Password==""{
   c.JSON(400,gin.H{"error":"Email and Password required"})
 return
}

 hash,err:=bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost)
  if err!=nil{
  
  c.JSON(500,gin.H{"error":"Error While Hashing"})
 return
}

  user.Password = string(hash)
  user.Role = "user"

 if err := database.DB.Create(&user).Error;err!=nil{
   c.JSON(400,gin.H{"error":"email already exist"})
  return
}


 c.JSON(200,gin.H{"message":"user Created"})
 

}
