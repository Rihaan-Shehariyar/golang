package main

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

	n := len(arr)

	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = key

	}
}

func SelectionSort(arr []int) {

	n := len(arr)

	for i := 0; i < n-1; i++ {
		minIndex := i

		for j := 1; j < n; j++ {
			if arr[minIndex] > arr[j] {
				minIndex = j
			}
		}

		arr[i], arr[minIndex] = arr[minIndex], arr[i]

	}

}
