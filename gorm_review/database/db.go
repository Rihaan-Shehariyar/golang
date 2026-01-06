package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(){

  dsn:= "host=localhost user=postgres password=Rihaan@123 dbname=user port=5432 sslmode=disable"

 db,err:= gorm.Open(postgres.Open(dsn),&gorm.Config{})

 if err!=nil{
   log.Println("COULDNT CONNECT DB")
 }

 DB=db

}