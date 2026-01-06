package main

import "fmt"

// import (
// 	"fmt"

// 	// "golang.org/x/tools/go/analysis/passes/nilfunc"
// )

// type node struct {
// 	data int
// 	next *node
// }

// type linkedList struct {
// 	head   *node
// 	length int
// }

// func (l *linkedList) insert(value int) {

// 	newValue := &node{data: value}

// 	if l.head == nil {

// 		l.head = newValue
// 		l.length++
// 		return

// 	}

// 	cur := l.head
// 	for cur.next != nil {
// 		cur = cur.next
// 	}

// 	cur.next = newValue
// 	l.length++
// }

// func (l *linkedList) search(target int) bool {

// 	curr := l.head

// 	for curr != nil {
// 		if curr.data == target {
// 			return true
// 		}
//      curr = curr.next

// 	}
// 	return false
// }

// func (l *linkedList) delete(value int) bool {
// 	if l.head == nil {
// 		return false
// 	}

// 	if l.head.data == value {

// 		l.head = l.head.next
// 		l.length--
// 		return true
// 	}

// 	curr := l.head

// 	for curr.next != nil && curr.next.data != value {

// 		curr = curr.next
// 	}
// 	if curr.next == nil {

// 		return false
// 	}
// 	curr.next = curr.next.next
// 	l.length--
// 	return true
// }

// func (l *linkedList) print() {

// 	curr := l.head

//   for curr != nil{

//   fmt.Printf("%d-> ",curr.data)
//   curr = curr.next
// }

// }

// func main(){

//  var list linkedList

// list.insert(10)
// list.insert(20)
// list.insert(30)
// list.insert(40)

// fmt.Println("After Inserts : ")
// list.print()

// fmt.Println("search result : ",list.search(50))

// list.delete(20)
// fmt.Println("After deletion")
// list.print()

// }


type node struct{

  data int
  next *node
}

type linkedList struct{
  
   head *node  

}

func (l *linkedList)append(value int){
  
    newNode := &node{data: value}

   if l.head == nil{
   l.head = newNode
    return
}
curr := l.head

for curr.next!=nil{
  curr = curr.next
}
curr.next = newNode
}

func (l *linkedList)fromArray(arr []int){
   
  for _,v := range arr{
 
    l.append(v) 
}
}

func (l *linkedList) print(){

   curr:=l.head

  for curr !=nil{
   
   fmt.Printf("%d ->",curr.data)
   curr = curr.next
}
fmt.Println("nil")
 }

func main(){

  arr := []int{1,2,3,4,5}

var list linkedList

  list.fromArray(arr)

list.print()

}