package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// === 1. Create hello.txt ===
	err := os.WriteFile("hello.txt", []byte("Hello, Radosław!\nThis file is embedded using go:embed."), 0644)
	if err != nil {
		fmt.Println("Error creating hello.txt:", err)
		return
	}

	// === 2. Create static directory ===
	err = os.MkdirAll("static", os.ModePerm)
	if err != nil {
		fmt.Println("Error creating static directory:", err)
		return
	}

	// === 3. Create static/index.html ===
	htmlContent := `<!DOCTYPE html>
<html>
<head>
    <title>Embedded Page</title>
    <link rel="stylesheet" href="style.css">
</head>
<body>
    <h1>Hello from embedded HTML!</h1>
</body>
</html>`
	err = os.WriteFile(filepath.Join("static", "index.html"), []byte(htmlContent), 0644)
	if err != nil {
		fmt.Println("Error creating index.html:", err)
		return
	}

	// === 4. Create static/style.css ===
	cssContent := `body {
    background-color: #f0f0f0;
    font-family: sans-serif;
}`
	err = os.WriteFile(filepath.Join("static", "style.css"), []byte(cssContent), 0644)
	if err != nil {
		fmt.Println("Error creating style.css:", err)
		return
	}

	// === 5. Read and display contents ===
	fmt.Println("\n📄 hello.txt:")
	hello, _ := os.ReadFile("hello.txt")
	fmt.Println(string(hello))

	fmt.Println("\n🌐 static/index.html:")
	index, _ := os.ReadFile("static/index.html")
	fmt.Println(string(index))

	fmt.Println("\n🎨 static/style.css:")
	style, _ := os.ReadFile("static/style.css")
	fmt.Println(string(style))
}
