package main

import "fmt"

// -------------EXAMPLE-1----------------------

// Define an interface
type Logger interface {
	Log(msg string)
}

// Implement the interface
type ConsoleLogger struct{}

// Implement the Log method
func (c ConsoleLogger) Log(msg string) {
	fmt.Println("LOG:", msg)
}

// Function that takes the interface as a parameter
func run(l Logger) {
	l.Log("Starting app...")
}

// -------------EXAMPLE-2----------------------

// Define another interface
type Shape interface {
	Area() float64
}

// Implement the interface with a struct
type Rectangle struct {
	Width, Height float64
}

// Implement the Area method
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Function that takes the Shape interface as a parameter
func printArea(s Shape) {
	fmt.Println("Area:", s.Area())
}

// -------------------------------------------------

func main() {

	// -------------EXAMPLE-1----------------------

	// Create an instance of ConsoleLogger
	logger := ConsoleLogger{}

	// Pass the instance to the run function
	run(logger)

	// -------------EXAMPLE-2----------------------

	// Create an instance of Rectangle
	rect := Rectangle{Width: 10, Height: 5}

	// Pass the instance to the printArea function
	printArea(rect)

	// -------------------------------------------------
}
