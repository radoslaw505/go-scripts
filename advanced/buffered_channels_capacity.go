package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a buffered channel with capacity of 2
	ch := make(chan int, 2)

	// Start a goroutine to read from the channel with a delay
	go func() {
		for val := range ch {
			fmt.Println("📥 Received:", val)
			time.Sleep(1 * time.Second) // Simulate slow processing
		}
	}()

	// Send 4 values into the channel
	for i := 1; i <= 4; i++ {
		fmt.Println("📤 Sending:", i)
		ch <- i // First 2 sends are non-blocking, next 2 will block until space is freed
		fmt.Println("✅ Sent:", i)
	}

	// Close the channel after sending
	close(ch)

	// Wait to allow all messages to be processed
	time.Sleep(5 * time.Second)
}
