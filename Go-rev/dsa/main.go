package main

import "fmt"

// type Node struct {
// 	Value int
// 	Next  *Node
// }

// func main() {

// 	n1 := &Node{Value: 10}
// 	n2 := &Node{Value: 20}
//     n3 := &Node{Value: 30}

// 	n1.Next = n2

// 	curr := n1

// 	for curr.Next != nil {
// 		curr = curr.Next
// 	}

//  curr.Next = n3

// }

// func InsertHead(head *Node,value int) *Node	{

//  newNode := &Node{Value: 30}
//  newNode.Next = head

//  head = newNode

//  return head

// }

type Node struct {
	Value int
	Next  *Node
}

type linkedList struct{
  head *Node
}

func(l *linkedList)Insert(value int){
   
 newNode := &Node{Value: value}
  
if l.head == nil{
   l.head = newNode
   return
}

 curr := l.head
 for curr.Next != nil{
  curr = curr.Next
}

 curr.Next = newNode

}

func (l *linkedList)InsertHead(value int){
  newNode := &Node{Value: value}

 newNode.Next = l.head
 l.head = newNode
  
}

func (l *linkedList)FromArray(arr []int){
  for _, v := range arr{
  l.Insert(v)
}
}

func (l *linkedList)Print(){
 curr := l.head
 
 for curr !=nil{
  fmt.Printf("%d -> ",curr.Value)
  curr = curr.Next
}

 fmt.Println("nil")
 
}

func main(){

 var list linkedList
 var fromar linkedList
  
 arr := []int{1,3,5,7}

 list.Insert(10)
 list.Insert(20)
 list.Insert(30)
 list.Insert(40)
 list.Insert(50)
 list.Insert(60)
list.InsertHead(0)

list.Print()
fromar.FromArray(arr)
 fromar.Print()

}