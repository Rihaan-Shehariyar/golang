package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestPling(t *testing.T){

   gin.SetMode(gin.TestMode)

  w := httptest.NewRecorder()

  r := gin.Default()
  
  r.GET("/pling",pling)


  req,err := http.NewRequest("GET","/pling",nil)


  if err!=nil{
   
t.Fatal(err)
  
}
    r.ServeHTTP(w,req)

  if w.Code !=http.StatusOK{

    t.Errorf("expected status 200,but %d",w.Code)
  
}
  
}