package main

import (
	"fmt"
	"gorm/database"
	"gorm/models"
)


func CreateUser(){
 
  user := models.User{
    Name: "Rihaan",
    Email: "rihaan@gmail.com",
    Password: "Rihaan@123",
}
   
  database.DB.Create(&user)

}

func ReadUsers(){
    var users []models.User
  
   result:= database.DB.Find(&users)
  
   if result.Error!=nil{
    fmt.Println("ERROR",result.Error)
   return
}
 
  fmt.Println(users)
}


func main() {
	database.Connect()

  database.DB.AutoMigrate(
   
   &models.User{},
   &models.Product{},
   &models.Order{},
   &models.OrderItems{},
   
   
)
  CreateUser()
  ReadUsers()


} 