package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// type Student struct {
// 	Username string `json:"username" binding:"required"`
// 	Email    string `json:"email" binding:"required,email"`
// 	Password string `json:"password" binding:"required,min=6"`
// }

// func register(c *gin.Context){

//   var std1 Student

//   if err:= c.ShouldBindJSON(&std1); err!=nil{

//     c.JSON(400,gin.H{
//     "error" : "Invalid Json",
// })
// return
// }

// c.JSON(200,gin.H{

//  "message" : "Login Successfull",
//   "data" : std1,
// })

// }

// func main(){

//   r:=gin.Default()

//   r.POST("/register",register)
//  r.Run(":8080")
// }



type Register struct{
  
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required,min=6"` 
}

type Login struct{

    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required,min=6"` 
}


var users = map[string]string{}

func Registration(c *gin.Context){

    var reg Register

  if err:= c.ShouldBindJSON(&reg) ; err !=nil{
  
     c.JSON(400,gin.H{

    "error" : err.Error(),
})

return
}

hashpassword,err:= bcrypt.GenerateFromPassword(

   []byte(reg.Password),
   bcrypt.DefaultCost,
)

if err !=nil{

  c.JSON(400,gin.H{

   "error" : "error while hashing",
})

return
}

users[reg.Username] = string(hashpassword)

 c.JSON(200,gin.H{
    "message" : "Registration Successfull",
   
})

}


func loginhandle(c *gin.Context){

   var req Login

 if err := c.ShouldBindJSON(&req) ; err!=nil{
   
     c.JSON(400,gin.H{

    "error" : err.Error(),
})

  return
}

  hashpassword,exist := users[req.Username]

  if !exist{

    c.JSON(http.StatusUnauthorized,gin.H{
   "error" : "Invalid UserName or Password",
})
  return
}

 err := bcrypt.CompareHashAndPassword(
   []byte(hashpassword),
   []byte(req.Password),
)

 if err!=nil{

   c.JSON(http.StatusUnauthorized,gin.H{
   "error" : "Invalid UserName or Password",
})

return
}
  c.JSON(http.StatusOK,gin.H{

   "message" : "Login Succesfull",
})
}

func main(){

  r := gin.Default()

  r.POST("/register",Registration)
  r.POST("/login",loginhandle)
 r.Run(":8080")
}