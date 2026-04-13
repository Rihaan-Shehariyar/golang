package main

import (
	"context"
	"fmt"
	"time"
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



func worker(ctx context.Context) {

	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Work finished")

	default:
		fmt.Println("Work cancelled : ", ctx.Err())

	}
}


func main(){

 ctx,cancel := context.WithTimeout(context.Background(),2 *time.Second)
 defer cancel()

 go worker(ctx)

 time.Sleep(4 * time.Second)

}