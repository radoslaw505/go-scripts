package main

import (
	"fmt"
	"time"
)

/*
deadlock_fixed_channel.go

- Fixes unbuffered and cyclic deadlocks by starting receivers, using buffering, or seeding channels.
*/

func fixUnbuffered() {
	ch := make(chan int)
	go func() {
		val := <-ch
		fmt.Println("received", val)
	}()
	ch <- 1 // now there's a receiver, send succeeds
}

func fixWithBuffered() {
	ch := make(chan int, 1) // buffer size 1 allows one send without receiver
	ch <- 42                // does not block
	fmt.Println("sent on buffered channel without immediate receiver")
	fmt.Println("receiving:", <-ch)
}

func fixCyclic() {
	a := make(chan struct{}, 1)
	b := make(chan struct{}, 1)

	a <- struct{}{} // seed 'a' so the first goroutine can proceed

	go func() {
		<-a
		fmt.Println("goroutine1 got a, sending to b")
		b <- struct{}{}
	}()

	go func() {
		<-b
		fmt.Println("goroutine2 got b, sending to a")
		a <- struct{}{}
	}()

	time.Sleep(200 * time.Millisecond)
}

func main() {
	fixUnbuffered()
	fixWithBuffered()
	fixCyclic()
	fmt.Println("fixes applied, exiting")
}
