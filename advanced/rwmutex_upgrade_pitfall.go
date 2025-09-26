package main

import (
	"fmt"
	"sync"
	"time"
)

/*
rwmutex_upgrade_pitfall.go

- Demonstrates the pitfall of trying to "upgrade" from RLock to Lock.
- RUnlock followed by Lock is not atomic; other writers can intervene.
- This reproducer shows how a writer can sneak in between RUnlock and Lock.
- Run: go run rwmutex_upgrade_pitfall.go
*/

func main() {
	var mu sync.RWMutex
	value := 1
	var wg sync.WaitGroup

	// Reader that wants to upgrade to writer
	wg.Add(1)
	go func() {
		defer wg.Done()
		mu.RLock()
		v := value
		fmt.Println("reader observed value:", v)
		mu.RUnlock()

		// Attempt to upgrade: this is not atomic!
		time.Sleep(50 * time.Millisecond) // allow writer to run
		mu.Lock()
		fmt.Println("reader upgraded and setting value from", value, "to 100")
		value = 100
		mu.Unlock()
	}()

	// Writer that runs in between
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(20 * time.Millisecond)
		mu.Lock()
		fmt.Println("writer acquired lock and sets value to 2")
		value = 2
		mu.Unlock()
	}()

	wg.Wait()
	fmt.Println("final value:", value)
}
