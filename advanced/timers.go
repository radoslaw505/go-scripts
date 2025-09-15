package main

import (
	"fmt"
	"time"
)

func main() {
	// Example 1: Basic timer using time.NewTimer
	fmt.Println("Example 1: NewTimer")
	timer1 := time.NewTimer(2 * time.Second)
	fmt.Println("Waiting for timer1 to fire...")
	<-timer1.C // Blocks until timer fires
	fmt.Println("timer1 fired!")

	// Example 2: Using time.After for one-shot delay
	fmt.Println("\nExample 2: time.After")
	fmt.Println("Waiting for 1 second...")
	<-time.After(1 * time.Second)
	fmt.Println("time.After fired!")

	// Example 3: Stopping a timer before it fires
	fmt.Println("\nExample 3: Stop a timer")
	timer2 := time.NewTimer(3 * time.Second)
	stopped := timer2.Stop()
	if stopped {
		fmt.Println("timer2 stopped before firing")
	} else {
		fmt.Println("timer2 already fired or stopped")
	}

	// Example 4: Resetting a timer
	fmt.Println("\nExample 4: Reset a timer")
	timer3 := time.NewTimer(1 * time.Second)
	<-timer3.C // Wait for it to fire

	// Reset the timer to fire again after 2 seconds
	timer3.Reset(2 * time.Second)
	fmt.Println("timer3 reset, waiting again...")
	<-timer3.C
	fmt.Println("timer3 fired after reset")

	// Example 5: Using timer with select for timeout
	fmt.Println("\nExample 5: Select with timer")
	ch := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		ch <- "data received"
	}()

	select {
	case msg := <-ch:
		fmt.Println("Received:", msg)
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout: no data received")
	}
}
