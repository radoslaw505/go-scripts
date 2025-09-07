package main

import "fmt"

func main() {
	// =======================
	// Basic Array Declaration
	// =======================
	var numbers [5]int
	fmt.Println("Initial array:", numbers)

	// =======================
	// Assigning Values
	// =======================
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	numbers[4] = 50

	fmt.Println("Updated array:", numbers)

	// =======================
	// Array Properties & Access
	// =======================
	fmt.Println("Length of array:", len(numbers))
	fmt.Println("First element:", numbers[0])
	fmt.Println("Second element:", numbers[1])
	fmt.Println("Third element:", numbers[2])
	fmt.Println("Fourth element:", numbers[3])
	fmt.Println("Fifth element:", numbers[4])

	// =======================
	// Iterating with for loop
	// =======================
	fmt.Println("Iterating over array:")
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("Element at index %d: %d\n", i, numbers[i])
	}

	// =======================
	// Iterating with for-range
	// =======================
	fmt.Println("Iterating over array with for-range:")
	for index, value := range numbers {
		fmt.Printf("Element at index %d: %d\n", index, value)
	}

	// =======================
	// Array Initialization with Values
	// =======================
	arr := [3]string{"Go", "is", "fun"}
	fmt.Println("Initialized array:", arr)

	// =======================
	// Multi-dimensional Arrays
	// =======================
	var matrix [2][3]int = [2][3]int{
		{1, 2, 3},
		{0, 0, 0},
	}
	matrix[1][0] = 4
	matrix[1][1] = 5
	matrix[1][2] = 6
	fmt.Println("Multi-dimensional array (matrix):", matrix)

	// =======================
	// Iterating Multi-dimensional Arrays
	// =======================
	fmt.Println("Iterating over multi-dimensional array:")
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("Element at [%d][%d]: %d\n", i, j, matrix[i][j])
		}
	}

	fmt.Println("Length of multi-dimensional array:", len(matrix), "x", len(matrix[0]))

	// Comparison of two arrays
	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{4, 5, 6}
	fmt.Println("arr1 == arr2:", arr1 == arr2)

	// =======================
	// Array with Inferred Size
	// =======================
	inferredArray := [...]int{1, 2, 3, 4, 5}
	fmt.Println("Inferred size array:", inferredArray)

	// =======================
	// Array Copying
	// =======================
	copiedArray := numbers
	fmt.Println("Copied array:", copiedArray)

	// Modify the copied array and show that the original remains unchanged
	copiedArray[0] = 999
	fmt.Println("Modified copied array:", copiedArray)
	fmt.Println("Original array after modifying copied array:", numbers)

	// =======================
	// Array of Structs
	// =======================
	type Point struct {
		X int
		Y int
	}
	var points [2]Point
	points[0] = Point{X: 1, Y: 2}
	points[1] = Point{X: 3, Y: 4}
	fmt.Println("Array of structs:", points)

	// Iterate over the array of structs
	for i, p := range points {
		fmt.Printf("Point at index %d: (%d, %d)\n", i, p.X, p.Y)
	}

	// =======================
	// Zero Value Array
	// =======================
	var zeroArray [3]int
	fmt.Println("Zero value array:", zeroArray)
	fmt.Println("Length of zero value array:", len(zeroArray))

	// =======================
	// Blank Identifier Example
	// =======================
	_ = zeroArray // Using blank identifier to avoid unused variable error

	// Array with blank identifier in for-range loop
	fmt.Println("Iterating with blank identifier:")
	for _, value := range numbers {
		fmt.Println("Value:", value)
	}

	// Example with an error (uncomment to see the error)
	// for index, value := range numbers {	// declared but not used
	// 	fmt.Println("Value:", value)
	// }

	// =======================
	// End of Examples
	// =======================
	fmt.Println("Array examples completed.")
}
