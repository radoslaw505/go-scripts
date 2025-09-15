package main

import (
	"fmt"
	"sync"
	"time"
)

// Simulated task
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Mark this goroutine as done
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second) // Simulate work
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup
	numWorkers := 3

	// Launch multiple workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1) // Increment counter for each worker
		go worker(i, &wg)
	}

	wg.Wait() // Wait for all workers to finish
	fmt.Println("All workers completed")
}
