package main

import (
	"fmt"
	"time"
)

//для саблима, что бы не забыть ctrl + f3 и ctrl+d для переименования переменных.
func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}

	close(requests)

	limiter := time.Tick(time.Millisecond * 200)
	for req := range requests {
		//дальше не пойдет пока не получит в канал лимитера значение, в нашем случаи каждые 200 миллисек
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)

	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()

	burstyRequest := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequest <- i
	}
	close(burstyRequest)

	for req := range burstyRequest {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
