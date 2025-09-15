package main

import (
	"fmt"
	"time"

	"golang.org/x/time/rate"
)

func main() {
	limiter := rate.NewLimiter(2, 5) // 2 tokens/sec, burst size 5

	for i := 1; i <= 10; i++ {
		if limiter.Allow() {
			fmt.Println("Request", i, "allowed at", time.Now())
		} else {
			fmt.Println("Request", i, "denied at", time.Now())
		}
		time.Sleep(300 * time.Millisecond)
	}
}
