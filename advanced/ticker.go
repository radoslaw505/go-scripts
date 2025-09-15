package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Starting ticker...")

	// Create a ticker that ticks every second
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop() // Ensure ticker is stopped when done

	// i := 1
	// for range ticker.C {
	// 	i *= 2
	// 	fmt.Println(i)
	// }

	done := make(chan bool)

	// Stop the ticker after 5 ticks
	go func() {
		time.Sleep(5 * time.Second)
		done <- true
	}()

	count := 0
	for {
		select {
		case t := <-ticker.C:
			count++
			fmt.Println("Tick at", t)
		case <-done:
			fmt.Println("Ticker stopped after", count, "ticks")
			return
		}
	}
}
