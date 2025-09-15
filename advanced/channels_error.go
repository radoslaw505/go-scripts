package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// riskySend simulates a task that sends either a result or an error
func riskySend(id int, resultCh chan<- string, errorCh chan<- error) {
	delay := time.Duration(rand.Intn(500)) * time.Millisecond
	time.Sleep(delay) // simulate work

	// 30% chance of failure
	if rand.Float32() < 0.3 {
		errorCh <- errors.New(fmt.Sprintf("Sender %d failed", id))
		return
	}

	resultCh <- fmt.Sprintf("Sender %d succeeded after %v", id, delay)
}

func main() {
	rand.Seed(time.Now().UnixNano()) // seed random generator

	resultCh := make(chan string) // channel for successful results
	errorCh := make(chan error)   // channel for errors

	num := 5 // number of goroutines to launch
	for i := 1; i <= num; i++ {
		go riskySend(i, resultCh, errorCh)
	}

	// Collect results and errors from channels
	for i := 0; i < num; i++ {
		select {
		case res := <-resultCh:
			fmt.Println("✅", res)
		case err := <-errorCh:
			fmt.Println("❌", err)
		}
	}

	fmt.Println("📬 All messages processed")
}
