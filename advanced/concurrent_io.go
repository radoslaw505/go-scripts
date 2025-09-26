package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

/*
concurrent_io.go

- Demonstrates a worker pool for I/O-like tasks.
- Uses context cancellation for graceful shutdown when SIGINT or SIGTERM is received.
- Workers listen on a jobs channel, perform simulated I/O, and send results on a results channel.
- WaitGroup ensures we wait for workers to finish before closing results.
*/

func simulateIO(ctx context.Context, id int) (int, error) {
	// Simulate variable I/O latency per job.
	// Use time.After in a select so we can also return early if ctx is canceled.
	delay := time.Duration(200+id*50) * time.Millisecond
	select {
	case <-time.After(delay):
		// Normal completion: return a simple computed result.
		return id * 2, nil
	case <-ctx.Done():
		// If context was canceled while waiting, return the context error.
		return 0, ctx.Err()
	}
}

// worker performs jobs from the jobs channel and sends results to the results channel.
// It respects ctx for cancellation and uses wg.Done() to signal completion.
func worker(ctx context.Context, id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	// Continuously process incoming jobs until jobs channel is closed or context is canceled.
	for {
		select {
		case <-ctx.Done():
			// Context canceled: stop processing and return to let main wait on wg.
			fmt.Printf("worker %d received cancel\n", id)
			return
		case n, ok := <-jobs:
			if !ok {
				// jobs channel closed: nothing left to do.
				return
			}

			// Perform the simulated I/O work.
			res, err := simulateIO(ctx, n)
			if err != nil {
				// Job was canceled while running; log and continue to next job.
				fmt.Printf("worker %d canceled job %d\n", id, n)
				continue
			}

			// Attempt to send the result. If ctx is canceled while sending, exit.
			select {
			case results <- res:
				// Successfully sent result.
			case <-ctx.Done():
				// Cancellation detected while trying to send result.
				return
			}
		}
	}
}

func main() {
	// Create a cancellable root context. Calling cancel will notify all workers to stop.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // ensure resources are freed when main returns

	// Set up signal handling for graceful shutdown (SIGINT / SIGTERM).
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	// Start a goroutine to wait for an OS signal and then cancel the context.
	go func() {
		s := <-sig
		fmt.Println()
		fmt.Println("shutdown signal received:", s)
		cancel()
	}()

	const tasks = 12
	// Buffered channels sized for the number of tasks help avoid blocking producers immediately.
	jobs := make(chan int, tasks)
	results := make(chan int, tasks)

	var wg sync.WaitGroup
	const workers = 4

	// Start worker goroutines.
	wg.Add(workers)
	for w := 0; w < workers; w++ {
		go worker(ctx, w, jobs, results, &wg)
	}

	// Producer: enqueue jobs. It respects ctx so it stops enqueueing if canceled.
	go func() {
		for i := 0; i < tasks; i++ {
			select {
			case jobs <- i:
				// job enqueued
			case <-ctx.Done():
				// stop producing if canceled
				break
			}
		}
		// No more jobs will be sent.
		close(jobs)
	}()

	// Collector: wait for all workers to finish, then close results channel.
	go func() {
		wg.Wait()      // wait for worker goroutines to exit
		close(results) // signal to main that collection is complete
	}()

	// Read and print results until results channel is closed.
	for r := range results {
		fmt.Println("result:", r)
	}

	// Final message after graceful shutdown and all results processed.
	fmt.Println("all done, exiting")
}
