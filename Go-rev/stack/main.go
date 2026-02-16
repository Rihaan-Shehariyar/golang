package main

// import "fmt"

// type Stack struct {
// 	items []int
// }

// func (s *Stack) Push(val int) {

// 	s.items = append(s.items, val)

// }
// func (s *Stack) Pop() (int, bool) {
// 	if len(s.items) == 0 {
// 		return 0, false
// 	}

// 	topIndex := len(s.items) - 1
// 	val := s.items[topIndex]

// 	s.items = s.items[:topIndex]
// 	return val, true

// }

// func (s *Stack) Peek() (int, bool) {

// 	if len(s.items) == 0 {
// 		return 0, false
// 	}

// 	return s.items[len(s.items)-1], true

// }

// func (s Stack) Print() {
// 	fmt.Println(s.items)
// }

// func main() {

// 	stack := Stack{}

// 	stack.Push(10)
// 	stack.Push(20)
// 	stack.Push(30)
// 	stack.Push(40)

// 	fmt.Println("Stack after pushes:")
// 	stack.Print()

//  stack.Pop()
//  stack.Print()

// }

// Linked List

type Node struct {
	Value int
	Next  *Node
}

type Stack struct {
	Top *Node
}

func (s *Stack) Push(val int) {

	newNode := &Node{Value: 10}

	newNode.Next = s.Top
	s.Top = newNode

}

func(s *Stack)Pop() int{
 
  if s.Top == nil {
	return 0
  }

  val := s.Top.Value
 s.Top = s.Top.Next
  
 return val
  
}
