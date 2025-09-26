package main

import (
	"bytes"
	"fmt"
	"sync"
	"time"
)

/*
pool_example.go

- Demonstrates a sync.Pool for reusing byte buffers to reduce allocations.
- Shows safe pattern: get, reset, use, put.
- Also demonstrates that Pool may discard objects under GC (not shown here).
*/

var bufPool = sync.Pool{
	New: func() interface{} {
		// allocate buffer with capacity to avoid repeated allocations
		return bytes.NewBuffer(make([]byte, 0, 1024))
	},
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	// Acquire buffer from pool
	b := bufPool.Get().(*bytes.Buffer)
	// Reset before use to ensure clean state
	b.Reset()

	// use buffer
	fmt.Fprintf(b, "worker-%d time=%v", id, time.Now().Format("15:04:05.000"))
	fmt.Println(b.String())

	// Put back into pool for reuse
	bufPool.Put(b)
}

func main() {
	var wg sync.WaitGroup
	const workers = 10
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go worker(i, &wg)
	}
	wg.Wait()

	// Example: Get a buffer when pool empty returns New()
	b := bufPool.Get().(*bytes.Buffer)
	fmt.Println("got buffer of len/cap:", b.Len(), b.Cap())
	bufPool.Put(b)
}
