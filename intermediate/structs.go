package main

import (
	"fmt"
	"strconv"
)

func main() {

	// Defining a simple struct
	type Person struct {
		Name string
		Age  int
	}

	// Creating an instance of the struct
	p := Person{Name: "Radosław", Age: 30}

	// Accessing struct fields
	println("Name:", p.Name) // Output: Name: Radosław
	println("Age:", p.Age)   // Output: Age: 30

	// Modifying struct fields
	p.Age = 31
	println("Updated Age:", p.Age) // Output: Updated Age: 31

	// Structs with Pointers
	type Point struct {
		X, Y int
	}

	// Creating a pointer to a struct
	pt := &Point{X: 10, Y: 20}

	// Accessing fields via pointer
	println("Point X:", pt.X) // Output: Point X: 10
	println("Point Y:", pt.Y) // Output: Point Y: 20

	// Modifying fields via pointer
	pt.X = 15
	println("Updated Point X:", pt.X) // Output: Updated Point X: 15

	// Anonymous Structs
	user := struct {
		Name string
		Age  int
	}{
		Name: "Radosław",
		Age:  30,
	}

	println("User Name:", user.Name) // Output: User Name: Radosław
	println("User Age:", user.Age)   // Output: User Age: 30

	// Nested Structs
	type Address struct {
		City, State string
	}

	type Employee struct {
		Name    string
		Age     int
		Address Address
	}

	e := Employee{
		Name: "Radosław",
		Age:  30,
		Address: Address{
			City:  "Warsaw",
			State: "Mazovia",
		},
	}

	println("Employee Name:", e.Name)           // Output: Employee Name: Radosław
	println("Employee City:", e.Address.City)   // Output: Employee City: Warsaw
	println("Employee State:", e.Address.State) // Output: Employee State: Mazovia

	// Structs in Functions
	pp := Person{Name: "Radosław", Age: 31}

	person := func(p Person) string {
		return p.Name + " is " + strconv.Itoa(p.Age) + " years old."
	}(pp)

	fmt.Println(person) // Output: Radosław is 31 years old.

	// Embedded Structs
	type Animal struct {
		Name string
	}

	type Dog struct {
		Animal // Embedding Animal struct
		Breed  string
	}

	d := Dog{
		Animal: Animal{Name: "Buddy"},
		Breed:  "Golden Retriever",
	}

	println("Dog Name:", d.Name)   // Output: Dog Name: Buddy
	println("Dog Breed:", d.Breed) // Output: Dog Breed: Golden Retriever
}
