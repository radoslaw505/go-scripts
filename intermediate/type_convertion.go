package main

import (
	"fmt"
	"strconv"
)

func main() {
	// int to float
	i := 42
	f := float64(i)
	fmt.Println("Float:", f)

	// float to int
	f2 := 99.99
	i2 := int(f2)
	fmt.Println("Int:", i2)

	// string to int
	s := "123"
	num, _ := strconv.Atoi(s)
	fmt.Println("Parsed int:", num)

	// int to string
	str := strconv.Itoa(num)
	fmt.Println("String:", str)

	// string to bytes
	bytes := []byte("GoLang")
	fmt.Println("Bytes:", bytes)

	// bytes to string
	back := string(bytes)
	fmt.Println("Back to string:", back)
}
