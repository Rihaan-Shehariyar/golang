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

func Search(root *Node, value int) bool {
	if root == nil {
		return false
	}

	if root.value == value {
		return true
	}

	if value < root.value {
		return Search(root.left, value)
	}

	return Search(root.right, value)

}

// func Delete(root *Node, value int) *Node {
//    if root==nil{
//   return nil
// }

//  if value < root.value{
//   root.left =  Delete(root.left,value)
// } else if value > root.value {
// 	root.right = Delete(root.right,value)
// }else{
//  if root.left == nil {
// 	return root.right
//  }
// if root.right == nil{
//    return root.left
// }
// }

// }

func Inorder(root *Node) {
	if root == nil {
		return
	}

	Inorder(root.left)
	fmt.Println(root.value)
	Inorder(root.right)

}

func PreOrder(root *Node) {
	if root == nil {
		return
	}

	fmt.Println(root.value, " ")
	PreOrder(root.left)
	PreOrder(root.right)

}

func PostOrder(root *Node) {

	if root == nil {
		return
	}

	PreOrder(root.left)
	PostOrder(root.right)
	fmt.Println(root.value, " ")

}

func main() {

	var root *Node
	root = Insert(root, 10)
	root = Insert(root, 20)

	fmt.Println(Search(root, 30))

	// Inorder(root)

	// root := &Node{value: 10}

	// root.left = &Node{value: 50}
	// root.right = &Node{value: 3}

	// root.left.left = &Node{value: 7}
	// root.left.right = &Node{value: 1}

	// Inorder(root)

}
