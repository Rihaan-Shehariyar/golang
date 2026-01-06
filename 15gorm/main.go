package main

import (
	database "gorm/DATABASE"
	"gorm/models"
	"log"
)

// import (
// 	database "gorm/database"
// 	models "gorm/models"
// 	"log"
// )

// func main() {
// 	database.Connect()

//  err := database.DB.AutoMigrate(
//   &models.User{},
//   &models.Product{},
//   &models.Orders{},
//   &models.OrderItem{},
// )

//  if err!=nil{
//   log.Fatal("Migration Failed:",err)
// }

//  log.Println("Application started successfully")

// }


func main(){
    
  database.Connect()

 err:= database.DB.AutoMigrate(
  &models.User{},
   &models.Product{},
 &models.Orders{},
 &models.OrderItem{},
)
 
  if err!=nil{
  log.Fatal("MIgration failed",err)
} 
 
  // log.Pri)
}