package main

import "fmt"

type node struct {
	data int
	prev *node
	next *node
}

type doublyList struct {
	head *node
	tail *node
}

func (l *doublyList) append(value int) {

	newNode := &node{data: value}

	if l.head == nil {

		l.head = newNode
		l.tail = newNode
		return
	}

	l.tail.next = newNode
	newNode.prev = l.tail
	l.tail = newNode

}

func (l *doublyList) insert(value int) {

	newNode := &node{data: value}

	if l.head == nil {

		l.head = newNode
		l.tail = newNode
		return
	}
	l.head.prev = newNode
	newNode.next = l.head
	l.head = newNode
}

func (l *doublyList) deleteTail() {

	if l.head == nil {

		return
	}

	if l.head == l.tail {

		l.head = nil
		l.tail = nil
		return
	}

	l.tail = l.tail.prev
	l.tail.next = nil
}

func (l *doublyList) deleteHead() {

	if l.head == nil {

		return
	}

	if l.head == l.tail {

		l.head = nil
		l.tail = nil
		return
	}

	l.head = l.head.next
	l.head.prev = nil
}

func (l doublyList) printhead() {

	curr := l.head

	for curr != nil {

		fmt.Printf("%d ->",curr.data)
      curr = curr.next
       
	}
fmt.Println("nil")
}

func (l doublyList)printtail(){
    curr:=l.tail
     
    for curr != nil{
    
     fmt.Printf("%d ->",curr.data)
     curr = curr.prev
}

fmt.Println("nil")
} 

func main() {

    var list doublyList

  list.insert(1)
  list.insert(2)
  list.insert(3)
  list.append(4)
  list.deleteHead()
  list.deleteTail()
  

   list.printhead()
   list.printtail()
  
}