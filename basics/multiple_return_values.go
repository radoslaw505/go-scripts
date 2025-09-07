package main

import "fmt"

func main() {

	// Function returning multiple values syntax:
	// func <function_name>(<parameters list>) (<return_type1>, <return_type2>, ...) { <function_body> }

	// Example function that divides two numbers and returns the quotient and remainder
	// Calling the function and handling multiple return values
	quotient, remainder := divide(11, 2)
	if remainder != 0 {
		fmt.Println("Remainder:", remainder)
	}
	fmt.Println("Quotient:", quotient)

	// Calling the function that handles errors
	result, err := safeDivide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

}

// Function that divides two numbers and returns the quotient and remainder
func divide(a int, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

// Error handling example with multiple return values
func safeDivide(a int, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero is not allowed")
	}
	return a / b, nil
}
