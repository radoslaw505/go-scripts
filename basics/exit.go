package main

import (
	"fmt"
	"os"
)

func main() {

	// Example of function with 128 exit code
	exitWithCode128()

	fmt.Println("Exiting the program with status code 1")
	// Exit the program with a non-zero status code to indicate an error.
	// Note: os.Exit does not run deferred functions.
	// Uncomment the line below to see it in action.
	os.Exit(1)

	// This line will not be executed if os.Exit is called above.
	// fmt.Println("This line will not be executed if os.Exit is called")
}

// Example of function with 128 exit code
func exitWithCode128() {
	fmt.Println("Exiting the program with status code 128")
	os.Exit(128)
}
