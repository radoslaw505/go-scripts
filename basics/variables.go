package main

import "fmt"

// Global variable, accessible throughout the package.
// := cannot be used for global variables.
var globalVar = "I am global"

// Variable to override in main function
var name string = "Alice"

func main() {
	// Declare variables of different types.
	var name string = "John"
	var age int = 30
	var height float64 = 5.9
	var isEmployed bool = true

	count := 10 // Short variable declaration

	// Default values
	// Declare a variable without initialization
	var defaultInt int
	var defaultString string
	var defaultBool bool
	// Pointers, slices, maps, channels, and interfaces default to nil

	// Print default values
	fmt.Println("Default int:", defaultInt)       // 0
	fmt.Println("Default string:", defaultString) // ""
	fmt.Println("Default bool:", defaultBool)     // false

	// Print the value of count
	fmt.Println("Count:", count)

	var lastName = "Smith" // Type inferred as string

	// Print the types of the variables.
	fmt.Printf("Type of lastName: %T\n", lastName)

	// Print the values of the variables.
	fmt.Println("Name:", name, lastName)
	fmt.Println("Age:", age)
	fmt.Println("Height:", height)
	fmt.Println("Employed:", isEmployed)

	// -----------------------------
	// Scope of Variables
	// -----------------------------

	// Accessing local variable
	scopeExample()
	// fmt.Println(localVar) // This would cause an error: undefined localVar

	// Accessing global variable
	fmt.Println(globalVar)
}

func scopeExample() {
	// Local variable, only accessible within this function.
	var localVar = "I am local"
	fmt.Println(localVar)
}
