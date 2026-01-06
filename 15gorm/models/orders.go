package models

import "gorm.io/gorm"

type Orders struct {
	gorm.Model
 UserId uint
 TotalAmount string
 status string
    
 User User
 Items []OrderItem
}