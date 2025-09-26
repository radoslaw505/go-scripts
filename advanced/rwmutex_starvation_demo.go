package main

import (
	"fmt"
	"sync"
	"time"
)

/*
rwmutex_starvation_demo.go

- Demonstrates potential writer starvation when readers continuously acquire RLock.
- This program spawns continuous readers and a writer that waits; writer may be delayed.
- Note: Go's RWMutex attempts to avoid starvation, but heavy reader churn can still delay writers.
- Run: go run rwmutex_starvation_demo.go
*/

func main() {
	var mu sync.RWMutex
	stop := make(chan struct{})
	var wg sync.WaitGroup

	// Continuous readers
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for {
				select {
				case <-stop:
					return
				default:
					mu.RLock()
					// short read
					time.Sleep(1 * time.Millisecond)
					mu.RUnlock()
				}
			}
		}(i)
	}

	// Writer that tries to acquire exclusive lock periodically
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			start := time.Now()
			mu.Lock()
			fmt.Println("writer acquired lock after", time.Since(start))
			// simulate write
			time.Sleep(5 * time.Millisecond)
			mu.Unlock()
			time.Sleep(50 * time.Millisecond)
		}
	}()

	// Run for a short time then stop readers
	time.Sleep(1 * time.Second)
	close(stop)
	wg.Wait()
	fmt.Println("done")
}
