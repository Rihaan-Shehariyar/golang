package main

// import "fmt"

func QuickSort(arr []int, low, high int) {

	if low < high {
		pivotIndex := Partition(arr, low, high)
		QuickSort(arr, low, pivotIndex-1)

		QuickSort(arr, pivotIndex+1, high)
	}

}

func Partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1

}

// func main() {
// 	arr := []int{8, 3, 1, 7, 0, 10, 2}

// 	fmt.Println("Before:", arr)

// 	QuickSort(arr, 0, len(arr)-1)

// 	fmt.Println("After :", arr)
// }
