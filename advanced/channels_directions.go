package main

import "fmt"

// ping sends a message into the 'pings' channel.
// The channel is send-only: chan<- string
func ping(pings chan<- string, msg string) {
	pings <- msg // Send the message into the channel
}

// pong receives a message from 'pings' and sends it into 'pongs'.
// 'pings' is receive-only: <-chan string
// 'pongs' is send-only: chan<- string
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings // Receive message from 'pings'
	pongs <- msg   // Send message into 'pongs'
}

func main() {
	// Create two buffered channels with capacity 1
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	// Send a message using the ping function
	ping(pings, "passed message")

	// Transfer the message from pings to pongs
	pong(pings, pongs)

	// Receive and print the final message from pongs
	fmt.Println(<-pongs) // Output: passed message
}
