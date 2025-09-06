package main

import "fmt"

func main() {
	// Demonstrating basic arithmetic operations
	a := 10
	b := 3

	fmt.Println("Addition:", a+b)       // Addition
	fmt.Println("Subtraction:", a-b)    // Subtraction
	fmt.Println("Multiplication:", a*b) // Multiplication
	fmt.Println("Division:", a/b)       // Division
	fmt.Println("Modulus:", a%b)        // Modulus
	fmt.Println("Increment:", a+1)      // Increment
	fmt.Println("Decrement:", b-1)      // Decrement

	fmt.Println("-----")

	// Demonstrating comparison operators
	fmt.Println("Equal:", a == b)            // Equal
	fmt.Println("Not Equal:", a != b)        // Not Equal
	fmt.Println("Greater Than:", a > b)      // Greater Than
	fmt.Println("Less Than:", a < b)         // Less Than
	fmt.Println("Greater or Equal:", a >= b) // Greater or Equal
	fmt.Println("Less or Equal:", a <= b)    // Less or Equal

	// Demonstrating logical operators
	x := true
	y := false

	fmt.Println("Logical AND:", x && y) // Logical AND
	fmt.Println("Logical OR:", x || y)  // Logical OR
	fmt.Println("Logical NOT:", !x)     // Logical NOT

	// Demonstrating bitwise operators
	c := 5 // 0101 in binary
	d := 3 // 0011 in binary

	fmt.Println("Bitwise AND:", c&d)      // Bitwise AND
	fmt.Println("Bitwise OR:", c|d)       // Bitwise OR
	fmt.Println("Bitwise XOR:", c^d)      // Bitwise XOR
	fmt.Println("Bitwise AND NOT:", c&^d) // Bitwise AND NOT
	fmt.Println("Left Shift:", c<<1)      // Left Shift
	fmt.Println("Right Shift:", c>>1)     // Right Shift

	// Demonstrating assignment operators
	e := 10
	e += 5
	fmt.Println("After += :", e) // After +=

	e -= 3
	fmt.Println("After -= :", e) // After -=

	e *= 2
	fmt.Println("After *= :", e) // After *=

	e /= 4
	fmt.Println("After /= :", e) // After /=

	e %= 3
	fmt.Println("After %= :", e) // After %=

	e <<= 1
	fmt.Println("After <<= :", e) // After <<=

	e >>= 1
	fmt.Println("After >>= :", e) // After >>=

	e &= 2
	fmt.Println("After &= :", e) // After &=

	e |= 1
	fmt.Println("After |= :", e) // After |=

	e ^= 3
	fmt.Println("After ^= :", e) // After ^=

	e &^= 1
	fmt.Println("After &^= :", e) // After &^=

	fmt.Println("-----")

	// Demonstrating string concatenation
	str1 := "Hello, "
	str2 := "World!"
	str3 := str1 + str2
	fmt.Println("String Concatenation:", str3) // String Concatenation

	// Demonstrating string comparison
	fmt.Println("String Equal:", str1 == str2)            // String Equal
	fmt.Println("String Not Equal:", str1 != str2)        // String Not Equal
	fmt.Println("String Greater Than:", str1 > str2)      // String Greater Than
	fmt.Println("String Less Than:", str1 < str2)         // String Less Than
	fmt.Println("String Greater or Equal:", str1 >= str2) // String Greater or Equal
	fmt.Println("String Less or Equal:", str1 <= str2)    // String Less or Equal

	fmt.Println("-----")

	// Demonstrating pointer operations
	var num int = 42
	var ptr *int = &num // Pointer to num

	fmt.Println("Value of num:", num)                  // Value of num
	fmt.Println("Address of num:", &num)               // Address of num
	fmt.Println("Value of ptr (Address of num):", ptr) // Value of ptr (Address of num)
	fmt.Println("Value at ptr:", *ptr)                 // Value at ptr

	*ptr = 100                                                   // Changing value at the address
	fmt.Println("New value of num after changing via ptr:", num) // New value of num after changing via ptr

	fmt.Println("-----")

	// Demonstrating the comma-ok idiom with maps
	sampleMap := map[string]int{"one": 1, "two": 2, "three": 3}

	value, ok := sampleMap["two"]
	if ok {
		fmt.Println("Key 'two' found with value:", value) // Key 'two' found with value
	} else {
		fmt.Println("Key 'two' not found")
	}

	value, ok = sampleMap["four"]
	if ok {
		fmt.Println("Key 'four' found with value:", value)
	} else {
		fmt.Println("Key 'four' not found") // Key 'four' not found
	}

	fmt.Println("-----")

	// Demonstrating the use of the blank identifier
	_, exists := sampleMap["three"]
	if exists {
		fmt.Println("Key 'three' exists in the map") // Key 'three' exists in the map
	}

	// Ignoring the value and just checking existence
	if _, exists := sampleMap["five"]; !exists {
		fmt.Println("Key 'five' does not exist in the map") // Key 'five' does not exist in the map
	}

	fmt.Println("End of operator demonstrations.")

}
