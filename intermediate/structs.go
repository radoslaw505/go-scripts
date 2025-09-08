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

	// Struct Tags
	type Product struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Price int    `json:"price"`
	}

	prod := Product{ID: 1, Name: "Laptop", Price: 1000}

	println("Product ID:", prod.ID) // Output: Product ID: 1

	// Embedded Structs
	type Contact struct {
		Email string
		Phone string
	}

	type User struct {
		Name    string
		Age     int
		Contact // Embedded struct
	}

	u := User{
		Name: "Radosław",
		Age:  30,
		Contact: Contact{
			Email: "radoslaw@example.com",
			Phone: "123-456-7890",
		},
	}

	println("User Name:", u.Name)           // Output: User Name: Radosław
	println("User Age:", u.Age)             // Output: User Age: 30
	println("User Email:", u.Contact.Email) // Output: User Email: radoslaw@example.com
	println("User Phone:", u.Contact.Phone) // Output: User Phone: 123-456-7890

	// Accessing embedded struct fields directly
	println("User Email (direct):", u.Email) // Output: User Email (direct):
}
