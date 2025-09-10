package main

import (
	"embed"
	"fmt"
	"net/http"
)

// === 1. Embed a single file as string ===

//go:embed hello.txt
var helloText string

// === 2. Embed a single file as []byte ===

//go:embed hello.txt
var helloBytes []byte

// === 3. Embed multiple files into embed.FS ===

//go:embed static/*
var staticFiles embed.FS

func main() {
	// Print embedded string
	fmt.Println("Embedded string from hello.txt:")
	fmt.Println(helloText)

	// Print embedded bytes
	fmt.Println("\nEmbedded bytes from hello.txt:")
	fmt.Println(string(helloBytes))

	// Read embedded HTML file
	html, err := staticFiles.ReadFile("static/index.html")
	if err != nil {
		fmt.Println("Error reading index.html:", err)
		return
	}
	fmt.Println("\nEmbedded HTML content:")
	fmt.Println(string(html))

	// Serve embedded files via HTTP
	fmt.Println("\nStarting server at http://localhost:8080/")
	http.Handle("/", http.FileServer(http.FS(staticFiles)))
	http.ListenAndServe(":8080", nil)
}
