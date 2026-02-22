package main

import (
	"fmt"
)

// import (
// 	"fmt"
// 	"sync"
// )

// // import "fmt"

// // "fmt"
// // "time"

// // import (
// // 	"fmt"
// // 	"runtime"
// // 	"sync"
// // 	"time"
// // )

// // func job(id int, wg *sync.WaitGroup) {

// // 	defer wg.Done()

// // 	fmt.Println("Job Started : ", id)
// // 	time.Sleep(1 * time.Second)
// // 	fmt.Println("Job Finished : ", id)

// // }
// // func main() {
// // 	var wg sync.WaitGroup

// // 	for i := 1; i <= 3; i++ {

// // 		wg.Add(1)
// // 		go job(i, &wg)

// // 	}

// // 	wg.Wait()
// // runtime.GOMAXPROCS(3)
// // 	fmt.Println("All Jobs Completed")

// // }

// // func worker(id int, ch chan int) {
// // 	ch <- id * 2
// // }

// // func main() {

// // 	ch := make(chan int)

// // 	for i := 1; i <= 3; i++ {
// // 		go worker(i, ch)
// // 	}

// // 	for i := 1; i <= 3; i++ {
// // 		fmt.Println(<-ch)
// // 	}

// // }

// // func worker(id int,ch chan string){

// // //  time.Sleep(time.Duration(id)*time.Second)
// //  ch <- fmt.Sprintf("Worker %d done",id)
// // }

// //  func main(){

// //  ch1 :=make(chan string)
// //  ch2 := make(chan string)

// //  go worker(1,ch1)
// //  go worker(2,ch2)

// //  select {

// //  case r := <-ch1:
// //   fmt.Println(r)
// //  case r := <-ch2:
// //   fmt.Println(r)
// // }
// // }

// // func main() {

// // 	odd := make(chan bool)
// // 	even := make(chan bool)

// // 	go func() {
// // 		for i := 1; i <= 10; i += 2 {

// // 			<-odd
// // 			fmt.Println("Odd:", i)
// // 			even <- true
// // 		}
// // 	}()

// // 	go func() {
// // 		for i := 2; i <= 10; i += 2 {
// // 			<-even
// // 			fmt.Println("Even:", i)
// // 			odd <- true
// // 		}
// // 	}()

// // 	odd <- true
// // 	select {}

// // }

// // func producer(ch chan int){
// // for i:=1;i<=5;i++{
// //   fmt.Println("Produced: ",i)
// //   ch <- i
// //   time.Sleep(500*time.Millisecond)
// // }
// //  close(ch)
// // }

// // func consumer(ch chan int){

// //  for job := range ch{
// //  fmt.Println("consumed:",job)
// // }

// // }

// // func main(){

// //  ch := make(chan int)

// //  go producer(ch)
// //  consumer(ch)

// // }

// var counter int
// var mu sync.Mutex

// func worker(wg *sync.WaitGroup){

//  defer wg.Done()

//  mu.Lock()
//  counter ++
//  mu.Unlock()

// }

// func main(){

//  var wg sync.WaitGroup

//  for i:=0;i<1000;i++{
//  wg.Add(1)
//  go worker(&wg)
// }

// wg.Wait()
// fmt.Println("Final COunter",counter)

// }


// var wg sync.WaitGroup

// func main(){
//   odd :=make(chan bool)
//   even := make(chan bool)

//  go func() {
// 	for i := 1; i < 11; i+=2 {
//        <-odd
//      fmt.Println("Odd : ",i)
//       	even <- true
// 	}
//  }()

//  go func() {
// 	for i:=2; i<11; i +=2{
//     <-even
//    fmt.Println("even:",i)
//     odd <-true
// } 
//  }()
 
//   odd<-true
//   select{}
// } 




// func main(){
 
//  var wg sync.WaitGroup
 
//  num := make(chan int)
//  wg.Add(5)
//  for i := 1; i <6 ; i++ {
    
// 	 go func() {
//        defer wg.Done()
// 		num <- i
// 	 }()

//   fmt.Print(<-num)
 
//  }
//  wg.Wait()
// }




// func main(){

//  arr := []int{1,2,3,4}
//  b := make(chan int)
//  go func() {
// 	for _,x := range arr{
//   b <- x *2
//  }

// close(b)
//  }()
 
//   for y := range b{
//  fmt.Print(y) 
// }
// }


func main(){
 sl := []int{10, 20, 30, 40, 50}
 ch := make(chan int)
 sum := 0

 go func() {
	for _,x := range sl{
  ch <- x
}
 close(ch)
 }()

 for x := range ch{
  sum = sum + x
}

fmt.Println(sum)

}