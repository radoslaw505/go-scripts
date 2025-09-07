package main

import "fmt"

func main() {

	// Using different fmt functions
	// Print does not add a newline at the end
	fmt.Print("Hello")

	// Println adds a newline at the end
	fmt.Println("World")

	// Printf allows formatted strings
	fmt.Printf("Hello, %s\n", "John")

	// Formatting with Verbs
	name := "Radosław"
	age := 30
	fmt.Printf("Name: %s, Age: %d\n", name, age)

	// Input Functions
	var name1 string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name1)
	fmt.Println("Hello,", name1)

	// String Formatting
	str := fmt.Sprintf("Pi is approximately %.2f", 3.14159)
	fmt.Println(str) // Output: Pi is approximately 3.14

	// Error Handling
	connErr := fmt.Errorf("connection refused")
	err := fmt.Errorf("failed to connect: %v", connErr)
	fmt.Println(err)

	// Formatting Structs
	p := Person{"Radosław", 30}
	fmt.Printf("Person: %+v\n", p) // Output: Person: {Name:Radosław Age:30}

	// Sprintln example
	greeting := fmt.Sprintln("Hello", "from", "Go!")
	fmt.Print(greeting) // Output: Hello from Go!

	// Sprintf example
	intro := fmt.Sprintf("My name is %s and I am %d years old.", name, age)
	fmt.Println(intro) // Output: My name is Radosław and I am 30 years old.

	// Errorf example
	fileErr := fmt.Errorf("file not found: %s", "data.txt")
	fmt.Println(fileErr) // Output: file not found: data.txt

	// Using the Stringer interface
	fmt.Println(p) // Output: Radosław (30 years old)
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s (%d years old)", p.Name, p.Age)
}
