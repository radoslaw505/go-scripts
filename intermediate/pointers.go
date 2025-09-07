package main

import "fmt"

func main() {

	// Declare an integer variable 'a' and assign it the value 42.
	a := 42
	// Declare a pointer variable 'p' that holds the address of 'a'.
	p := &a

	// Print the value of 'a' using the pointer 'p'.
	fmt.Println("Value of a using pointer p:", *p)

	// Modify the value of 'a' through the pointer 'p'.
	*p = 100

	// Print the modified value of 'a'.
	fmt.Println("Modified value of a:", a)
	fmt.Println("Value of a using pointer p after modification:", *p)

	// Print the address of 'a' and the value of pointer 'p' to show they are the same.
	fmt.Println("Address of a:", &a)
	fmt.Println("Value of pointer p (address of a):", p)

	// nil pointer example
	var pNil *int
	if pNil == nil {
		fmt.Println("pNil is a nil pointer")
	}

	// Assign the address of 'a' to the nil pointer 'pNil'.
	pNil = &a
	fmt.Println("Value of a using pNil after assignment:", *pNil)

	// Change the value of 'a' through 'pNil'.
	*pNil = 200
	fmt.Println("Value of a after modifying through pNil:", a)

	// Function that increments an integer value using a pointer.
	num := 10
	fmt.Println("Value of num before increment:", num)
	increment(&num)
	fmt.Println("Value of num after increment:", num)
}

func increment(n *int) {
	*n++
}
