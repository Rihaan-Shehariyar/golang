package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Email    string `gorm:"not null,unique"`
	Password string `gorm:"not null"`
    deletet_at gorm.DeletedAt
    Profile Profile `gorm:"constraint:OnDelete:CASCADE"`
}