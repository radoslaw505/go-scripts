package main

import (
	"fmt"
)

func main() {
	// Create a buffered channel with capacity of 2
	ch := make(chan string, 2)

	// Send two messages into the channel
	ch <- "Hello"
	ch <- "Buffered Channel"

	// Receive and print the messages
	fmt.Println(<-ch) // Output: Hello
	fmt.Println(<-ch) // Output: Buffered Channel

	// Note: If you try to send a third message without receiving,
	// it will block because the buffer is full.
}
