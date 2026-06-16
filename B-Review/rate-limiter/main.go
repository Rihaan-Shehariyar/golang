package main

import (
	"fmt"
	"sync"
)

// import (
// 	"fmt"
// 	"time"

// 	"golang.org/x/time/rate"
// )

// func main() {
// 	limiter := rate.NewLimiter(2, 4)

// 	for i := 0; i < 10; i++ {
// 		if limiter.Allow() {
// 			fmt.Println("allowed", i)
// 		} else {
// 			fmt.Println("blocked", i)
// 		}

// 		time.Sleep(200 * time.Millisecond)

// 	}
// }

// func worker(ctx context.Context) {

// 	select {
// 	case <-time.After(3 * time.Second):
// 		fmt.Println("Work finished")

// 	default:
// 		fmt.Println("Work cancelled : ", ctx.Err())

// 	}
// }


func worker(jobwoker int,jobs <-chan int , wg *sync.WaitGroup){

 defer wg.Done()
 
 for job := range jobs{
 
  fmt.Printf("Worker %d proccesed job %d ",jobwoker,job)
 
}

}


func main(){

  jobs := make(chan int)


 var wg sync.WaitGroup

 numworkers := 3

 for i := 1; i <= numworkers; i++ {
	wg.Add(1)
  go worker(numworkers,jobs,&wg)
 }


 for j := 1; j <=5; j++ {
	jobs <- j
 }

close(jobs)
 

wg.Wait()

}