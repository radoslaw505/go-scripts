package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
race_fixed_atomic.go

- Fixes the race by using atomic operations for the counter.
- Best for simple counters or flags where full mutex semantics are unnecessary.
- Run: go run -race race_fixed_atomic.go to confirm no race reports.
*/

func main() {
	var wg sync.WaitGroup

	var counter int64 // must be 64-bit aligned on some platforms
	const goroutines = 100
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func() {
			// Atomic increment avoids data race without a mutex.
			atomic.AddInt64(&counter, 1)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("expected", goroutines, "got", atomic.LoadInt64(&counter))
}
