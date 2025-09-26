package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	// Atomic counter
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
	fmt.Println("Final atomic counter value:", atomic.LoadInt64(&counter)) // Safe read

	// Normal to compare
	var t_wg sync.WaitGroup
	t_value := 0

	for i := 0; i < 1000; i++ {
		t_wg.Add(1)
		go func() {
			defer t_wg.Done()
			t_value++ // Unsafe increment
		}()
	}

	t_wg.Wait()
	fmt.Println("Final normal counter value:", t_value) // Unsafe read
}
