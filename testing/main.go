package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

)

// import (
// 	"log"
// 	"net/http"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"github.com/google/uuid"
// )

// type User struct {
// 	Name  string `json : "username"`
// 	Email string `json : "email"`
// }

// var sessions = map[string]string{}

// func login(c *gin.Context){

//   var req User

//    if err:=c.ShouldBindJSON(&req) ; err!=nil{

//      c.JSON(http.StatusBadRequest,gin.H{

//       "error" : "Invalid JSON",
// })

// return
// }

//   if req.Name != "admin" || req.Email != "1234"{

//     c.JSON(http.StatusBadRequest,gin.H{
//      "error" : "INvalid username and password",
// })
//   return
// }

//   sessionID := uuid.NewString()
//   sessions[sessionID] = req.Name

//   c.SetCookie(
//     "session_id",
//      sessionID,
//     3600,
//      "/",
//      "",
//      false,
//     true,
// )

//    c.JSON(http.StatusOK,gin.H{

//     "message" : "LOGIN SUCCESS",
// })
// }

// func authMiddleWare() gin.HandlerFunc{
// return func(ctx *gin.Context) {

//      sessionID,err := ctx.Cookie("session_id")

//     if err!=nil{
//    ctx.JSON(http.StatusUnauthorized,gin.H{

//      "error" : "Not logged in",
// })
//   ctx.Abort()
//   return
// }
//     username,exists := sessions[sessionID]

//    if !exists{

//      ctx.JSON(http.StatusUnauthorized,gin.H{
//     "error" : "SESSION EXPIRED",
// })
//   ctx.Abort()
//   return
// }
//     ctx.Set("username",username)
//     ctx.Next()
// }

// }

// func dashboard(c *gin.Context){

//    username,_ :=c.Get("username")

//   c.JSON(http.StatusOK,gin.H{

//    "message" : "Welcome to dashboard",
//     "user" : username,
// })

// }

// func logout(c *gin.Context){

//    sessionId,err := c.Cookie("session_id")

//   if err == nil{

//    delete(sessions,sessionId)
// }

//    c.SetCookie(
//    "session_id",
//     "",
//      -1,
//      "/",
//     "",
//     false,
//     true,
// )
//   c.JSON(http.StatusOK,gin.H{
//     "message" : "Logged Out",
// })
// }

// func authLogger()gin.HandlerFunc{

//    return func(ctx *gin.Context) {

//     start := time.Now()

//   ctx.Next()

//    latency := time.Since(start)
//   status := ctx.Writer.Status()

//   log.Printf(

//    "[Auth] %s %s | %d | %v",
//   ctx.Request.Method,
//   ctx.Request.URL.Path,
//   status,
// latency,
// )

// }

// }

// func main(){

//  r:= gin.Default()

//  api := r.Group("/api")
// {
//  api.POST("/login",authLogger(),login)
//  api.POST("/logout",authLogger(),logout)

// api.GET("/users",authMiddleWare(),dashboard)

// }

// r.Run(":8080")

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
var sessions = map[string]string{}

func register(c *gin.Context){

    var reg Register

  if err := c.ShouldBindJSON(&reg); err!=nil{

c.JSON(http.StatusBadRequest,gin.H{
      
     "error" : "Invalid JSON",
})

  return
}

  hashpassword,err := bcrypt.GenerateFromPassword(
     []byte(reg.Password),
     bcrypt.DefaultCost,
) 
   
  if err != nil{

    c.JSON(http.StatusBadRequest,gin.H{

     "error" : "Error while Hashing",
})
return
}

  users[reg.Username] = string(hashpassword)


   c.JSON(200,gin.H{

   "message" : "REgistratio succesfull",

})
   
}


func loginRequest(c *gin.Context){
   
    var log Login

   if err:= c.ShouldBindJSON(&log)
   
      err!=nil{

     c.JSON(400,gin.H{
    
    "error" : "Invalid JSON",
})
 return
}
    
  hashpassword,exist:=users[log.Username]

 if !exist{
    
   c.JSON(http.StatusUnauthorized,gin.H{
    "Error" : "USername not exist,Please register",
})

return
} 
  err := bcrypt.CompareHashAndPassword(
    []byte(hashpassword),
    []byte(log.Password),
)
      
  if err!=nil{
    c.JSON(http.StatusUnauthorized,gin.H{

    "error"  : "Invalid password",
}) 
  return
}


  sessionID := uuid.NewString()
  sessions[sessionID] = log.Username

  c.SetCookie(
    "session_id",
     sessionID,
     3600,
     "/",
     "",
    false,
     true,
 )

  c.JSON(http.StatusOK,gin.H{

    "message" : "Logged In SUCCESFULL",
})
   
}

func authMiddleWare()gin.HandlerFunc{
 return func(c *gin.Context) {
    
     sessionId,err := c.Cookie("session_id")
   if err!=nil{

   c.JSON(http.StatusUnauthorized,gin.H{

     "message" : "LOGIN TO continue",
})
c.Abort()
return 
}

  username,exist := sessions[sessionId]
 
  if !exist{
   
     c.JSON(http.StatusUnauthorized,gin.H{

     "message" : "sesssion EXPIRED",
})
c.Abort()
return 
}

 c.Set("username",username)
 
 c.Next()
}
 
}

func home(c *gin.Context){
    
    username,_:=c.Get("username")

     c.JSON(200,gin.H{
    "username" : username,

  })
}
    

func main(){

  r := gin.Default()
  r.POST("/login", loginRequest)

}



