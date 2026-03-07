package main

import (
	"fmt"
	"strings"
)

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

func CheckPalindrome(s string) bool {

	i := 0
	j := len(s) - 1

	for i < j {

		if s[i] != s[j] {
			return false
		}

		i++
		j--
	}

	return true

}

func VowelCount(s string) int {

	count := 0

	for _, ch := range s {

		if strings.Contains("aeiou", string(ch)) {
			count++
		}
	}

	return count

}

func main() {

	b := "Hello"
	c := "madam"
 
	fmt.Println(ReverseString(b))
	fmt.Println(CheckPalindrome(c))
	fmt.Println(VowelCount(c))

}
