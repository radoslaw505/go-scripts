package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Define subcommands
	greetCmd := flag.NewFlagSet("greet", flag.ExitOnError)
	calcCmd := flag.NewFlagSet("calc", flag.ExitOnError)

	// Flags for greet
	greetName := greetCmd.String("name", "Radosław", "Name to greet")
	greetLoud := greetCmd.Bool("loud", false, "Shout the greeting")

	// Flags for calc
	calcOp := calcCmd.String("op", "add", "Operation: add or sub")
	calcA := calcCmd.Int("a", 0, "First number")
	calcB := calcCmd.Int("b", 0, "Second number")

	// Ensure a subcommand is provided
	if len(os.Args) < 2 {
		fmt.Println("Expected 'greet' or 'calc' subcommands")
		os.Exit(1)
	}

	// Switch based on subcommand
	switch os.Args[1] {
	case "greet":
		greetCmd.Parse(os.Args[2:])
		message := fmt.Sprintf("Hello, %s!", *greetName)
		if *greetLoud {
			message = strings.ToUpper(message)
		}
		fmt.Println(message)

	case "calc":
		calcCmd.Parse(os.Args[2:])
		var result int
		switch *calcOp {
		case "add":
			result = *calcA + *calcB
		case "sub":
			result = *calcA - *calcB
		default:
			fmt.Println("Unknown operation:", *calcOp)
			os.Exit(1)
		}
		fmt.Printf("Result: %d\n", result)

	default:
		fmt.Println("Unknown subcommand:", os.Args[1])
		os.Exit(1)
	}
}
