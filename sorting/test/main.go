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

	for i := 0; i < n-1; i++ {
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

type Node struct {
	data  int
	right *Node
	left  *Node
}

func Insert(root *Node, value int) *Node {
	if root == nil {
		return &Node{data: value}
	}

	if value < root.data {
		root.left = Insert(root.left, value)
	} else {
		root.right = Insert(root.right, value)
	}

	return root

}

func Search(root *Node, value int) bool {

	if root == nil {
		return false
	}

	if root.data == value {
		return true
	}

	if value < root.data {
		return Search(root.left, value)
	}

	return Search(root.right, value)

}

func Delete(root *Node, value int) *Node {

	if root == nil {
		return nil
	}

	if value < root.data {
		root.left = Delete(root.left, value)
	} else if value > root.data {
		root.right = Delete(root.right, value)
	} else {
		if root.left == nil {
			return root.right
		}

		if root.right == nil {
			return root.left
		}

		minNode := findMin(root.right)
		root.data = minNode.data
		root.right = Delete(root.right, minNode.data)

	}
	return root

}

func findMin(root *Node) *Node {
	for root.left != nil {
		root = root.left
	}
	return root
}

type Heap struct {
	arr []int
}

func (h *Heap) InsetMin(value int) {

	h.arr = append(h.arr, value)
	index := len(h.arr) - 1

	for index > 0 {

		parent := (index - 1) / 2

		if h.arr[parent] > h.arr[index] {
			h.arr[parent], h.arr[index] = h.arr[index], h.arr[parent]
			index = parent
		} else {
			break
		}

	}

}
func (h *Heap) ExtractMIn() int {

	if len(h.arr) == 0 {
		return -1
	}

	min := h.arr[0]
	n := len(h.arr) - 1
	h.arr[0] = h.arr[n]
	h.arr = h.arr[:n]

	h.HeapifyDown(0)

	return min

}

func (h *Heap) HeapifyDown(index int) {

	size := len(h.arr)

	for {

		left := 2*index + 1
		right := 2*index + 2
		smallest := index

		if left < size && h.arr[left] < h.arr[smallest] {
			smallest = left
		}

		if right < size && h.arr[right] < h.arr[smallest] {
			smallest = right
		}

		if index == smallest {
			break
		}

		h.arr[index], h.arr[smallest] = h.arr[smallest], h.arr[index]
		index = smallest

	}

}

func BubbleSortu(arr []int) {

	n := len(arr)

	for i := 0; i < n-1; i++ {
		swapped := false

		for j := 0; j < n-i-1; j++ {
			if arr[j+1] < arr[j] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}

		if !swapped {
			break
		}

	}

}

func InsertionSorts(arr []int) {

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
