package main

import (
	"fmt"
	"sync"
	"time"
)

// import (
// 	"context"
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func Worker(ctx context.Context, id int, jobs <-chan int, wg *sync.WaitGroup) {

// 	defer wg.Done()

// 	for {
// 		select {
// 		case job, ok := <-jobs:

// 			if !ok {
// 				fmt.Println("Worker", id, "Stopping")
// 				return
// 			}

// 			fmt.Println("worker", id, "Proccesing", job)
// 			time.Sleep(time.Second)

// 		case <-ctx.Done():
// 			fmt.Println("Worker", id, "Cancelled")
// 			return
// 		}
// 	}

// }

// func main() {

// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	jobs := make(chan int)
// 	var wg sync.WaitGroup

// 	workCount := 3

// 	for i := 0; i < workCount; i++ {

// 		wg.Add(1)
// 		go Worker(ctx, i, jobs, &wg)

// 	}

// 	go func() {
// 		for j := 0; j <= 7; j++ {
// 			jobs <- j
// 		}

// 		close(jobs)
// 	}()

// 	wg.Wait()

// 	fmt.Println("All workers Finished")

// }

// func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	for job := range jobs {
// 		fmt.Println("worker", id, "processing", job)
// 		time.Sleep(time.Second)
// 	}

// }

// func main() {
// 	jobs := make(chan int, 10)

// 	var wg sync.WaitGroup

// 	for i := 1; i <= 3; i++ {
// 		wg.Add(1)
// 		go worker(i, jobs, &wg)
// 	}

// 	for j := 1; j <= 10; j++ {

// 		jobs <- j

// 	}

// 	close(jobs)
// 	wg.Wait()

// }

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {

	defer wg.Done()

	for job := range jobs {
		fmt.Println("Worker", id, "Processing job ", job)
		time.Sleep(time.Second)
	}

}

func main(){


 var wg sync.WaitGroup

 jobs := make(chan int,10)

 for i:=1 ; i<=3 ; i++{
 wg.Add(1)
 go worker(i,jobs,&wg)
 
}

 for j:=1 ; j<=10 ; j++{
  jobs <- j
}

 close(jobs)

 wg.Wait()

 
}




func StartWorkerPool(numWorkers int,jobs []int){

 jobch := make(chan int)

 var wg sync.WaitGroup

  for i:=1 ; i < numWorkers ; i++{
 
 wg.Add(1)

 go func(worker int) {
   defer wg.Done()

 for job := range jobch{
   fmt.Printf("Worker %d processed %d job",worker,job)
}
 
 }(i)
 
 for _,job := range jobs{
   jobch <- job
}

close(jobch)

wg.Wait()
}
 
}