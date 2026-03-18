package main

import (
	"fmt"
	"strings"
)

func main() {

	text := "Hello Good Morning Hello Good Night GO GO GO"

	freq := make(map[string]int)

	words := strings.Fields(text)

	for _, word := range words {
		freq[word]++
	}

	fmt.Println(freq)

}
