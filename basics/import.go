// Declares the package name. 'main' is used for executable programs.
package main

// Imports required packages:
// - fmt: for formatted I/O
// - net/http: for HTTP client functionality
import (
	foo "fmt"
	bar "net/http"
)

// The main function is the entry point of the program.
func main() {
	// Print a message to the console.
	foo.Println("Hello, Go Standard Library!")

	// Send an HTTP GET request to a public API endpoint.
	resp, err := bar.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		// Print an error message if the request fails.
		foo.Println("Error fetching data:", err)
		return
	}
	// Ensure the response body is closed after function exits.
	defer resp.Body.Close()

	// Check if the HTTP response status code is 200 (OK).
	if resp.StatusCode != bar.StatusOK {
		foo.Println("Error: Received non-200 response code! HTTP Response Status:", resp.Status)
		return
	}

	// Print a success message with the HTTP response status.
	foo.Println("Successfully fetched data! HTTP Response Status:", resp.Status)
}
