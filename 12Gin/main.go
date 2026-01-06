package main

import (
	// "encoding/json"

	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// func main() {

// 	r := gin.Default()

// r.GET("/",func(c *gin.Context) {

//    c.JSON(200,gin.H{

//      "message" : "GIN SERVER IS RUNNING",
// })
// })

// r.Run(":8080")
// }

// func main(){

//   r:= gin.Default()

//  r.GET("/",func(ctx *gin.Context) {

//   ctx.JSON(200,gin.H{"message" : "Home"})
// })

// r.GET("/api",func(ctx *gin.Context){

//   ctx.JSON(200,gin.H{"message":"API"})
// })
// api := r.Group("/api")
// {

//   api.GET("/hello",func(ctx *gin.Context) {

//    ctx.JSON(200,gin.H{"message":"API HELLO"})
// })

//    api.POST("/user",func(ctx *gin.Context){

//     ctx.JSON(200,gin.H{"message":"POST API"})
// })

// }

// r.Run(":8080")
// }

// type User struct{

//   Name string `json:"name"`
//   Age int     `json:"age"`

// }

// var users = []User{}

// func main(){

//    r := gin.Default()

// api := r.Group("/api")
// {

// api.GET("/users",getUsers)

// api.POST("/users",postUser)
// }

// r.Run(":8080")

// }

// func getUsers(c *gin.Context){
//   c.JSON(200,users)

// }

// func postUser(c *gin.Context){

// var user User

// if err:=c.ShouldBindJSON(&user); err!=nil{
//  c.JSON(400,gin.H{"error":"invalid JSon"})
//  return

// }

// users = append(users, user)

// c.JSON(201,user)
// }

// type LoginRequest struct{

//    Username string `json : "username"`
//    Password string  `json : "password"`
// }

// func login(c *gin.Context){

//   var req LoginRequest

//  if err:= c.ShouldBindBodyWithJSON(&req); err!=nil{

//     c.JSON(http.StatusBadRequest, gin.H{

//       "error" : "Invalid JSON",
// })
//    return
// }

// if req.Username == "admin" && req.Password == "1234"{

//    c.JSON(http.StatusOK,gin.H{

//     "message" : "Login Succesfull",
// })

// return
// }

// c.JSON(http.StatusUnauthorized,gin.H{

//     "errror" : "INVALID USERNAME AND PASSWORD",
// })

// }

// func main(){

//   r := gin.Default()

//  r.POST("/login",login)

// r.Run(":8080")
// }





type LoginRequest struct{

    Username string `json:"username"`
    Password string `json:"password"` 
   
}


var Sessions = map[string]string{}


func SessionGenerator()string{
 
  const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ12345678890"

  rand.Seed(time.Now().UnixNano())

  b := make([]byte,32)

  for i:= range b{
  
   b[i] = letters[rand.Intn(len(letters))] 

   
}

 return string(b)
}

func login(c *gin.Context){
  
    var req LoginRequest   

   if err := c.ShouldBindJSON(&req) ; err!=nil{
  
     c.JSON(http.StatusBadRequest,gin.H{
     
    "error" : "Invalid JSON",})

  return

}
if req.Username != "admin" || req.Password != "1234" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid Username or Password",
		})
		return
	}

sessionID := SessionGenerator()
Sessions[sessionID] = req.Username


c.SetCookie(
    
  "session_id",
  sessionID,
   3600,
   "/",
   "",
   false,
    true,

)

c.JSON(200,gin.H{

   "message" : "Log In Successfull",
})
  
}  


func profile(c *gin.Context){
  
   sessionId,err := c.Cookie("session_id")

  if err!= nil{
   
   c.JSON(http.StatusUnauthorized,gin.H{
   "error" : "Not Logged IN",
})
return
}

  username,ok := Sessions[sessionId]

 if !ok{
  
   c.JSON(http.StatusUnauthorized,gin.H{
    "error"  : "session Expired",

})

return
}

c.JSON(http.StatusOK,gin.H{
   "message" : "welcome",
   "username": username,
})
}

func logout(c *gin.Context){

   sessionId , err := c.Cookie("session_id")

   if  err == nil{

     delete(Sessions,sessionId)
}

 c.SetCookie("session_id","",-1,"/","",false,true)
c.JSON(http.StatusOK,gin.H{
  
    "message" : "Logged Out",
})
}

func loggingMiddleware() gin.HandlerFunc{
  return func(ctx *gin.Context) {
   
   start := time.Now()

 ctx.Next()

 
  latency := time.Since(start)
  
 status := ctx.Writer.Status()
 method := ctx.Request.Method
 path := ctx.Request.URL.Path

fmt.Printf(
  "[Auth] %s %s -> %d (%v)\n",
  method,path,status,latency,

)
}

}


func authMiddleware()gin.HandlerFunc{

    return func(ctx *gin.Context) {

       sessionID,err := ctx.Cookie("session_id")
      if err!=nil{
     
    ctx.JSON(http.StatusNotAcceptable,gin.H{
    "error" : "Invalid error",
})
    ctx.Abort()
    return 
}
  
   username,ok := Sessions[sessionID]
 
   if !ok {
  
   ctx.JSON(http.StatusNotAcceptable,gin.H{
    
      "error" : "Expired Session",
 })
  ctx.Abort()
  return 
}
   ctx.Set("username",username)
   ctx.Next()
}

}

func main(){
   r := gin.Default()

 Logmiddle := loggingMiddleware()
 authMiddle := authMiddleware()
 
  r.POST("/login",Logmiddle,login)
  r.POST("/logout",Logmiddle,logout)
  
  protected := r.Group("/")
  protected.Use(authMiddle)
  protected.GET("/profile",profile)
r.Run(":8080")

}