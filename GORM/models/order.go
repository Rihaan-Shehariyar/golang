package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
   UserId uint
   TotalAmount float64
   status string

User User
Items []OrderItems

}  