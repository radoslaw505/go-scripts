package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// Copy from string to stdout
	r := strings.NewReader("Hello, Radosław!\n")
	io.Copy(os.Stdout, r)

	// ReadAll
	data, _ := io.ReadAll(strings.NewReader("Read this fully"))
	fmt.Println("ReadAll:", string(data))

	// Pipe example
	pr, pw := io.Pipe()
	go func() {
		io.WriteString(pw, "Piped message")
		pw.Close()
	}()
	buf := new(bytes.Buffer)
	buf.ReadFrom(pr)
	fmt.Println("Pipe output:", buf.String())
}
