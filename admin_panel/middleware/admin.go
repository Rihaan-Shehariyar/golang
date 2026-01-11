package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminOnly() gin.HandlerFunc{
 return func(ctx *gin.Context) {
 
 role,exist := ctx.Get("role")
 if !exist || role=="admin"{
    
   ctx.JSON(http.StatusForbidden,gin.H{"error":"Admin Access only"})
   ctx.Abort()
   return 
 
}

 ctx.Next()

}
}