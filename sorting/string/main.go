package main

import "fmt"

func ReverseString(s string) string {
	b := []byte(s)

	i := 0
	j := len(b) - 1

	for i < j {

		b[i], b[j] = b[j], b[i]
		i++
		j--

	}

	return string(b)

}

func main() {
	b := "Hello"
	
	fmt.Println(ReverseString(b))

}
