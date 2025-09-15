package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Create a channel to receive OS signals
	sigChan := make(chan os.Signal, 1)

	// Notify on SIGINT (Ctrl+C) and SIGTERM
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Create a channel to signal graceful shutdown
	done := make(chan bool)

	// Start a goroutine to handle signals
	go func() {
		sig := <-sigChan
		fmt.Println("\nReceived signal:", sig)
		fmt.Println("Cleaning up before exit...")
		time.Sleep(1 * time.Second) // Simulate cleanup
		done <- true
	}()

	fmt.Println("App is running. Press Ctrl+C to exit.")
	<-done // Block until signal is received
	fmt.Println("App exited gracefully.")
}
