package main

import (
	"fmt"
	"os"
)

func main() {
	// Set environment variable
	os.Setenv("GREETING", "Hello, Radosław!")

	// Get environment variable
	fmt.Println("GREETING:", os.Getenv("GREETING"))

	// Lookup with existence check
	value, ok := os.LookupEnv("GREETING")
	fmt.Println("GREETING exists?", ok, "| Value:", value)

	// Unset variable
	os.Unsetenv("GREETING")
	fmt.Println("GREETING after unset:", os.Getenv("GREETING"))

	// List all environment variables
	fmt.Println("\nAll environment variables:")
	for _, env := range os.Environ() {
		fmt.Println(env)
	}
}
