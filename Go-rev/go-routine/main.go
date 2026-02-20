package main

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


