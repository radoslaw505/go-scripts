package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	// Sorting integers
	nums := []int{5, 2, 6, 3, 1}
	sort.Ints(nums)
	fmt.Println("Sorted ints:", nums)

	// Sorting strings
	fruits := []string{"banana", "apple", "cherry"}
	sort.Strings(fruits)
	fmt.Println("Sorted strings:", fruits)

	// Sorting custom structs by age
	people := []Person{
		{"John", 28},
		{"Alice", 23},
		{"Bob", 21},
	}
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println("Sorted people by age:", people)
}
