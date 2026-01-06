package database


import (
	"log"
    "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(){
    
  dsn:="host=localhost user=postgres password=Rihaan@123 dbname=ecommerce port=5432 sslmode=disable"
 
  db,err :=gorm.Open(postgres.Open(dsn),&gorm.Config{})

  if err!=nil{
     log.Fatal("error while connecting",err)
}
  
  DB=db

  log.Println("DATABASE CONNECTED SUCCESSFULLY")
 
}