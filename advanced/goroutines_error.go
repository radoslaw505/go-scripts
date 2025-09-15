package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// riskyGoroutine simulates a task that may fail randomly
func riskyGoroutine(id int) error {
	delay := time.Duration(rand.Intn(500)) * time.Millisecond
	time.Sleep(delay) // simulate work

	// 30% chance of failure
	if rand.Float32() < 0.3 {
		return errors.New(fmt.Sprintf("Goroutine %d failed", id))
	}

	fmt.Printf("✅ Goroutine %d completed after %v\n", id, delay)
	return nil
}

func main() {
	rand.Seed(time.Now().UnixNano()) // seed random generator
	var wg sync.WaitGroup            // WaitGroup to wait for all goroutines

	num := 5 // number of goroutines to launch
	for i := 1; i <= num; i++ {
		wg.Add(1) // increment WaitGroup counter
		go func(id int) {
			defer wg.Done() // decrement counter when done
			if err := riskyGoroutine(id); err != nil {
				fmt.Println("❌", err)
			}
		}(i) // pass loop variable as argument
	}

	wg.Wait() // block until all goroutines finish
	fmt.Println("🏁 All goroutines finished")
}
