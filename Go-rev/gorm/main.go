package main

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"not null"`
	Email      string `gorm:"uniqueIndex"`
    Orders []Order
	Created_at time.Time
}

type Order struct{
  ID uint `gorm:"primaryKey"`
  UserId uint `gorm:"foriegnKey"`
  Total int
  Created_at time.Time 
}

var DB *gorm.DB

func Connect(){
 dsn := "host=localhost user=postgres password=Rihaan@123 dbname=ecommerce port=5432 sslmode=disable"

 db,err := gorm.Open(postgres.Open(dsn),&gorm.Config{})
 
if err!=nil {
	log.Fatalf("error while connecting")
}
 
DB=db


}

func main(){
Connect()

 DB.AutoMigrate(&User{})

 DB.Create(&User{
 Name: "name",
})

}