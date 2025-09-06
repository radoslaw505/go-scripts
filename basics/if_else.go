package main

import "fmt"

func main() {
	// Example variables
	a := 10
	b := 20

	// If-Else statement to compare two numbers
	if a > b {
		fmt.Println("a is greater than b")
	} else if a < b {
		fmt.Println("a is less than b")
	} else {
		fmt.Println("a is equal to b")
	}

	// Nested If statement
	if a != 0 {
		if b%a == 0 {
			fmt.Println("b is divisible by a")
		} else {
			fmt.Println("b is not divisible by a")
		}
	}

	// Using short variable declaration in if statement
	if c := a + b; c > 25 {
		fmt.Println("The sum of a and b is greater than 25:", c)
	} else {
		fmt.Println("The sum of a and b is 25 or less:", c)
	}

	fmt.Println("End of If-Else examples")
}
