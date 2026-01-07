package handlers

import (
	
	"jwt/database"
	"jwt/models"
	"jwt/utils"

	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Signup(c *gin.Context){

 var req AuthRequest

 if err:= c.ShouldBindJSON(&req) ; err!=nil{
 
  c.JSON(http.StatusUnauthorized,gin.H{
   "error" : "Invalid Json",
})
  return
}

 hashpassword,err := bcrypt.GenerateFromPassword([]byte(req.Password),bcrypt.DefaultCost)
 if err!=nil{
 c.JSON(http.StatusInternalServerError,gin.H{
   "error" : "Error in Hashing",
})
 return
}

 user := models.User{
  Name: req.Name,
  Email: req.Email,
  Password: string(hashpassword),
}

 if err:= database.DB.Create(&user).Error; err!=nil{
  c.JSON(http.StatusInternalServerError,gin.H{
   "error" : "Unable to Create Db",
})
 return
}
 
 c.JSON(http.StatusCreated,gin.H{"message":"User CReated"})

}


func Login(c *gin.Context){
 var req AuthRequest
 var user models.User

 c.ShouldBindJSON(&req)
  err := database.DB.Where("email=?",req.Email).First(&user).Error

 if err!=nil{
 c.JSON(http.StatusUnauthorized,gin.H{
 "error" : "INvalid email",
}) 
 return
}

 res:=bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(req.Password))
 if res !=nil{
 c.JSON(http.StatusUnauthorized,gin.H{
 "error" : "Invalid Password",
})
 return
}
 
 access,_ := utils.Accestoken(user.ID)
 refresh,_ := utils.RefreshToken(user.ID)

 user.RefreshToken = refresh
 database.DB.Save(&user)

 c.JSON(200,gin.H{"access_token" : access,"refresh_token" : refresh})


}

func Refresh(c *gin.Context){
  
 var body struct{ 

 RefreshToken string `json:"refresh_token"`
}

	c.ShouldBindJSON(&body)

 claims,err:=utils.ParseToken(body.RefreshToken)
 	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid refresh token"})
		return
	}

 var user models.User
 database.DB.First(&user,uint(claims["user_id"].(float64)))

if user.RefreshToken !=body.RefreshToken{
 c.JSON(http.StatusUnauthorized,gin.H{
 "error" : "token Mismatch",
})
 return
}

 access,_ := utils.Accestoken(user.ID)
 c.JSON(http.StatusOK, gin.H{"access_token": access})
}