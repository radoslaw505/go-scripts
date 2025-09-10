package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	original := "Hello, Radosław!"

	// Standard Base64 encoding
	encoded := base64.StdEncoding.EncodeToString([]byte(original))
	fmt.Println("Encoded:", encoded)

	// Decoding
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Decode error:", err)
		return
	}
	fmt.Println("Decoded:", string(decoded))

	// URL-safe encoding
	urlEncoded := base64.URLEncoding.EncodeToString([]byte("https://golang.org"))
	fmt.Println("URL-safe:", urlEncoded)
}
