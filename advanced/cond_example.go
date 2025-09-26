package main

import (
	"fmt"
	"sync"
	"time"
)

/*
cond_example.go

- Demonstrates sync.Cond usage to coordinate a producer and multiple consumers.
- Consumers wait until the queue has items. Producer signals consumers when items are available.
- Waiters re-check predicate in a loop to avoid spurious wakeups.
*/

func main() {
	var mu sync.Mutex
	cond := sync.NewCond(&mu)

	queue := make([]int, 0)
	const maxItems = 10

	// consumer: waits for items and processes them
	consumer := func(id int, wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			mu.Lock()
			for len(queue) == 0 {
				// Wait releases mu and suspends; reacquires mu on wake.
				cond.Wait()
			}
			// pop item
			v := queue[0]
			queue = queue[1:]
			mu.Unlock()

			// process outside lock
			fmt.Printf("consumer %d processed %d\n", id, v)
			if v == -1 {
				// -1 is sentinel to stop consumer
				return
			}
			// simulate work
			time.Sleep(100 * time.Millisecond)
		}
	}

	var wg sync.WaitGroup
	// start consumers
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go consumer(i, &wg)
	}

	// producer
	go func() {
		for i := 0; i < maxItems; i++ {
			time.Sleep(50 * time.Millisecond) // produce at some rate
			mu.Lock()
			queue = append(queue, i)
			// signal one waiter; use Broadcast() to wake all
			cond.Signal()
			mu.Unlock()
			fmt.Println("produced", i)
		}
		// send stop sentinels to consumers
		mu.Lock()
		for i := 0; i < 3; i++ {
			queue = append(queue, -1)
		}
		cond.Broadcast() // wake all consumers to consume sentinels
		mu.Unlock()
	}()

	wg.Wait()
	fmt.Println("all consumers done")
}
