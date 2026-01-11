package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
  Name string `json:"name" gorm:"not null"`
  Email string `json:"email" gorm:"unique"`
  Password string `json:"password" gorm:"omitempty"`
  Role string `json:"role" gorm:"default:user"`
}

