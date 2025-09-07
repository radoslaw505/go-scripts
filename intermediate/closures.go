package main

import "fmt"

func main() {

	// Get the closure
	add := adder(0)

	// Use the closure
	// Each call to add updates the sum variable in the closure's environment
	fmt.Println("Adding 3:", add(3))   // 3
	fmt.Println("Adding 7:", add(7))   // 10
	fmt.Println("Adding 10:", add(10)) // 20
	fmt.Println("Adding 5:", add(5))   // 25
}

// adder returns a closure that sums up integers.
func adder(i int) func(int) int {
	// This is called when adder is called
	sum := i
	fmt.Println("Initializing sum to :", sum)

	// This is the closure
	// This is called when the returned function is called
	return func(x int) int {
		sum += x
		return sum
	}
}
