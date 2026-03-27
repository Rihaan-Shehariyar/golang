package main

import "fmt"

// Heap Sort implementation
func HeapSort(arr []int) {
	n := len(arr)

	// Build Max Heap
	// Heap Property:
	// In a MAX HEAP, parent >= children
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// Extract elements from heap
	for i := n - 1; i > 0; i-- {
		// Move current root to end
		arr[0], arr[i] = arr[i], arr[0]

		// Restore heap property
		heapify(arr, i, 0)
	}
}

// Heapify subtree rooted at index i
func heapify(arr []int, n int, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest)
	}
}

func main() {
	arr := []int{4, 10, 3, 5, 1}
	HeapSort(arr)
	fmt.Println(arr)
}
