package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Example 1: Using context.TODO()
	fmt.Println("Example 1: context.TODO()")

	// context.TODO is used when you're not sure which context to use yet.
	// It's a placeholder and should be replaced with a real context later.
	ctxTODO := context.TODO()

	// You can still use TODO context like any other context
	go func(ctx context.Context) {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("Work done with TODO context")
		case <-ctx.Done():
			fmt.Println("TODO context canceled:", ctx.Err())
		}
	}(ctxTODO)

	time.Sleep(2 * time.Second)

	// Example 2: Using context.WithCancel to cancel a goroutine
	fmt.Println("Example 2: WithCancel")
	ctx1, cancel1 := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				// This block runs when cancel() is called
				fmt.Println("Goroutine canceled:", ctx.Err())
				return
			default:
				fmt.Println("Working...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}(ctx1)

	time.Sleep(2 * time.Second)
	cancel1() // Cancel the context to stop the goroutine
	time.Sleep(1 * time.Second)

	// Example 3: Using context.WithTimeout to enforce a deadline
	fmt.Println("\nExample 3: WithTimeout")
	ctx2, cancel2 := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel2() // Always call cancel to release resources

	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Finished work")
	case <-ctx2.Done():
		// This block runs when the timeout expires
		fmt.Println("Timeout occurred:", ctx2.Err())
	}

	// Example 4: Passing values through context
	fmt.Println("\nExample 4: WithValue")
	type key string
	userKey := key("userID")

	// Attach a value to the context
	ctx3 := context.WithValue(context.Background(), userKey, 42)
	processRequest(ctx3, userKey)

	// Example 5: Using context.WithDeadline()
	fmt.Println("\nExample 5: context.WithDeadline()")

	// Set an absolute deadline 3 seconds from now
	deadline := time.Now().Add(3 * time.Second)
	ctxDeadline, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel() // Always cancel to release resources

	// Simulate a goroutine doing work
	go func(ctx context.Context) {
		select {
		case <-time.After(5 * time.Second):
			fmt.Println("Finished long task")
		case <-ctx.Done():
			// This will run if the deadline is exceeded before task completes
			fmt.Println("Deadline exceeded:", ctx.Err())
		}
	}(ctxDeadline)

	time.Sleep(4 * time.Second)
}

// processRequest retrieves a value from the context
func processRequest(ctx context.Context, k interface{}) {
	val := ctx.Value(k)
	if val != nil {
		fmt.Println("Retrieved value from context:", val)
	} else {
		fmt.Println("No value found for key:", k)
	}
}
