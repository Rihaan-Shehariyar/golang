package main

import (
	"fmt"
	"sync"
)

var pool = sync.Pool{
	New: func() any {
		fmt.Println("creating a new Buffer")
		return make([]byte, 0, 10)
	},
}

func worker(id int, wg *sync.WaitGroup) {

	defer wg.Done()

	buff := pool.Get().([]byte)

	buff = append(buff, byte(id))

	fmt.Println("worker", id, "buff:", buff)

	buff = buff[:0]
	pool.Put(buff)
}

func main() {

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()

}
