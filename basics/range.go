package main

import "fmt"

func main() {

	// Range over strings
	str := "Hello, Go!"
	fmt.Println("Range over string:")
	for index, char := range str {
		fmt.Printf("Index: %d, Character: %c\n", index, char)
	}

	// Range over arrays
	arr := [5]int{10, 20, 30, 40, 50}
	fmt.Println("\nRange over array:")
	for index, value := range arr {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// Range over slices
	slice := []string{"Go", "is", "fun"}
	fmt.Println("\nRange over slice:")
	for index, value := range slice {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}

	// Range over maps
	m := map[string]int{
		"Alice":   30,
		"Bob":     25,
		"Charlie": 35,
	}
	fmt.Println("\nRange over map:")
	for key, value := range m {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}

	// Range over channels
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)

	fmt.Println("\nRange over channel:")
	for value := range ch {
		fmt.Printf("Received value: %d\n", value)
	}

	// Using blank identifier to ignore index or value
	fmt.Println("\nUsing blank identifier:")
	for _, value := range arr {
		fmt.Printf("Value: %d\n", value)
	}
	for index := range slice {
		fmt.Printf("Index: %d\n", index)
	}

	// Modifying slice elements during range
	fmt.Println("\nModifying slice elements during range:")
	nums := []int{1, 2, 3, 4, 5}
	for index := range nums {
		nums[index] *= 2
	}
	fmt.Println("Modified slice:", nums)

	// Note: Modifying map during range is not safe and will cause a runtime panic.
	// Uncommenting the following code will demonstrate this behavior.
	/*
		fmt.Println("\nModifying map during range (unsafe):")
		for key := range m {
			if key == "Bob" {
				delete(m, key) // This will cause a panic
			}
		}
	*/

	// To safely modify a map during iteration, collect keys to delete first.
	fmt.Println("\nSafely modifying map during range:")
	keysToDelete := []string{}
	for key := range m {
		if key == "Bob" {
			keysToDelete = append(keysToDelete, key)
		}
	}
	for _, key := range keysToDelete {
		delete(m, key)
	}
	fmt.Println("Modified map:", m)

	// Note: The order of iteration over maps is not guaranteed to be the same each time.
	// This is due to the internal implementation of maps in Go.
	fmt.Println("\nIterating over map multiple times to show order is not guaranteed:")
	for i := 0; i < 3; i++ {
		fmt.Printf("Iteration %d:\n", i+1)
		for key, value := range m {
			fmt.Printf("Key: %s, Value: %d\n", key, value)
		}
		fmt.Println()
	}

	fmt.Println("End of range examples.")
}
