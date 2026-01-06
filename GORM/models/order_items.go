package models

import "gorm.io/gorm"

type OrderItems struct {
	gorm.Model
   ProductId uint
   OrderId uint
   Quantity int
   Price float64
  
  Product Product
  Order Order

}