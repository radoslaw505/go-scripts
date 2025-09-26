package main

import (
	"fmt"
	"sync"
	"time"
)

/*
deadlock_mutex.go

- Demonstrates mutex-related deadlocks and how to avoid them.
- selfDeadlock shows that recursive locking blocks; keep mutex usage non-recursive.
- lockOrderDeadlock shows potential deadlock when acquiring mutexes in opposite orders.
- lockOrderFixed demonstrates consistent lock ordering to avoid deadlock.
*/

func selfDeadlock() {
	var mu sync.Mutex
	fmt.Println("demonstrating self-deadlock: locking twice in same goroutine")
	mu.Lock()
	fmt.Println("first lock acquired")
	// mu.Lock() // uncommenting this line would block forever because Mutex is not recursive
	mu.Unlock()
}

func lockOrderDeadlock() {
	var muA, muB sync.Mutex
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		muA.Lock()
		fmt.Println("goroutine1 locked A")
		time.Sleep(50 * time.Millisecond)
		muB.Lock()
		fmt.Println("goroutine1 locked B")
		muB.Unlock()
		muA.Unlock()
	}()

	go func() {
		defer wg.Done()
		muB.Lock()
		fmt.Println("goroutine2 locked B")
		time.Sleep(50 * time.Millisecond)
		muA.Lock()
		fmt.Println("goroutine2 locked A")
		muA.Unlock()
		muB.Unlock()
	}()

	wg.Wait()
	fmt.Println("lockOrderDeadlock finished (if this printed, no deadlock occurred this run)")
}

func lockOrderFixed() {
	var muA, muB sync.Mutex
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		muA.Lock()
		fmt.Println("goroutine1 locked A")
		time.Sleep(50 * time.Millisecond)
		muB.Lock()
		fmt.Println("goroutine1 locked B")
		muB.Unlock()
		muA.Unlock()
	}()

	go func() {
		defer wg.Done()
		// same order as above to avoid deadlock
		muA.Lock()
		fmt.Println("goroutine2 locked A")
		time.Sleep(10 * time.Millisecond)
		muB.Lock()
		fmt.Println("goroutine2 locked B")
		muB.Unlock()
		muA.Unlock()
	}()

	wg.Wait()
	fmt.Println("lockOrderFixed finished safely")
}

func main() {
	selfDeadlock()
	// lockOrderDeadlock() // may deadlock depending on scheduling; use with caution
	lockOrderFixed()
}
