package main

import "fmt"

func main() {

	// Creating an instance of Person with embedded Address
	p := Person{
		Name: "Alice",
		Age:  30,
		Address: Address{
			City:  "Wonderland",
			State: "Fantasy",
		},
	}

	// Accessing fields from both Person and embedded Address
	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
	fmt.Println("City:", p.City)   // Accessing embedded struct field directly
	fmt.Println("State:", p.State) // Accessing embedded struct field directly

	// Modifying fields
	p.City = "New Wonderland"
	p.Age = 31

	fmt.Println("Updated City:", p.City)
	fmt.Println("Updated Age:", p.Age)

	// Demonstrating that embedded struct fields can be accessed directly
	fmt.Printf("Full Address: %s, %s\n", p.Address.City, p.Address.State)

	// Embedding structs within structs
	type Company struct {
		Name    string
		Address Address
	}

	c := Company{
		Name: "Tech Corp",
		Address: Address{
			City:  "Tech City",
			State: "Innovation",
		},
	}

	fmt.Println("Company Name:", c.Name)
	fmt.Println("Company City:", c.Address.City)
	fmt.Println("Company State:", c.Address.State)

	// Using methods with embedded structs
	fmt.Println(p.Greet())
	fmt.Println(c.Name, "is located in", c.Address.City)
}

// Defining embedded structs
type Address struct {
	City, State string
}

type Person struct {
	Name string
	Age  int
	Address
}

// Methods with embedded structs
func (p Person) Greet() string {
	return fmt.Sprintf("Hello, my name is %s and I live in %s, %s.", p.Name, p.City, p.State)
}
