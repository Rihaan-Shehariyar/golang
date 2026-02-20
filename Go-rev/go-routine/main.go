package main

import (
	"fmt"
	// "time"
)

// import (
// 	"fmt"
// 	"runtime"
// 	"sync"
// 	"time"
// )

// func job(id int, wg *sync.WaitGroup) {

// 	defer wg.Done()

// 	fmt.Println("Job Started : ", id)
// 	time.Sleep(1 * time.Second)
// 	fmt.Println("Job Finished : ", id)

// }
// func main() {
// 	var wg sync.WaitGroup

// 	for i := 1; i <= 3; i++ {

// 		wg.Add(1)
// 		go job(i, &wg)

// 	}

// 	wg.Wait()
// runtime.GOMAXPROCS(3)
// 	fmt.Println("All Jobs Completed")

// }

// func worker(id int, ch chan int) {
// 	ch <- id * 2
// }

// func main() {

// 	ch := make(chan int)

// 	for i := 1; i <= 3; i++ {
// 		go worker(i, ch)
// 	}

// 	for i := 1; i <= 3; i++ {
// 		fmt.Println(<-ch)
// 	}

// }


// func worker(id int,ch chan string){
 
// //  time.Sleep(time.Duration(id)*time.Second)
//  ch <- fmt.Sprintf("Worker %d done",id)
// }

//  func main(){

//  ch1 :=make(chan string)
//  ch2 := make(chan string)

//  go worker(1,ch1)
//  go worker(2,ch2)


 
//  select {

//  case r := <-ch1:
//   fmt.Println(r)
//  case r := <-ch2:
//   fmt.Println(r)
// }
// }
