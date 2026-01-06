package main

import "fmt"

func main() {

	var ptr *int
    x := 10
    ptr = &x

	fmt.Println(*ptr)

}