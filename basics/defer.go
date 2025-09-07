package main

func main() {

	// Defer statement syntax:
	// defer <function_call>

	// Example function to demonstrate defer
	deferExample()

	// Multiple defer statements
	multipleDefers()

	// Defer with named return values
	incrementWithDefer(10)
}

// Function to demonstrate the use of defer
func deferExample() {
	// This message will be printed first
	println("Start of function")

	// Defer the execution of this function until the surrounding function returns
	defer println("This is deferred and will run last")

	// This message will be printed second
	println("End of function")
}

// Function to demonstrate multiple defer statements
func multipleDefers() {
	defer println("First deferred call")
	defer println("Second deferred call")
	defer println("Third deferred call")

	println("Function body executing")
}

// Funtion to demonstrate incrementing a value with defer
func incrementWithDefer(i int) {
	defer println("Value in defer:", i)
	i++
	println("Value before defer:", i)
}
