package models

type Users struct {
	ID    uint
	Name  string  `gorm:"not null"`
	Email string `gorm:"unique"`
	Age   int    `gorm:"not null" binding:"required"`
}
