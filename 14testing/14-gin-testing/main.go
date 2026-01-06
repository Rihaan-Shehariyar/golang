package main

import "github.com/gin-gonic/gin"

func pling(c *gin.Context){

   c.JSON(200,gin.H{

   "message" : "pong",
})
}