package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// 1. Run a simple command
	fmt.Println("Running 'echo Hello, Go!'")
	cmd1 := exec.Command("echo", "Hello, Go!")
	output1, err := cmd1.Output()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Output:", string(output1))
	}

	// 2. Capture output from a command
	fmt.Println("\nRunning 'date'")
	cmd2 := exec.Command("date")
	output2, err := cmd2.Output()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Current date:", string(output2))
	}

	// 3. Send input to a command and read transformed output
	fmt.Println("\nRunning 'tr a-z A-Z' with input")
	var out bytes.Buffer
	cmd3 := exec.Command("tr", "a-z", "A-Z")
	cmd3.Stdin = bytes.NewBufferString("hello, world!")
	cmd3.Stdout = &out
	err = cmd3.Run()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Transformed output:", out.String())
	}

	// 4. Run a command in a different directory
	fmt.Println("\nRunning 'ls -l' in /tmp")
	cmd4 := exec.Command("ls", "-l")
	cmd4.Dir = "/tmp"
	output4, err := cmd4.Output()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Contents of /tmp:\n", string(output4))
	}

	// 5. Run a command with custom environment variables
	fmt.Println("\nRunning 'printenv LANG'")
	cmd5 := exec.Command("printenv", "LANG")
	cmd5.Env = append(os.Environ(), "LANG=pl_PL.UTF-8")
	output5, err := cmd5.Output()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("LANG environment:", string(output5))
	}

	// 6. Check if a command exists
	fmt.Println("\nChecking if 'nosuchcommand' exists")
	_, err = exec.LookPath("nosuchcommand")
	if err != nil {
		fmt.Println("Command not found.")
	} else {
		fmt.Println("Command exists.")
	}
}
