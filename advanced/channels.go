package main

import "fmt"

func main() {
	ch := make(chan string)

	// Send data into the channel
	go func() {
		ch <- "Hello from the channel!"
	}()

	// Receive data from the channel
	msg := <-ch
	fmt.Println("Received:", msg)
}
