package main

import "fmt"

func main() {

	// Example usage of a recursive function
	fmt.Println("Factorial of 5 is:", factorial(5))

	// Example usage of Fibonacci function
	fmt.Println("Fibonacci of 7 is:", fibonacci(7))

	// Example usage of Fibonacci with memoization
	fmt.Println("Fibonacci with memoization of 420 is:", fibonacciMemo(420))

}

// Example of a simple recursive function
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// Another example: Fibonacci sequence using recursion
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// Example of a simple recursive function with memoization
var fibMemo = map[int]int{}

func fibonacciMemo(n int) int {
	if n <= 1 {
		return n
	}
	if val, exists := fibMemo[n]; exists {
		return val
	}
	fibMemo[n] = fibonacciMemo(n-1) + fibonacciMemo(n-2)
	return fibMemo[n]
}
