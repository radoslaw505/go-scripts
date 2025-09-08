package main

import "fmt"

func main() {

	// Example usage of the generic functions
	nums := []int{1, 2, 3, 4, 5}
	evens := Filter(nums, func(n int) bool { return n%2 == 0 })
	squared := Map(nums, func(n int) int { return n * n })

	fmt.Println("Evens:", evens)
	fmt.Println("Squared:", squared)

	a, b := swap(1, 2)
	fmt.Println("Before swap:", 1, 2)
	fmt.Println("Swapped:", a, b)

	// More examples with different types

	filter := Filter([]string{"apple", "banana", "cherry"}, func(s string) bool { return len(s) > 5 })
	mapped := Map([]string{"apple", "banana", "cherry"}, func(s string) string { return s + "!" })

	fmt.Println("Filtered:", filter)
	fmt.Println("Mapped:", mapped)
}

// Simple generic functions for filtering and mapping slices
func swap[T any](a, b T) (T, T) {
	return b, a
}

// Filter returns a new slice containing only the elements that satisfy the predicate function.
func Filter[T any](slice []T, predicate func(T) bool) []T {
	var result []T
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// Map returns a new slice containing the results of applying the mapper function to each element of the input slice.
func Map[T any, R any](slice []T, mapper func(T) R) []R {
	result := make([]R, len(slice))
	for i, v := range slice {
		result[i] = mapper(v)
	}
	return result
}
