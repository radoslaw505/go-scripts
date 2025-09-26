package main

import (
	"fmt"
	"sync"
)

/*
race_fixed_mutex.go

- Fixes the race by using sync.Mutex to protect the shared counter.
- Run: go run -race race_fixed_mutex.go to confirm no race reports.
*/

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	counter := 0
	const goroutines = 100
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func() {
			// Lock around the critical section to serialize access.
			mu.Lock()
			counter++
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("expected", goroutines, "got", counter)
}
