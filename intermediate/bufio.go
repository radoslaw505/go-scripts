package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// === 1. Buffered Reading from a File ===
	fmt.Println("Reading from file:")
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Print("Line: ", line)
	}

	// === 2. Buffered Writing to a File ===
	fmt.Println("\nWriting to file:")
	outFile, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer outFile.Close()

	writer := bufio.NewWriter(outFile)
	writer.WriteString("Hello, Radosław!\nThis is written using bufio.Writer.\n")
	writer.Flush() // Important: flush to write buffered data

	// === 3. Scanning Input from stdin ===
	fmt.Println("\nType something (Scanner will read it):")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input := scanner.Text()
		fmt.Println("You typed:", input)
	}

	// === 4. Using bufio.NewReadWriter ===
	fmt.Println("\nUsing ReadWriter:")
	rw := bufio.NewReadWriter(bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout))
	rw.Writer.WriteString("Enter your name: ")
	rw.Writer.Flush()
	name, _ := rw.Reader.ReadString('\n')
	fmt.Println("Hello,", strings.TrimSpace(name))
}
