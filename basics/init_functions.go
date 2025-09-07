package main

import "fmt"

// Example on init function
func init() {
	fmt.Println("Initializing...")
}

// Second init function
func init() {
	fmt.Println("Setting up...")
}

// Main function

func main() {
	fmt.Println("Main function")
}
