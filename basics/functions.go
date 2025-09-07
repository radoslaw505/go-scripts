package main

import "fmt"

func main() {

	// func <function_name>(<parameters list>) <return_type> { <function_body> }

	fmt.Println("Hello from the main function!")

	// Function with no parameters and no return value
	greet()

	fmt.Println("Sum:", add(3, 5))

	// Anonymous function assigned to a variable
	var multiply = func(a int, b int) int {
		return a * b
	}
	fmt.Println("Product:", multiply(4, 6))

	// Immediately Invoked Function Expression (IIFE)
	func() {
		fmt.Println("This is an IIFE!")
	}()

	// IIFE with parameters
	func(name string) {
		fmt.Println("Hello,", name)
	}("Go Developer")

	// IIFE with parameters and return value
	result := func(a int, b int) int {
		return a - b
	}(10, 4)
	fmt.Println("Difference:", result)

	// Using higher-order function
	sum := applyOperation(7, 3, add)
	fmt.Println("Sum using higher-order function:", sum)

	// Function returning another function
	double := getMultiplier(2)
	fmt.Println("Double of 5:", double(5))
}

// Function with parameters and return value
func greet() {
	fmt.Println("Greetings from the greet function!")
}

// Function with parameters and return value
func add(a int, b int) int {
	return a + b
}

// Higher-order function: takes a function as a parameter
func applyOperation(a int, b int, operation func(int, int) int) int {
	return operation(a, b)
}

// Function that returns another function
func getMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}
