package models

type User struct{
  Id uint
  Name string `json:"name"`
  Email string `json:"email" gorm:"unique"`
  

}

type Post struct{
 
 UserId uint `g`

}