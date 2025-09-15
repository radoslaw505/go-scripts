package main

import (
	"fmt"
	"time"
)

func printMessage(msg string, delay time.Duration) {
	time.Sleep(delay)
	fmt.Println(msg)
}

func main() {
	fmt.Println("Main function started")

	go printMessage("First goroutine finished", 1*time.Second)
	go printMessage("Second goroutine finished", 2*time.Second)
	go printMessage("Third goroutine finished", 3*time.Second)

	// Wait long enough for all goroutines to complete
	time.Sleep(4 * time.Second)

	fmt.Println("Main function ended")
}
