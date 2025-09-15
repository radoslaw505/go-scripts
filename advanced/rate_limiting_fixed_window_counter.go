package main

import (
	"fmt"
	"time"
)

var requestCount int
var windowStart = time.Now()

const limit = 5
const windowSize = time.Second * 10

func allowRequest() bool {
	now := time.Now()
	if now.Sub(windowStart) > windowSize {
		windowStart = now
		requestCount = 0
	}
	requestCount++
	return requestCount <= limit
}

func main() {
	for i := 1; i <= 10; i++ {
		if allowRequest() {
			fmt.Println("Request", i, "allowed")
		} else {
			fmt.Println("Request", i, "denied")
		}
		time.Sleep(1 * time.Second)
	}
}
