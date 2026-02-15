package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

 dsn:="host=localhost user=postgres dbname=admin port=5432 sslmode=disable"


 db,err:=gorm.Open(postgres.Open(dsn),&gorm.Config{})
 if err!=nil{
      panic("CanNot Connect to DB")
}
 
 DB = db
 

}