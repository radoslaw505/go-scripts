package main

import "fmt"

func main() {
	// Using a for loop to print numbers from 1 to 5
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// Print a separator
	fmt.Println("-----")

	// Using a for loop to iterate over a slice
	numbers := []int{10, 20, 30, 40, 50}
	for index, value := range numbers {
		fmt.Println("Index:", index, "Value:", value)
	}

	// Print a separator
	fmt.Println("-----")

	// Using a for loop to iterate over a map
	person := map[string]string{"name": "Alice", "city": "Wonderland"}
	for key, value := range person {
		fmt.Println("Key:", key, "Value:", value)
	}

	// Print a separator
	fmt.Println("-----")

	// Using a for loop as a while loop
	count := 0
	for count < 3 {
		fmt.Println("Count is:", count)
		count++
	}

	// Print a separator
	fmt.Println("-----")

	// Infinite loop with a break condition
	sum := 0
	for {
		sum += 1
		if sum >= 5 {
			break
		}
	}
	fmt.Println("Final Sum:", sum)

	// Print a separator
	fmt.Println("-----")

	// Using continue to skip even numbers
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Odd Number:", i)
	}

	// Print a separator
	fmt.Println("-----")

	// Nested loops to print a multiplication table
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("%d * %d = %d\n", i, j, i*j)
		}
	}

	// Print a separator
	fmt.Println("-----")

	// While loop using for
	n := 5
	for n > 0 {
		fmt.Println("Countdown:", n)
		n--
	}

	// Print a separator
	fmt.Println("-----")

	// Using a for loop to iterate over a string
	str := "GoLang"
	for index, char := range str {
		fmt.Printf("Index: %d, Character: %c\n", index, char)
	}

	fmt.Println("-----")

	// Using a outer loop with a labeled break
	// What is outer loop? A labeled break statement allows you to break out of an outer loop from within a nested loop.
OuterLoop:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if i*j > 4 {
				break OuterLoop
			}
			fmt.Printf("%d * %d = %d\n", i, j, i*j)
		}
	}

	fmt.Println("Exited the outer loop")

	fmt.Println("-----")

	// While loop with break
	counter := 0
	for {
		if counter >= 7 {
			break
		}
		fmt.Println("Counter:", counter)
		counter++
	}
}
