package main

import "fmt"

func main() {

	// Basic Formatting Verbs
	name := "Radosław"
	age := 30
	fmt.Printf("Name: %q, Age: %d\n", name, age)

	// Controlling Width and Precision
	pi := 3.14159
	fmt.Printf("Pi: %6.2f\n", pi) // Output: "  3.14"
	// %6.2f means a total width of 6 characters, with 2 digits after the decimal point.

	// Padding with Zeros
	number := 42
	fmt.Printf("Padded Number: %05d\n", number) // Output: "00042"

	// Formatting Booleans
	isActive := true
	fmt.Printf("Is Active: %t\n", isActive) // Output: "Is Active: true"

	// Formatting Hexadecimal and Octal
	value := 255
	fmt.Printf("Hex: %x, Octal: %o\n", value, value) // Output: "Hex: ff, Octal: 377"

	// Using Sprintf to format and return a string
	formatted := fmt.Sprintf("Hello, %s!", "world")
	fmt.Println(formatted)

	// Formatting Structs
	type Person struct {
		Name string
		Age  int
	}
	p := Person{"Alice", 28}
	fmt.Printf("Person: %+v\n", p) // Output: "Person: {Name:Alice Age:28}"
}
