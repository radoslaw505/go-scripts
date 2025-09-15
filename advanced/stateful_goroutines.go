package main

import (
	"fmt"
	"time"
)

// Define message types
type increment struct{}
type getCount struct {
	reply chan int
}

func counterActor(incrChan <-chan increment, getChan <-chan getCount) {
	count := 0
	for {
		select {
		case <-incrChan:
			count++
		case msg := <-getChan:
			msg.reply <- count
		}
	}
}

func main() {
	incrChan := make(chan increment)
	getChan := make(chan getCount)

	// Start the stateful goroutine
	go counterActor(incrChan, getChan)

	// Send some increments
	for i := 0; i < 5; i++ {
		incrChan <- increment{}
	}

	// Query the current count
	replyChan := make(chan int)
	getChan <- getCount{reply: replyChan}
	fmt.Println("Counter value:", <-replyChan)

	time.Sleep(1 * time.Second)
}
