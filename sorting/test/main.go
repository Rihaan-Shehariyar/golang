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

func SelectionSort(arr []int) {

	size := len(arr)

	for i := 0; i < size-1; i++ {
		minIndex := i

		for j := i + 1; j < size; j++ {
			if arr[minIndex] > arr[j] {
				minIndex = j
			}
		}

		arr[i], arr[minIndex] = arr[minIndex], arr[i]
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

type Node struct {
	data  int
	left  *Node
	right *Node
}

func Insert(root *Node, value int) *Node {

	if root == nil {
		return &Node{data: value}
	}

	if root.data > value {
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

	if value > root.data {
		return Search(root.right, value)
	}

	return Search(root.left, value)

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

		minNode := minNode(root.right)
		root.data = minNode.data
		root.right = Delete(root.right, minNode.data)

	}

	return root

}

func minNode(root *Node) *Node {

	for root.left != nil {
		root = root.left
	}

	return root

}

type Heap struct {
	arr []int
}

func (h *Heap) HeapifyUp(value int) {

	h.arr = append(h.arr, value)
	index := len(h.arr) - 1

	for index > 0 {

		parent := (index - 1) / 2

		if h.arr[index] > h.arr[parent] {
			h.arr[index], h.arr[parent] = h.arr[parent], h.arr[index]
			index = parent
		} else {
			break
		}

	}

}

func (h Heap) ExtractMax() int {

	if len(h.arr) == 0 {
		return -1
	}

	n := len(h.arr) - 1
	maxVale := h.arr[0]
	h.arr[0] = h.arr[n]
	h.arr = h.arr[:n]

	h.HeapifyDown(0)
	return maxVale

}

func (h *Heap) HeapifyDown(index int) {

	size := len(h.arr)

	for {

		left := 2*index + 1
		right := 2*index + 2
		largest := index

		if left < size && h.arr[left] > h.arr[largest] {
			largest = left
		}

		if right < size && h.arr[right] > h.arr[largest] {
			largest = right
		}

		if largest == index {
			break
		}

		h.arr[index], h.arr[largest] = h.arr[largest], h.arr[index]
		index = largest

	}

}

func BinarySearch(arr []int, target int) int {

	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := low + (high-low)/2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}

	}
	return -1
}

func PreOrder(root *Node) {
	if root == nil {
		return
	}

	fmt.Println(root.data)
	PreOrder(root.left)
	PreOrder(root.right)

}

func PostOrder(root *Node) {
	if root == nil {
		return
	}

	PostOrder(root.left)
	PostOrder(root.right)
	fmt.Println(root.data)

}

func Inorder(root *Node) {
	if root == nil {
		return
	}

	Inorder(root.left)
	fmt.Println(root.data)
	Inorder(root.right)

}
