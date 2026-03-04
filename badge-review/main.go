package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup){

   defer wg.Done()

   fmt.Println("Woker",id,"job",jobs)

}

 type User struct{
   Name string
 }
 func Print[T any](a T)(){
   fmt.Println(a)
 }


 type Node struct{
   Data int
   Next *Node
 }

 type LinkedList struct{
   Head *Node
 }

 func Center(l *LinkedList)*Node{

//   1,2,3,4,5

    left := l.Head
    right := l.Head.Next

   for right!=nil && right.Next != nil{
     left = left.Next
     right = right.Next.Next
 }
   
  return left

 }