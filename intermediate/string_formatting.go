package main

import (
	"fmt"
)

type User struct {
	Name string
	Age  int
}

func main() {
	// 🖨️ Basic printing
	fmt.Print("Hello, ")
	fmt.Println("world!")

	// 🧪 Sprintf for dynamic strings
	greeting := fmt.Sprintf("Hi %s, welcome to Go!", "Radosław")
	fmt.Println(greeting)

	// 🔤 Formatting verbs
	fmt.Printf("String: %s\n", "golang")
	fmt.Printf("Integer: %d\n", 42)
	fmt.Printf("Float: %f\n", 3.14159)
	fmt.Printf("Boolean: %t\n", true)
	fmt.Printf("Type: %T\n", 3.14)
	fmt.Printf("Default format: %v\n", true)
	fmt.Printf("Go syntax format: %#v\n", []int{1, 2, 3})

	// 📏 Width and precision
	fmt.Printf("Right-aligned int: %6d\n", 7)
	fmt.Printf("Float with 2 decimals: %.2f\n", 3.14159)

	// 🧩 Struct formatting
	user := User{Name: "Radosław", Age: 30}
	fmt.Printf("Struct default: %v\n", user)
	fmt.Printf("Struct with field names: %+v\n", user)
	fmt.Printf("Struct as Go syntax: %#v\n", user)

	// 🧼 Padding and alignment
	fmt.Printf("|%-10s|%10s|\n", "left", "right")

	// 🧾 Literal percent sign
	fmt.Printf("Progress: 80%% complete\n")

	// 🧪 Custom formatted message
	message := formatUserMessage("Radosław", 30)
	fmt.Println(message)

}

func formatUserMessage(name string, age int) string {
	return fmt.Sprintf("👤 %s is %d years old and learning Go like a boss!", name, age)
}
