package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	// Create a new http client
	client := &http.Client{}

	// GET Request
	resp, err := client.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("Error making GET request: ", err)
		return
	}
	defer resp.Body.Close()

	// Read and print response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading body response: ", err)
		return
	}
	defer resp.Body.Close()

	// Print body
	fmt.Println(body)
	fmt.Println(string(body))

	// Print request status
	fmt.Println("Request status: ", resp.Status)
	fmt.Println("Request status code: ", resp.StatusCode)
}
