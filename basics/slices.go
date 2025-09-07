package main

import "fmt"

func main() {
	// Declare and initialize a slice of integers.
	numbers := []int{10, 20, 30, 40, 50}

	// Print the entire slice.
	fmt.Println("Slice:", numbers)

	// Print the length of the slice.
	fmt.Println("Length of slice:", len(numbers))

	// Access and print individual elements of the slice.
	fmt.Println("First element:", numbers[0])
	fmt.Println("Second element:", numbers[1])
	fmt.Println("Third element:", numbers[2])
	fmt.Println("Fourth element:", numbers[3])
	fmt.Println("Fifth element:", numbers[4])

	// Iterate over the slice using a traditional for loop.
	fmt.Println("Iterating over slice with for loop:")
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("Element at index %d: %d\n", i, numbers[i])
	}

	// Iterate over the slice using a for-range loop.
	fmt.Println("Iterating over slice with for-range:")
	for index, value := range numbers {
		fmt.Printf("Element at index %d: %d\n", index, value)
	}

	// Demonstrate slicing operations.
	subSlice1 := numbers[1:4] // Elements from index 1 to 3
	subSlice2 := numbers[:3]  // First three elements
	subSlice3 := numbers[2:]  // Elements from index 2 to the end

	fmt.Println("Sub-slice from index 1 to 3:", subSlice1)
	fmt.Println("First three elements:", subSlice2)
	fmt.Println("Elements from index 2 to end:", subSlice3)

	// Append new elements to the slice.
	numbers = append(numbers, 60, 70)
	fmt.Println("Slice after appending elements:", numbers)

	// Create a slice with a specified length and capacity using make.
	madeSlice := make([]int, 5, 10) // Length 5, Capacity 10
	fmt.Println("Made slice:", madeSlice)
	fmt.Println("Length of made slice:", len(madeSlice))
	fmt.Println("Capacity of made slice:", cap(madeSlice))

	// Append elements to the made slice to demonstrate capacity growth.
	madeSlice = append(madeSlice, 1, 2, 3, 4, 5, 6)
	fmt.Println("Made slice after appending elements:", madeSlice)
	fmt.Println("Length of made slice after appending:", len(madeSlice))
	fmt.Println("Capacity of made slice after appending:", cap(madeSlice))

	// Two dimensional slice (slice of slices).
	twoD := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("Two-dimensional slice:", twoD)

	// Accessing elements in a two-dimensional slice.
	fmt.Println("Element at [0][0]:", twoD[0][0])
	fmt.Println("Element at [1][2]:", twoD[1][2])
	fmt.Println("Element at [2][1]:", twoD[2][1])

	// Iterating over a two-dimensional slice.
	fmt.Println("Iterating over two-dimensional slice:")
	for i := 0; i < len(twoD); i++ {
		for j := 0; j < len(twoD[i]); j++ {
			fmt.Printf("Element at [%d][%d]: %d\n", i, j, twoD[i][j])
		}
	}

	fmt.Println("Length of two-dimensional slice:", len(twoD), "x", len(twoD[0]))

	// Note: Slices cannot be compared directly using '==' operator.
	// Uncommenting the following lines will result in a compilation error.
	// slice1 := []int{1, 2, 3}
	// slice2 := []int{1, 2, 3}
	// fmt.Println("slice1 == slice2:", slice1 == slice2) // Error!

}
