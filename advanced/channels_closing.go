package main

import (
	"fmt"
	"time"
)

// producer sends a few values into the channel and then closes it.
func producer(ch chan<- int) {
	for i := 1; i <= 3; i++ {
		ch <- i // Send values to channel
		time.Sleep(500 * time.Millisecond)
	}
	close(ch) // Close the channel to signal no more values will be sent
}

func main() {
	// Create an unbuffered channel
	ch := make(chan int)

	// Start the producer in a separate goroutine
	go producer(ch)

	// Example 1: Receiving values using range until channel is closed
	fmt.Println("Reading from channel using range:")
	for val := range ch {
		fmt.Println("Received:", val)
	}

	// Example 2: Receiving from a closed channel returns zero value
	closedCh := make(chan string, 1)
	closedCh <- "hello"
	close(closedCh)

	val1 := <-closedCh // "hello"
	val2 := <-closedCh // "" (zero value for string)

	fmt.Println("\nReading from closed buffered channel:")
	fmt.Println("First read:", val1)
	fmt.Println("Second read (after close):", val2)

	// Example 3: Using comma-ok idiom to detect channel closure
	ch2 := make(chan int, 1)
	ch2 <- 42
	close(ch2)

	val, ok := <-ch2
	fmt.Printf("\nComma-ok idiom:\nReceived: %d, Open: %v\n", val, ok)

	val, ok = <-ch2
	fmt.Printf("Received: %d, Open: %v\n", val, ok)

	// Example 4: Attempting to send on a closed channel (uncomment to see panic)
	/*
	   ch3 := make(chan int)
	   close(ch3)
	   ch3 <- 99 // ❌ PANIC: send on closed channel
	*/

	// Example 5: Attempting to close a channel twice (uncomment to see panic)
	/*
	   ch4 := make(chan int)
	   close(ch4)
	   close(ch4) // ❌ PANIC: close of closed channel
	*/
}
