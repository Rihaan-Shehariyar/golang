package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(value int) {

	s.items = append(s.items, value)
}

func (s *Stack) Pop() int {

	if len(s.items) == 0 {

		return -1
	}

	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return top
}

func (s *Stack) peek()int{

 if len(s.items) == 0{
  
   return -1
}
  
  return s.items[len(s.items)-1]
}

func main() {

	s := Stack{}

	s.Push(10)
	s.Push(20)
	s.Push(30)
	s.Push(40)
	s.Push(50)
	s.Push(60)


	fmt.Println(s)

s.Pop()
s.Pop()
s.Pop()
s.Pop()


fmt.Println(s)
fmt.Println(s.peek())

}