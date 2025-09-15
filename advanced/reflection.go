package main

import (
	"fmt"
	"reflect"
)

// Define a sample struct
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// Create an instance of Person
	p := Person{Name: "Radosław", Age: 30}

	// Get the reflect.Type and reflect.Value
	t := reflect.TypeOf(p)
	v := reflect.ValueOf(p)

	fmt.Println("Type:", t.Name())
	fmt.Println("Kind:", t.Kind())

	// Iterate over struct fields
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		fmt.Printf("Field %d: %s (%s) = %v\n", i, field.Name, field.Type, value)
		fmt.Printf("Tag: json=\"%s\"\n", field.Tag.Get("json"))
	}

	// Modify a value using reflection
	fmt.Println("\nModifying Age using reflection...")
	vp := reflect.ValueOf(&p).Elem() // Get addressable value
	ageField := vp.FieldByName("Age")
	if ageField.CanSet() {
		ageField.SetInt(35)
	}

	fmt.Println("Updated struct:", p)
}
