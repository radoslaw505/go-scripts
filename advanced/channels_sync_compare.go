package main

import (
	"fmt"
	"time"
)

func unsynchronizedWorker(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(1 * time.Second)
	fmt.Printf("Worker %d finished\n", id)
}

func synchronizedWorker(id int, ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- fmt.Sprintf("Worker %d done", id)
}

func main() {
	fmt.Println("🔓 Unsynchronized:")
	go unsynchronizedWorker(1)
	go unsynchronizedWorker(2)
	time.Sleep(2 * time.Second) // Wait manually

	fmt.Println("\n🔒 Synchronized:")
	ch := make(chan string)
	go synchronizedWorker(3, ch)
	go synchronizedWorker(4, ch)

	// Receive messages from both workers
	fmt.Println("✅", <-ch)
	fmt.Println("✅", <-ch)
}
