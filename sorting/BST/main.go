package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

func Insert(root *Node, value int) *Node {

	if root == nil {
		return &Node{value: value}
	}

	if value < root.value {
		root.left = Insert(root.left, value)
	} else if value > root.value {
		root.right = Insert(root.right, value)
	}

	return root

}

func Inorder(root *Node){
  if root == nil {
	return
  }
 
 Inorder(root.left)
 fmt.Println(root.value)
 Inorder(root.right)
 
}

func main() {

	var root *Node
	root = Insert(root, 10)
	root = Insert(root, 20)

	Inorder(root)
}



