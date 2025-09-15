package main

import (
	"fmt"
	"time"
)

func worker(id int, ch chan int) {
	fmt.Printf("Worker %d: Sending data...\n", id)
	ch <- id // blocks until main receives
	fmt.Printf("Worker %d: Sent data\n", id)
}

func main() {
	ch := make(chan int) // unbuffered channel

	// Launch 3 workers
	for i := 1; i <= 3; i++ {
		go worker(i, ch)
	}

	// Receive from channel 3 times
	for i := 1; i <= 3; i++ {
		val := <-ch // blocks until a worker sends
		fmt.Printf("Main: Received %d\n", val)
	}

	time.Sleep(1 * time.Second) // allow goroutines to finish
}
