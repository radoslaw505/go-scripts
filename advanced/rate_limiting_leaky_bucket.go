package main

import (
	"fmt"
	"time"
)

var bucket = make(chan int, 3) // bucket capacity
const leakRate = time.Second   // one request per second

func leak() {
	for range time.Tick(leakRate) {
		select {
		case req := <-bucket:
			fmt.Println("Processed request", req, "at", time.Now())
		default:
			// No requests to process
		}
	}
}

func main() {
	go leak()

	for i := 1; i <= 10; i++ {
		select {
		case bucket <- i:
			fmt.Println("Accepted request", i)
		default:
			fmt.Println("Rejected request", i, "(bucket full)")
		}
		time.Sleep(300 * time.Millisecond)
	}

	time.Sleep(5 * time.Second) // Let leak process remaining requests
}
