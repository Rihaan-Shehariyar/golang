package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func TestCookieValidation(t *testing.T) {
	gin.SetMode(gin.TestMode)

  r := gin.Default()
  r.POST("/login",loginRequest)
  
  hashpassword,_ := bcrypt.GenerateFromPassword([]byte("123456"),bcrypt.DefaultCost)
  users["admin"] = string(hashpassword)

  body := `{"username":"admin","password":"123456"}`
  req,_ := http.NewRequest("POST","/login",strings.NewReader(body))
 
 w := httptest.NewRecorder()

 r.ServeHTTP(w,req)

 if w.Code != http.StatusOK{
  t.Fatalf("expected code got %d",w.Code)
 
}
cookies := w.Result().Cookies()

 if len(cookies) == 0{
   t.Fatal("No Cookie Availbale")
}

 if cookies[0].Name != "session_id"{
   t.Fatal("no session created")
}   


}

func TestWithoutCookie(t *testing.T) {
   gin.SetMode(gin.TestMode)

 r:= gin.Default()
 r.GET("/home",authMiddleWare(),home)

  req,_ := http.NewRequest("GET","/login",nil)
  w := httptest.NewRecorder()
 
 r.ServeHTTP(w,req)
 
  if w.Code ==http.StatusOK{
   t.Fatalf("expected 400,got %d ",w.Code)
}
 
}