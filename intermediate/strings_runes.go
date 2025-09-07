package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	// Example of a string containing both ASCII and non-ASCII characters
	str := "Hello, 世界"
	fmt.Println(str) // Output: Hello, 世界

	// Common string operations
	s := "golang"

	// Length in bytes
	fmt.Println(len(s)) // 6

	// Accessing individual bytes
	fmt.Println(s[0]) // 103 ('g')

	// String concatenation
	s2 := s + " rocks"
	fmt.Println(s2) // golang rocks

	// Example of rune literals
	var r rune = '世'
	fmt.Println(r)        // 19990
	fmt.Printf("%c\n", r) // 世

	// Iterating over runes
	newStr := "Hello, 世界"

	for i, r := range newStr {
		fmt.Printf("Index: %d, Rune: %c\n", i, r)
	}

	// Converting Between Strings and Runes
	str1 := "Go语言"
	runes1 := []rune(str1)
	fmt.Println(runes1) // [71 111 35821 35328]

	// Converting back to string
	runes2 := []rune{71, 111, 35821, 35328}
	str2 := string(runes2)
	fmt.Println(str2) // Go语言

	// Unicode and Encoding
	str3 := "世界"
	fmt.Println(str3[0]) // 228 (not '世', but the first byte of UTF-8 encoding)

	// Counting runes
	fmt.Println(len([]rune(str3))) // 2

	// Counting runes using utf8 package
	str4 := "Hello, 世界"
	fmt.Println("Byte length:", len(str4))
	fmt.Println("Rune count:", utf8.RuneCountInString(str4))
}
