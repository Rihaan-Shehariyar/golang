package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
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

func findMin(node *Node) *Node {

	for node.left != nil {
		node = node.left
	}
	return node
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

func main() {

	var root *Node

	arr := []int{1, 4, 6, 7, 8, 9}

	for _, v := range arr {
		root = Insert(root, v)
	}

	PostOrder(root)

}


func DeleteN(root *Node,value int)*Node{

 if root == nil {
	return nil
 }

 if root.data > value {
    root.right = Delete(root.right,value)
 }else if root.data < value {
	 root.left = Delete(root.left,value)
 }else{
 
 if root.left == nil {
	return root.right
 }
 if root.right == nil {
	return root.left
 }

 min := findMin(root.right)
 root.data = min.data
 root.right = Delete(root.right,min.data)

}

return root
}
