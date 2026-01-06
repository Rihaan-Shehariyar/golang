package main

import "fmt"

func main() {
	
	U1 := User{"John", 21, true}
	fmt.Println(U1)
	fmt.Printf("Name is %v and he is a %v years \n",U1.Name,U1.Age)
	U1.ChangeAge()

   var ptr *int
   fmt.Println(ptr)
}
type User struct {
		Name    string
		Age     int
		Success bool
	}

func (p User) ChangeAge(){
	p.Age = 28
	fmt.Printf("Age : %v",p.Age)
}


