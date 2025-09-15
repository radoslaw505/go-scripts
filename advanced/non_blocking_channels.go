package main

import (
	"fmt"
	"time"
)

func main() {
	// Example 1: Non-blocking receive from an empty channel
	ch1 := make(chan int)

	select {
	case val := <-ch1:
		fmt.Println("Received from ch1:", val)
	default:
		// Executes immediately because ch1 has no data
		fmt.Println("ch1 is empty, no value received")
	}

	// Example 2: Non-blocking send to a full buffered channel
	ch2 := make(chan string, 1)
	ch2 <- "first" // Fill the buffer

	select {
	case ch2 <- "second":
		fmt.Println("Sent 'second' to ch2")
	default:
		// Executes because ch2 is full
		fmt.Println("ch2 is full, could not send 'second'")
	}

	// Example 3: Non-blocking receive from multiple channels
	ch3 := make(chan string)
	ch4 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch3 <- "message from ch3"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch4 <- "message from ch4"
	}()

	// Poll channels without blocking
	for i := 0; i < 3; i++ {
		select {
		case msg := <-ch3:
			fmt.Println("Received:", msg)
		case msg := <-ch4:
			fmt.Println("Received:", msg)
		default:
			// No channels ready yet
			fmt.Println("No messages available")
		}
		time.Sleep(500 * time.Millisecond)
	}

	// Example 4: Non-blocking with timeout using time.After
	ch5 := make(chan int)

	select {
	case val := <-ch5:
		fmt.Println("Received from ch5:", val)
	case <-time.After(1 * time.Second):
		// Executes if no message arrives within 1 second
		fmt.Println("Timeout: no message received from ch5")
	}
}
