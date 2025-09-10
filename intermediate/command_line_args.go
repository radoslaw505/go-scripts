package main

import (
	"flag"
	"fmt"
)

func main() {
	task := flag.String("task", "", "Task to add")
	priority := flag.Int("priority", 1, "Priority level (1-3)")
	verbose := flag.Bool("verbose", false, "Enable verbose output")

	flag.Parse()

	if *task == "" {
		fmt.Println("Please provide a task using -task")
		return
	}

	if *verbose {
		fmt.Printf("Task: %s\nPriority: %d\n", *task, *priority)
	} else {
		fmt.Printf("Task added: %s\n", *task)
	}
}
