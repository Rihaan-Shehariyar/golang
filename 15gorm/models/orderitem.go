package models

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	OrdersId  uint
	ProductId uint
	Quantity  int
	Price     float64

	Orders  Orders
	Product Product
}