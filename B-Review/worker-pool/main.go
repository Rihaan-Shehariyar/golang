package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Worker(ctx context.Context, id int, jobs <-chan int, wg *sync.WaitGroup) {

	defer wg.Done()

	for {
		select {
		case job, ok := <-jobs:

			if !ok {
				fmt.Println("Worker", id, "Stopping")
				return
			}

			fmt.Println("worker", id, "Proccesing", job)
			time.Sleep(time.Second)

		case <-ctx.Done():
			fmt.Println("Worker", id, "Cancelled")
			return
		}
	}

}

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	jobs := make(chan int)
	var wg sync.WaitGroup

	workCount := 3

	for i := 0; i < workCount; i++ {

		wg.Add(1)
 	go Worker(ctx, i, jobs, &wg)

	}

	go func() {
		for j := 0; j <= 7; j++ {
			jobs <- j
		}

		close(jobs)
	}()

	wg.Wait()

	fmt.Println("All workers Finished")

}
