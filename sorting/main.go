package main

import "fmt"

func BubbleSort(arr []int) {

	n := len(arr)

	for i := 0; i < n-1; i++ {

		swapped := false

		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}

		if !swapped {
			break
		}

	}

}

func InsertionSort(arr []int) {

	for i := 1; i < len(arr); i++ {

		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = key

	}

}

func main() {

	arr := []int{5, 6, 1, 7, 8, 8}
	BubbleSort(arr)
	fmt.Println(arr)

}
