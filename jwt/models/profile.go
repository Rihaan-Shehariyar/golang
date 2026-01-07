package models

import "gorm.io/gorm"

type Profile struct {
	gorm.Model
  UserId uint
  Age int 
  Bio string
}
