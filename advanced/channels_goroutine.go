package main

import (
	"fmt"
	"math/rand"
	"time"
)

// worker simulates a task that takes a random amount of time,
// then sends a message to the channel when done.
func worker(id int, ch chan<- string) {
	// Generate a random delay between 0–500 milliseconds
	delay := time.Duration(rand.Intn(500)) * time.Millisecond

	// Simulate work by sleeping
	time.Sleep(delay)

	// Send a completion message to the channel
	ch <- fmt.Sprintf("Worker %d finished after %v", id, delay)
}

func main() {
	// Seed the random number generator with current time
	rand.Seed(time.Now().UnixNano())

	// Create a channel for string messages
	ch := make(chan string)

	// Define how many workers to launch
	numWorkers := 5

	// Launch multiple worker goroutines
	for i := 1; i <= numWorkers; i++ {
		go worker(i, ch) // Start each worker concurrently
	}

	// Receive and print messages from all workers
	for i := 1; i <= numWorkers; i++ {
		msg := <-ch // Wait for a message from the channel
		fmt.Println("✅", msg)
	}

	// Final message after all workers have reported
	fmt.Println("🏁 All workers reported")
}
