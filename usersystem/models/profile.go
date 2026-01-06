package models

type Profile struct{
  id uint `gorm:"primaryKey"`
  UserID uint `gorm:"not null"`
  Age  int `gorm:"not null"`
  Bio string 
}

