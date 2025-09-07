package main

import "fmt"

func main() {

	// Panic syntax:
	// panic(<value>)

	// Example function to demonstrate panic and recover
	panicExample()

	// Example of handling panic in a function
	result, err := safeDivide(10, 2)
	fmt.Printf("Result: %d, Error: %v\n", result, err)

	result1, err1 := safeDivide(10, 0)
	fmt.Printf("Result: %d, Error: %v\n", result1, err1)

}

// Function to demonstrate the use of panic and recover
func panicExample() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println("Starting function")

	// Trigger a panic
	panic("Something went wrong!")

	// This line will not be executed
	// fmt.Println("This line will not be executed")
}

// Example function how to handle panic with defer and recover
func safeDivide(a int, b int) (int, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic in safeDivide:", r)
		}
	}()

	if b == 0 {
		panic("division by zero")
	}
	return a / b, nil
}
