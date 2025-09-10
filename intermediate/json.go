package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// Define a struct for structured JSON
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email,omitempty"` // omits if empty
	Password string `json:"-"`               // excluded from JSON
}

func main() {
	// === 1. Encode struct to JSON ===
	user := User{
		ID:    1,
		Name:  "Radosław",
		Email: "radoslaw@example.com",
	}

	jsonBytes, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		log.Fatal("Error encoding JSON:", err)
	}
	fmt.Println("Encoded JSON:\n", string(jsonBytes))

	// === 2. Decode JSON to struct ===
	jsonStr := `{"id":2,"name":"Alice","email":"alice@example.com"}`
	var decodedUser User
	err = json.Unmarshal([]byte(jsonStr), &decodedUser)
	if err != nil {
		log.Fatal("Error decoding JSON:", err)
	}
	fmt.Printf("\nDecoded Struct:\n%+v\n", decodedUser)

	// === 3. Decode JSON to map ===
	raw := `{"id":3,"name":"Bob","email":"bob@example.com"}`
	var data map[string]interface{}
	err = json.Unmarshal([]byte(raw), &data)
	if err != nil {
		log.Fatal("Error decoding to map:", err)
	}
	fmt.Println("\nDecoded Map:")
	for k, v := range data {
		fmt.Printf("%s: %v\n", k, v)
	}

	// === 4. Encode map to JSON ===
	encodedMap, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatal("Error encoding map:", err)
	}
	fmt.Println("\nRe-encoded Map JSON:\n", string(encodedMap))

	// === 5. Handle JSON array ===
	jsonArray := `[{"id":4,"name":"Charlie"},{"id":5,"name":"Diana"}]`
	var users []User
	err = json.Unmarshal([]byte(jsonArray), &users)
	if err != nil {
		log.Fatal("Error decoding JSON array:", err)
	}
	fmt.Println("\nDecoded Users from JSON Array:")
	for _, u := range users {
		fmt.Printf("%+v\n", u)
	}

	// Encode users slice back to JSON
	usersJSON, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		log.Fatal("Error encoding users slice:", err)
	}
	fmt.Println("\nRe-encoded Users JSON Array:\n", string(usersJSON))

	// === 6. Error handling with invalid JSON ===
	invalidJSON := `{"id":6,"name":"Eve",}` // Trailing comma is invalid
	var invalidUser User
	err = json.Unmarshal([]byte(invalidJSON), &invalidUser)
	if err != nil {
		fmt.Println("\nError decoding invalid JSON:", err)
	} else {
		fmt.Printf("Decoded invalid JSON (unexpectedly succeeded): %+v\n", invalidUser)
	}

	// === 7. Streaming JSON encoding/decoding ===
	// Encoding to a stream (stdout)
	fmt.Println("\nStreaming JSON Encoding:")
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(user); err != nil {
		log.Fatal("Error streaming JSON encoding:", err)
	}

	// Decoding from a stream (stdin)
	// Uncomment below lines to test decoding from stdin
	/*
		fmt.Println("\nEnter JSON data:")
		decoder := json.NewDecoder(os.Stdin)
		var inputUser User
		if err := decoder.Decode(&inputUser); err != nil {
			log.Fatal("Error streaming JSON decoding:", err)
		}
		fmt.Printf("Decoded from stream: %+v\n", inputUser)
	*/

	// Note: The stdin decoding part is commented out to avoid blocking execution.
}
