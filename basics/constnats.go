package main

import "fmt"

const pi = 3.14
const GRAVITY = 9.8

func main() {

	// Constants
	const MAX_RETRIES int = 5 // ALL_CAPS with underscores for constants

	// Override the constant value
	// pi = 3.1 // This will cause a compile-time error
	const pi = 3.14159 // Shadowing the outer pi constant

	// Print the constant to avoid unused variable errors
	fmt.Println(MAX_RETRIES)
	fmt.Println("Value of pi:", pi)
	fmt.Println("Value of gravity:", GRAVITY)

}
