package main

import "fmt"

func main() {

	// var num int

	// fmt.Println("Enter The number to check: ")
	// fmt.Scan(&num)

	// if num%2 == 0 {
	// 	fmt.Println("Even Number")
	// } else {
	//      fmt.Println("Odd Number")
	// }

	// var num1, num2 int
	// var op string

	// fmt.Println("Enter the Number")
	// fmt.Scan(&num1)
	// fmt.Println("Enter the Operator")
	// fmt.Scan(&op)
	// fmt.Println("Enter the number")
	// fmt.Scan(&num2)

	// switch op {
	// case "+":
	// 	fmt.Println("Result : ", num1+num2)
	// case "-":
	// 	fmt.Println("Result : ", num1-num2)
	// case "*":
	// 	fmt.Println("Result : ", num1*num2)
	// case "/":
	// 	if num2 == 0 {
	// 		fmt.Println("Cannot divide by zero")
	// 	} else {
	// 		fmt.Println("Result : ", num1/num2)
	// 	}

	// default:
	// 	fmt.Println("Invalid Operator")
	// }


  var num int
 
 fmt.Println("Enter The Number : ")
 fmt.Scan(&num)

 a,b := 0,1

 for i := 0; i < num; i++ {
	fmt.Print(a,"")
    next := a+b
    a= b
    b = next
    
 }

}
