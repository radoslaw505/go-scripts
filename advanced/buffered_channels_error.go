package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Simulated task that may succeed or fail
func doWork(id int) (string, error) {
	delay := time.Duration(rand.Intn(500)) * time.Millisecond
	time.Sleep(delay)

	// 30% chance of failure
	if rand.Float32() < 0.3 {
		return "", errors.New(fmt.Sprintf("Worker %d failed", id))
	}

	return fmt.Sprintf("Worker %d completed in %v", id, delay), nil
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Buffered channels for results and errors
	resultCh := make(chan string, 5)
	errorCh := make(chan error, 5)

	numWorkers := 5

	// Launch workers
	for i := 1; i <= numWorkers; i++ {
		go func(id int) {
			result, err := doWork(id)
			if err != nil {
				errorCh <- err // Send error to error channel
				return
			}
			resultCh <- result // Send result to result channel
		}(i)
	}

	// Collect results and errors
	for i := 0; i < numWorkers; i++ {
		select {
		case res := <-resultCh:
			fmt.Println("✅", res)
		case err := <-errorCh:
			fmt.Println("❌", err)
		}
	}

	fmt.Println("🏁 All workers processed")
}
