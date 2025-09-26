package main

import (
	"fmt"
	"time"
)

/*
deadlock_demo.go

- Example 1: unbuffered channel send with no receiver -> deadlock
- Example 2: cyclic waits between goroutines -> potential deadlock
- The runtime will detect global deadlock and panic when applicable.
*/

func deadlockUnbuffered() {
	ch := make(chan int)
	fmt.Println("Sending to unbuffered channel without receiver (will block)...")
	ch <- 1 // DEADLOCK: main goroutine blocks here
	fmt.Println("This line will never run")
}

func cyclicDeadlock() {
	a := make(chan struct{})
	b := make(chan struct{})

	go func() {
		<-a
		b <- struct{}{}
	}()

	go func() {
		<-b
		a <- struct{}{}
	}()

	time.Sleep(100 * time.Millisecond)
	fmt.Println("Both goroutines are blocked, program may deadlock if no progress occurs")
}

func main() {
	// Uncomment one at a time to observe behavior.
	// deadlockUnbuffered()
	cyclicDeadlock()
	time.Sleep(500 * time.Millisecond)
	fmt.Println("main returning (if you see this, no global deadlock detected yet)")
}
