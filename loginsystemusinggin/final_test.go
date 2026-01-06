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
   
   r:=gin.Default()
  r.POST("/login",loginRequest)
  
  hashpassword,_:=bcrypt.GenerateFromPassword([]byte("123456"),bcrypt.DefaultCost)
  Users["admin"] = string(hashpassword)

  body := `{"username":"admin","password":"123456}`
  req,_:= http.NewRequest("POST","/login",strings.NewReader(body))
 req.Header.Set("Header-Type","application/json")
 
 w:=httptest.NewRecorder()

 r.ServeHTTP(w,req)

 if w.Code != http.StatusOK{
   t.Fatalf("expected code 200,got %d",w.Code)
}
 cookies:=w.Result().Cookies()

 if len(cookies)==0{
   t.Fatal("No cookies found")
}
 
  sessionCookie := cookies[0]

 if sessionCookie.Name != "session_id"{
  t.Fatalf("expexted session_id got %s",sessionCookie.Name)
}

 if sessionCookie.Value == ""{
 t.Fatal("value is empty")
}


}