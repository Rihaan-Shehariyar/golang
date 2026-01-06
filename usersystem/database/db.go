package database

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(){
   

  dsn:="host=localhost name=postgres password=Rihaan@123 dbname=usersysytem  port=5432 disable=disabled"
 
 db,err:= gorm.Open(postgres.Open(dsn),&gorm.Config{})
if err!=nil{
  log.Fatal("couldnt connect gorm")
}

  sqlDB ,err:=db.DB()
if err!=nil{
  log.Fatal("Failed to connect sqlDb",sqlDB)
}
 
 sqlDB.SetMaxOpenConns(25)
 sqlDB.SetMaxIdleConns(10)
 sqlDB.SetConnMaxLifetime(time.Hour)

 DB = db
}