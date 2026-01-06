package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Register struct {
	Username string`json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

var Users = map[string]string{}
var Sessions = map[string]string{}

func registration(c *gin.Context){

    var reg Register
   
  if err:=c.ShouldBindJSON(&reg);err!=nil{
   
   c.JSON(400,gin.H{
    "error" : "Error in JSON",
})
 return
}
   
    hashpassword,err:=bcrypt.GenerateFromPassword([]byte(reg.Password),bcrypt.DefaultCost)

  if err!=nil{
       
    c.JSON(400,gin.H{
   "error" : "error while hasshing",
})
}
   Users[reg.Username] = string(hashpassword)

 c.JSON(200,gin.H{
  "message" : "Registration Completed",
})
   
}

func loginRequest(c *gin.Context){
   var log Login
  
  if err:=c.ShouldBindJSON(&log);err!=nil{
   
   c.JSON(400,gin.H{
    "error" : "Error in JSON",
})
 return
}
  
 hashpass,exist := Users[log.Username]

 if !exist{
  c.JSON(401,gin.H{
   "erroe" : "No username found",
})
}
  err:=bcrypt.CompareHashAndPassword([]byte(hashpass),[]byte(log.Password))

  if err!=nil{
     c.JSON(401,gin.H{
   "error" : "Incorrect Password",
})
 }
  
 sessionId := uuid.NewString()
 Sessions[sessionId] = log.Username

  c.SetCookie("session_id",sessionId,3600,"/","",false,true)

 c.JSON(200,gin.H{
   "message" : "Successfull",
})
    
}

func authMiddleWare()gin.HandlerFunc{
 return func(ctx *gin.Context) {
      
   sessionId,err := ctx.Cookie("session_id")

  if err!=nil{
  ctx.JSON(401,gin.H{
   "error" : "No session found",
})
 ctx.Abort()
 return 
}
   
    username,exist := Sessions[sessionId]
    if !exist{
    ctx.JSON(401,gin.H{
   "error" : "session Expired",
})
 ctx.Abort()
 return 
  
}
   ctx.Set("username",username)
  ctx.Next()
}
}
  
 func loggMiddleware()gin.HandlerFunc{
 return func(ctx *gin.Context) {
   
    start := time.Now()
  
   ctx.Next()
  
  latency := time.Since(start)
  status := ctx.Writer.Status()

  log.Printf(
  "[Auth] %s %s | %d |%v",ctx.Request.Method,ctx.Request.URL.Path,status,latency,
)
 
}
}



func home(c *gin.Context){
   username,_:= c.Get("username")
   
  c.JSON(200,gin.H{
  "message" : "welcome"+ username.(string),
})
}

func logOut(c *gin.Context){

    sessionId,err:=c.Cookie("username")
if err ==nil{
  delete(Sessions,sessionId)
}
 c.SetCookie("session_id","",-1,"/","",false,true)

}

func main(){
  
 r:=gin.Default()

 api:=r.Group("/api")
{
 api.POST("/reg",registration)
 api.POST("/login",loggMiddleware(),loginRequest)
 api.POST("/logout",logOut)
api.GET("/home",authMiddleWare(),home)
}

r.Run(":8080")
 
}