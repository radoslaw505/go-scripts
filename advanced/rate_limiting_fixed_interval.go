package main

import (
	"fmt"
	"time"
)

func main() {
	// Limit to 1 operation every 500ms
	rateLimit := time.NewTicker(500 * time.Millisecond)
	defer rateLimit.Stop()

	for i := 1; i <= 5; i++ {
		<-rateLimit.C // Wait for tick
		fmt.Println("Operation", i, "at", time.Now())
	}
}
