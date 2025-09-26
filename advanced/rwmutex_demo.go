package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
rwmutex_demo.go

- Demonstrates typical RWMutex usage with many readers and occasional writers.
- Shows how readers can run concurrently and writers are exclusive.
- Run: go run rwmutex_demo.go
*/

func main() {
	var mu sync.RWMutex
	counter := 0
	var wg sync.WaitGroup

	rand.Seed(time.Now().UnixNano())

	// Start many readers
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 20; j++ {
				// Readers use RLock to allow concurrency among readers.
				mu.RLock()
				_ = counter // read shared state
				// simulate read work
				time.Sleep(time.Duration(rand.Intn(5)) * time.Millisecond)
				mu.RUnlock()
				time.Sleep(5 * time.Millisecond)
			}
			fmt.Println("reader", id, "done")
		}(i)
	}

	// Start a few writers
	for w := 0; w < 3; w++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for k := 0; k < 5; k++ {
				// Writers must take exclusive Lock
				mu.Lock()
				counter++
				fmt.Printf("writer %d incremented counter to %d\n", id, counter)
				// simulate write work
				time.Sleep(10 * time.Millisecond)
				mu.Unlock()
				time.Sleep(30 * time.Millisecond)
			}
			fmt.Println("writer", id, "done")
		}(w)
	}

	wg.Wait()
	fmt.Println("final counter:", counter)
}
