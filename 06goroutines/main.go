// package main

// import "fmt"

// func main() {

// 	ch := make(chan string)

// 	go func() {
// 		ch <- "Hello from goroutine"
// 	}()

// 	msg := <-ch

// 	fmt.Println(msg)

// }

package main



import (
    "fmt"
    "sync"
)

var count = 0
var mu sync.Mutex

func increment() {

    mu.Lock()
    count++
    mu.Unlock()
}

func main() {
    var wg sync.WaitGroup

    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            increment()
            wg.Done()
        }()
    }

    wg.Wait()
    fmt.Println(count)
}
