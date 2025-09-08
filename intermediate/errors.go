package main

import (
	"errors"
	"fmt"
)

// CustomError defines a structured error with a code and message
type CustomError struct {
	Code    int
	Message string
}

// Error implements the error interface for CustomError
func (e CustomError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

// divide performs division and returns an error if dividing by zero
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// fetchData simulates a function that returns a custom error
func fetchData(id int) (string, error) {
	if id <= 0 {
		return "", CustomError{Code: 400, Message: "Invalid ID"}
	}
	// Simulate success
	return fmt.Sprintf("Data for ID %d", id), nil
}

// ValidationError method demonstrates another custom error type
type ValidationError struct {
	Code    int
	Message string
}

// Error implements the error interface for ValidationError
func (e *ValidationError) Error() string {
	return fmt.Sprintf("Validation Error %d: %s", e.Code, e.Message)
}

// Error functions demonstrate error handling in Go
func validation() error {
	return &ValidationError{Code: 422, Message: "Unprocessable Entity"}
}

func main() {
	// Example 1: Standard error
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Standard error:", err)
	} else {
		fmt.Println("Division result:", result)
	}

	// Example 2: Custom error
	data, err := fetchData(-1)
	if err != nil {
		// Use type assertion to extract custom error details
		if ce, ok := err.(CustomError); ok {
			fmt.Printf("Custom error caught: Code=%d, Message=%s\n", ce.Code, ce.Message)
		} else {
			fmt.Println("Unknown error:", err)
		}
	} else {
		fmt.Println("Fetched data:", data)
	}

	// Example 3: Validation error
	if err := validation(); err != nil {
		if ve, ok := err.(*ValidationError); ok {
			fmt.Printf("Validation error caught: Code=%d, Message=%s\n", ve.Code, ve.Message)
		} else {
			fmt.Println("Unknown error:", err)
		}
	}

}
