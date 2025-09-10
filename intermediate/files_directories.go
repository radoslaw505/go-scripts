package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Write to file
	f, _ := os.Create("demo.txt")
	fmt.Println("File created:", f.Name())
	defer f.Close()
	f.WriteString("Hello, Radosław!\nWelcome to Go file handling.\n")

	// Read and filter lines
	file, _ := os.Open("demo.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "Go") {
			fmt.Println("Filtered:", scanner.Text())
		}
	}

	// File info
	info, _ := os.Stat("demo.txt")
	fmt.Println("File size:", info.Size())
	fmt.Println("Name:", info.Name())
	fmt.Println("IsDir:", info.IsDir())
	fmt.Println("ModTime:", info.ModTime())
	fmt.Println("Mode:", info.Mode())

	// Modify permissions
	err := os.Chmod("demo.txt", 0644)
	if err != nil {
		fmt.Println("Error changing permissions:", err)
	} else {
		fmt.Println("Permissions changed to 0644")
	}

	// Modify timestamps
	err = os.Chtimes("demo.txt", info.ModTime(), info.ModTime())
	if err != nil {
		fmt.Println("Error changing times:", err)
	} else {
		fmt.Println("Timestamps updated:", info.ModTime())
	}

	// Move/Rename file
	err = os.Rename("demo.txt", "renamed_demo.txt")
	if err != nil {
		fmt.Println("Error renaming file:", err)
	} else {
		fmt.Println("File renamed to: renamed_demo.txt")
	}

	// Delete file
	err = os.Remove("renamed_demo.txt")
	if err != nil {
		fmt.Println("Error deleting file:", err)
	} else {
		fmt.Println("File deleted")
	}

	// Directory operations
	err = os.Mkdir("demoDir", 0755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
	} else {
		fmt.Println("Directory created: demoDir")
	}
	err = os.Remove("demoDir")
	if err != nil {
		fmt.Println("Error removing directory:", err)
	} else {
		fmt.Println("Directory removed: demoDir")
	}

	// Temp file
	tmp, _ := os.CreateTemp("", "demo-*.txt")
	defer tmp.Close()
	fmt.Println("Temp file:", tmp.Name())

	// Temp dir
	dir, _ := os.MkdirTemp("", "demo-dir-*")
	fmt.Println("Temp dir:", dir)
}
