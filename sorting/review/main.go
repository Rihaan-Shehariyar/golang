package main

import "fmt"

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

// func Twosums(arr []int, target int) []int {

// 	b := make(map[int]int)

// 	for i, x := range arr {
// 		b[x] = i
// 	}

// }



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
		root.right = Insert(root.right, value)
	} else {

		root.left = Insert(root.left, value)

	}

	return root

}

func Inorder(root *Node){

 if root==nil {
	return
 }

	Inorder(root.left)
	fmt.Println(root.data)
	Inorder(root.right)

}

func main(){

 var root *Node
 arr := []int{2,3,5,67,3}

 for _,x := range arr{

  Insert(root,x)

}
 fmt.Println("Inorder")
 Inorder(root)

}

type Heap struct {
	arr []int
}

func (h *Heap) Insert(value int) {

	h.arr = append(h.arr, value)
	index := len(h.arr) - 1

	h.HeapifyUp(index)
}

func (h *Heap) HeapifyUp(index int) {

	for index > 0 {

		parent := (index - 1) / 2

		if h.arr[index] > h.arr[parent] {
			h.arr[parent], h.arr[index] = h.arr[index], h.arr[parent]
			index = parent
		} else {
			break
		}

	}

}

func (h *Heap) ExtractMax() int {

	max := h.arr[0]
	n := len(h.arr) - 1
	h.arr[0] = h.arr[n]
	h.arr = h.arr[:n]

	h.HeapifyDown(0)
	return max

}

func (h *Heap) HeapifyDown(i int) {

	size := len(h.arr)
	for {

		left := 2*i + 1
		right := 2*i + 2
		largest := i

		if left < size && h.arr[left] > h.arr[largest] {
			largest = left
		}

		if right < size && h.arr[right] > h.arr[largest] {
			largest = right
		}

		if largest == i {
			break
		}

		h.arr[i], h.arr[largest] = h.arr[largest], h.arr[i]

	}

}

func main() {

	arr := Heap{}

	//  arr.arr=[]int{}

	//  for _,y := range x{
	//   arr.HeapifyUp(y)
	// }

	arr.Insert(10)
	arr.Insert(30)
	arr.Insert(50)
	arr.Insert(50)

	arr.ExtractMax()
	arr.ExtractMax()

	fmt.Println(arr.arr)

}
