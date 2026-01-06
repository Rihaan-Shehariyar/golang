package main

import (
	"encoding/json"
	"net/http"
)

// import "fmt"

// // import "fmt"

// // type Node struct {
// // 	data int
// // 	prev *Node
// // 	next *Node
// // }

// // type doublyLinked struct {
// // 	head *Node
// // }

// // func (l *doublyLinked) insert(val int) {

// // 	newNode := &Node{data: val}
// // 	if l.head == nil {

// // 		l.head = newNode
// // 		return
// // 	}

// // 	curr := l.head

// // 	for curr.next != nil {

// // 		curr = curr.next

// // 	}

// // 	curr.next = newNode
// // 	newNode.prev = curr

// // }

// // func (l *doublyLinked) delete(val int) {

// // 	if l.head == nil {

// // 		fmt.Println("Empty")
// // 	}

// //    if l.head.data == val{

// //    l.head = l.head.next
// // }

// //   curr:=l.head

// //    for curr.next!=nil && curr.next.data != val{

// //     curr = curr.next
// // }

// //   if curr.next ==nil{

// //    return
// // }

// //  curr.next = curr.next.next
// //  curr.next.prev = curr
// // }

// // func(l doublyLinked)print(){

// //    curr := l.head

// //    for curr != nil{

// //      fmt.Printf("%d ->",curr.data)
// //      curr = curr.next
// // }

// //  fmt.Println("nil")
// // }

// // func main(){

// //    var l doublyLinked

// //  l.insert(1)
// //  l.insert(2)
// //  l.insert(3)

// // // l.delete(2)

// // l.print()
// // }

// // type Node struct{

// //   data int
// //   next *Node

// // }

// // type stack struct{

// //     top *Node
// // }

// // func (s *stack)push(val int){

// //     newNode := &Node{data: val}

// // if s.top ==nil{

// //    s.top = newNode
// //    return
// // }

// //    newNode.next = s.top
// //    s.top = newNode

// // }

// // func (s *stack)pop(){

// //   if s.top == nil{
// //    return
// //  }

// //   s.top = s.top.next

// // }

// // func(s stack)peek()int{

// //   return s.top.data
// // }
// // func (s *stack)print(){

// //    curr:=s.top

// // for curr!=nil{

// //   fmt.Printf("%d ->",curr.data)
// //   curr = curr.next
// // }
// // fmt.Println("nil")
// // }

// // func main(){

// //     var s stack

// // s.push(10)
// // s.push(20)
// // s.push(30)

// // s.pop()

// // s.print()
// // }

// type node struct{

//   data int
//    next *node

// }

// type queue struct{

//   front *node
//   rear *node

// }

// func (q *queue)enqueue(val int){

//  newNode := &node{data: val}

//    if q.front == nil{

//     q.front = newNode
//     q.rear = newNode
//     return
// }

// q.rear.next = newNode
// q.rear = newNode

// }

// func (q *queue)dequeue(){

//    if q.front == nil{

//     return
// }

// q.front = q.front.next

// if q.front == nil{

//    q.front = nil
//   q.rear = nil
// }
// }

// func(q queue)print(){

// curr:=q.front

// for curr!=nil{

//    fmt.Printf("%d ->",curr.data)
//    curr = curr.next
// }

// fmt.Println("nil")
// }

// func main(){

//   var q queue

// q.enqueue(1)
// q.enqueue(2)
// q.enqueue(3)
// q.enqueue(4)

// q.dequeue()
// q.dequeue()
// q.print()

// }



func helloResponse(w http.ResponseWriter, r *http.Request){
 
  if r.Method != http.MethodGet{
    http.Error(w,"Method Not allowed",http.StatusMethodNotAllowed)
   return
}  
  
 w.Header().Set("Content-Type","application/json")

  response := map[string]string{
   "message" : "Hello",
}

 json.NewEncoder(w).Encode(response)

}


func main(){  
   
 http.HandleFunc("/hello",helloResponse)
 http.ListenAndServe(":8080",nil)

}


