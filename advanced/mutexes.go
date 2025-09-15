package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	var counter int
	var wg sync.WaitGroup

	// Launch 1000 goroutines to increment counter
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()   // Acquire lock
			counter++   // Critical section
			mu.Unlock() // Release lock
		}()
	}

	wg.Wait()
	fmt.Println("Final counter value:", counter)
}
