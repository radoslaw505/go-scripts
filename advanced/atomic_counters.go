package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64
	var wg sync.WaitGroup

	// Launch 1000 goroutines to increment the counter
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1) // Safe concurrent increment
		}()
	}

	wg.Wait()
	fmt.Println("Final counter value:", atomic.LoadInt64(&counter)) // Safe read
}
