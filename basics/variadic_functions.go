package main

import "fmt"

func main() {

	// Variadic function syntax:
	// func <function_name>(<parameters list>, <variadic_parameter> ...<type>) <return_type> { <function_body> }

	// Calling a variadic function with multiple arguments
	fmt.Println("Sum:", sum(1, 2, 3, 4, 5))

	// Calling a variadic function with a slice
	numbers := []int{6, 7, 8, 9, 10}
	fmt.Println("Sum with slice:", sum(numbers...))

	// Calling a variadic function with no arguments
	fmt.Println("Sum with no arguments:", sum())

	// Calling a variadic function that returns multiple values
	str, total := newSum("Total Sum:", 1, 2, 3)
	fmt.Println(str, total)

}

// Variadic function that sums an arbitrary number of integers
func sum(numbers ...int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}

func newSum(returnString string, numbers ...int) (string, int) {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return returnString, total
}
