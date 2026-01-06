package main

import (
	
	"fmt"
)

// import "fmt"

// type queue struct {
// 	items []int
// }

// func (q *queue) enqueue(value int) {

// 	q.items = append(q.items, value)
// }

// func (q *queue) dequeue() int {

// 	if len(q.items) == 0 {

// 		return -1
// 	}
// 	top := q.items[0]
// 	q.items = q.items[1:]
// 	return top
// }

// func (q *queue) peek() int {
// 	if len(q.items) == 0 {

// 		return -1
// 	}

// 	return q.items[0]

// }

// func main() {

// 	var q queue

// 	q.enqueue(10)
// 	q.enqueue(20)
// 	q.enqueue(30)

// fmt.Print(q.peek())

// 	q.dequeue()

// 	fmt.Println(q.items)

// }

// type Queue struct{

// data [5]int
// front int
// rear int
// size int
// capacity int

// }

// func newQueue() *Queue{

// return &Queue{

//     front: 0,
//     rear: 0,
//     size: 0,
//     capacity: 5,

// }
// }

// func (q *Queue) enqueue(value int){

//   if q.size == q.capacity{
//    fmt.Println("capacity full")
// return
// }

// q.data[q.rear] = value
// q.rear = (q.rear+1)%q.capacity
// q.size++
// }

// func (q *Queue) deQueue()int{

//  if q.size ==0{

//    return -1
// }

// val:= q.data[q.front]
// q.front =(q.front+1)%q.capacity
// q.size--
// return  val
// }

// func (q *Queue) print(){

//   if q.size ==0{

//    return
// }

// i := q.front
// count:=q.size
// for count>0{

//    fmt.Printf("%d ",q.data[i])
//    i = (i+1)%q.capacity
//    count --

// }
// fmt.Println()

// }

// func main(){

//  q := newQueue()

// q.enqueue(10)
// q.enqueue(20)
// q.enqueue(30)
// q.enqueue(40)
// q.enqueue(50)
// q.print()

// fmt.Println(q.deQueue())

// q.print()

// }


type node struct{
    
  data int
  next *node
  
}

type linkedList struct{

    front *node
    rear *node
    size int
   
}

func (q *linkedList)enqueue(value int){

    newNode := &node{data: value}

if q.rear == nil{

  q.rear = newNode
  q.front = newNode
  q.size ++
  return

}

q.rear.next = newNode
q.rear = newNode
q.size++

}

func (q *linkedList)dequeue()int{

   if q.front == nil{
     
   return -1
  
}

  value := q.front.data

  q.front = q.front.next
  

  if q.front == nil{

   q.rear = nil
 
}
 
  q.size--

  return value
}

func (q *linkedList)peek()int{

  if q.front == nil{
     
   return -1
  }

 return q.front.data
}

func (q linkedList)print(){
   
    if q.front == nil{
    
     fmt.Println("Empty Linked List") 
     return   
} 
   curr := q.front
  
  for curr != nil {

    fmt.Printf("%d ",curr.data)
    curr = curr.next
}

fmt.Println("nil")

  
}


func main(){

  var q linkedList

  q.enqueue(10)
  q.enqueue(20)
  q.enqueue(30)
  q.enqueue(40)


  q.print()
q.dequeue()
q.print()


q.print()
}