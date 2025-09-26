package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"
)

/*
parallel_cpu.go

- Demonstrates CPU-bound work distributed across worker goroutines.
- Uses runtime.GOMAXPROCS to allow true parallel execution on multiple OS threads.
- Uses context cancellation for graceful shutdown when SIGINT or SIGTERM is received.
- Workers compute a CPU-heavy function (naive recursive fib) and send results to a collector.
*/

// naive recursive fib to create CPU bound workload.
// This is intentionally inefficient to make CPU work measurable.
func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

// worker reads job integers from jobs channel, computes fib(n), and sends results.
// It checks ctx so it can stop work promptly when cancellation is requested.
func worker(ctx context.Context, id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			// Context canceled: log and return so wg.Done() is observed.
			fmt.Printf("worker %d stopping\n", id)
			return
		case n, ok := <-jobs:
			if !ok {
				// No more jobs: exit cleanly.
				return
			}

			// Compute the heavy function. This blocks the goroutine and uses CPU.
			res := fib(n)

			// Send result but exit early if canceled while sending.
			select {
			case results <- res:
			case <-ctx.Done():
				return
			}
		}
	}
}

func main() {
	// Allow Go to use all available CPU cores for parallel execution.
	// runtime.GOMAXPROCS(runtime.NumCPU()) is the default for modern Go versions,
	// but setting it explicitly documents intent.
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Create cancellable context for graceful shutdown.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Set up signal handling for SIGINT and SIGTERM.
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	// Goroutine to handle OS signal and cancel context so workers stop.
	go func() {
		s := <-sig
		fmt.Println()
		fmt.Println("shutdown signal received:", s)
		cancel()
	}()

	const tasks = 8
	jobs := make(chan int, tasks)
	results := make(chan int, tasks)

	var wg sync.WaitGroup
	const workers = 4

	// Start worker goroutines that perform CPU bound work.
	wg.Add(workers)
	for w := 0; w < workers; w++ {
		go worker(ctx, w, jobs, results, &wg)
	}

	// Producer: enqueue CPU tasks. It respects ctx cancellation.
	go func() {
		for i := 0; i < tasks; i++ {
			select {
			case jobs <- 35: // use a value that creates noticeable CPU work
			case <-ctx.Done():
				// stop producing if we were canceled
				break
			}
		}
		// Close jobs to indicate no further tasks.
		close(jobs)
	}()

	// Collector: wait for workers then close results to signal completion.
	go func() {
		wg.Wait()
		close(results)
	}()

	// Track elapsed time to observe speedup from parallelism.
	start := time.Now()
	count := 0

	// Consume results until results channel is closed.
	for r := range results {
		_ = r // result value unused here, but could be logged or aggregated
		count++
		fmt.Printf("collected %d/%d\n", count, tasks)
	}

	// Print summary including elapsed wall-clock time.
	fmt.Println("processed tasks", count)
	fmt.Println("elapsed", time.Since(start))
}
