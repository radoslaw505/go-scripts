package main

import "fmt"

func main() {
	// Example variables
	a := 10
	b := 20

	// Switch statement to compare two numbers
	switch {
	case a > b:
		fmt.Println("a is greater than b")
	case a < b:
		fmt.Println("a is less than b")
	default:
		fmt.Println("a is equal to b")
	}

	// Switch with multiple cases
	day := 3
	switch day {
	case 1, 2, 3:
		fmt.Println("It's the beginning of the week.")
	case 4, 5:
		fmt.Println("It's midweek.")
	case 6, 7:
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("Invalid day.")
	}

	// Switch with fallthrough
	// Note: fallthrough is used to continue to the next case
	num := 2
	switch num {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
		fallthrough
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Other number")
	}

	fmt.Println("End of Switch examples")

	// Type switch example
	checkType(42)
	checkType(3.14)
	checkType("Hello")
	checkType(true)
}

func checkType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("It's an integer")
	case int32:
		fmt.Println("It's an integer")
	case float64:
		fmt.Println("It's float")
	case string:
		fmt.Println("It's a string")
	default:
		fmt.Println("Unknown Type")
	}
}
