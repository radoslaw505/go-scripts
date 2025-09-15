package main

import (
	"fmt"
	"sync"
	"time"
)

// Worker function: receives tasks and sends results
func worker(id int, tasks <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		fmt.Printf("Worker %d processing task %d\n", id, task)
		time.Sleep(time.Second) // Simulate work
		results <- task * task  // Send result
	}
}

func main() {
	numWorkers := 3
	numTasks := 5

	tasks := make(chan int, numTasks)
	results := make(chan int, numTasks)
	var wg sync.WaitGroup

	// Start workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, tasks, results, &wg)
	}

	// Send tasks
	for j := 1; j <= numTasks; j++ {
		tasks <- j
	}
	close(tasks) // No more tasks

	// Wait for all workers to finish
	wg.Wait()
	close(results)

	// Collect results
	for res := range results {
		fmt.Println("Result:", res)
	}
}
