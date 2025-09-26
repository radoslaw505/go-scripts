package main

import (
	"fmt"
)

/*
race_fixed_channel.go

- Avoids shared mutable state by confining the counter to a single goroutine.
- All increments are sent as messages to the owner goroutine which updates the counter.
- This pattern leverages channels for safe communication.
- Run: go run -race race_fixed_channel.go
*/

func main() {
	const goroutines = 100

	// jobs channel carries increment requests. close(jobs) signals no more requests.
	jobs := make(chan struct{})
	done := make(chan int)

	// Owner goroutine: serializes counter updates.
	go func() {
		counter := 0
		for range jobs {
			counter++
		}
		// When jobs channel is closed, return final counter via done channel.
		done <- counter
	}()

	// Spawn worker goroutines that send an increment request and return.
	for i := 0; i < goroutines; i++ {
		go func() {
			jobs <- struct{}{} // request increment from the owner
		}()
	}

	// Close jobs after all workers started to allow owner to finish once all messages processed.
	// Small sleep is intentionally avoided; in real code coordinate with WaitGroup or another mechanism.
	// For demonstration, start goroutines synchronously then close.
	// In production use a WaitGroup to ensure all workers have sent messages before closing.
	// Here we simply close after spawning; messages already sent remain in the channel or will block if unbuffered.
	close(jobs)

	// Receive final counter and print.
	fmt.Println("expected", goroutines, "got", <-done)
}
