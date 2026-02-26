package main

import (
	"fmt"
	"time"

	"golang.org/x/time/rate"
)

func main() {
	limiter := rate.NewLimiter(2,4)


 for i:=0; i< 10;i++{
 if limiter.Allow() {
	fmt.Println("allowed",i)
 } else{
 fmt.Println("blocked",i)
}

 time.Sleep(200*time.Millisecond)

}
}