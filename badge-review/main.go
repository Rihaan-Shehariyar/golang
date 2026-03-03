package main

import (
	// "fmt"
	// "sync"
)

// func worker(id int, jobs <-chan int, wg *sync.WaitGroup){

//   defer wg.Done()

//   fmt.Println("Woker",id,"job",jobs)

// cte,json be
// }

// type User struct{
//   Name string
// }

// func (T )Print(){
   
// }


// type Node struct{
//   Data int
//   Next *Node
// }

// type LinkedList struct{
//   Head *Node
// }

// func Center(l *LinkedList)*Node{

//  // 1,2,3,4,5
//    left := l.Head
//    right := l.Head.Next

//   for right!=nil && right.Next != nil{
//     left = left.Next
//     right = right.Next.Next
// }
   
//  return left

// }



// func Palindrome(str string)bool{
 
//   string := []string{str}
//  reversed := 


//  for i :=len(str)-1 ; i <= 0  ; i++{

//    reversed = append()
 
// }



// }



// for i := 0; i < 3; i++ {}

// Salman Faris
// 13:59
// package main
// import "fmt"

// func main() {
//     for i := 0; i < 3; i++ {
//         go func() {
//             fmt.Println(i)
//         }()
//     }
// }

// Salman Faris
// 14:10
// "Hello 👋  Bro 😎 ! "

// Salman Faris
// 14:17
// package main
// import "fmt"

// type User struct{}

// func main() {
//     var u *User = nil
//     var i interface{} = u

//     if i == nil {
//         fmt.Println("nil")
//     } else {
//         fmt.Println("not nil")
//     }
// }

// Salman Faris
// 14:23
// Fan-in , Fan-out

// Salman Faris
// 14:27
// switch v := i.(type)

// Salman Faris
// 14:28
// Type Assertions

// Salman Faris
// 14:42
// JSONB in PostgreSQL

// Salman Faris
// 14:51
// ALTER TABLE dexgem_projects
// ALTER COLUMN growth TYPE integer;
// CREATE OR REPLACE FUNCTION set_growth()
// RETURNS TRIGGER AS $$
// BEGIN
//     IF NEW.initial_price > 0 AND NEW.published_price > 0 THEN
//         NEW.growth := ROUND(
//             NEW.initial_price::numeric / 
//             NULLIF(NEW.published_price, 0)
//         );
//     ELSE
//         NEW.growth := 0;
//     END IF;

//     RETURN NEW;
// END;
// $$ LANGUAGE plpgsql;
// CREATE TRIGGER trigger_set_growth
// BEFORE INSERT OR UPDATE
// ON dexgem_projects
// FOR EACH ROW
// EXECUTE FUNCTION set_growth();

// Salman Faris
// 14:53
// Growth            int                        `json:"growth" gorm:"->"`

// Salman Faris
// 15:19
// runes := []rune(s)


// $$ LANGUAGE plpgsql;
// CREATE TRIGGER trigger_set_growth
// BEFORE INSERT OR UPDATE
// ON dexgem_projects
// FOR EACH ROW
// EXECUTE FUNCTION set_growth();

// Salman Faris
// 14:53
// Growth            int                        `json:"growth" gorm:"->"`

// Salman Faris
// 15:19
// runes := []rune(s)

// Salman Faris
// 15:24
// Disadvantage of Golang
// Go Routine – Print Numbers
// Order with WaitGroup
// Range loop with string
// Polymorphism
// Unsafe Package
// Fan-in / Fan-out
// Type Switch / Type Assertion

// Salman Faris
// 15:25
// mbedding Interfaces
// CTE
// JSONB in PostgreSQL
// Run

// Salman Faris
// 15:27
// WITH cte_name AS (
//     SELECT ...
// )
// SELECT * FROM cte_name;