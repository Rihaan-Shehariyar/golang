package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(){

 dsn:="host=localhost user=postgres dbname=admin password=Rihaan@123 port=5432 sslmode=disable"

 db,err := gorm.Open(postgres.Open(dsn),&gorm.Config{})
if err!=nil{
   panic("Cannot connect DB")
}

 DB =db

}

