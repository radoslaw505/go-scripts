package main

import (
	"fmt"
	"sync"
)

/*
race_demo.go

- Demonstrates a data race on a shared counter.
- Run with the race detector to see the race: go run -race race_demo.go
- The result is nondeterministic and often less than the expected 100.
*/

func main() {
	var wg sync.WaitGroup

	// Shared variable accessed by many goroutines without synchronization.
	counter := 0

	const goroutines = 100
	wg.Add(goroutines)

	// Launch goroutines that increment the counter concurrently.
	for i := 0; i < goroutines; i++ {
		go func() {
			// Intentionally unsynchronized increment causes data race.
			counter++ // read-modify-write is not atomic
			wg.Done()
		}()
	}

	// Wait for all goroutines to finish.
	wg.Wait()

	// The printed counter is likely less than goroutines due to lost updates.
	fmt.Println("expected", goroutines, "got", counter)
}
