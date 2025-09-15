package main

import (
	"fmt"
	"time"
)

// simulateTask simulates a task that sends a message after a delay.
// It takes a name, a delay duration, and a send-only channel.
func simulateTask(name string, delay time.Duration, ch chan<- string) {
	time.Sleep(delay)                      // Simulate work by sleeping
	ch <- fmt.Sprintf("%s finished", name) // Send result to channel
}

func main() {
	// Create two unbuffered channels for communication
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Start two goroutines that simulate tasks
	go simulateTask("Task A", 2*time.Second, ch1)
	go simulateTask("Task B", 1*time.Second, ch2)

	// Use select to wait for whichever task finishes first
	select {
	case msg1 := <-ch1:
		fmt.Println("Received from ch1:", msg1)
	case msg2 := <-ch2:
		fmt.Println("Received from ch2:", msg2)
	case <-time.After(3 * time.Second):
		// Timeout case: executes if no other case is ready within 3 seconds
		fmt.Println("Timeout: no task completed in time")
	}
}
