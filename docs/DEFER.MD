# Using `defer` in Go

The `defer` statement in Go schedules a function call to be run after the function completes, just before it returns. It is commonly used for cleanup actions, such as closing files or releasing resources.

## Basic Syntax

```go
defer functionName(arguments)
```

## Example

```go
package main

import (
    "fmt"
)

func main() {
    fmt.Println("Start")
    defer fmt.Println("Deferred")
    fmt.Println("End")
}
```

**Output:**
```
Start
End
Deferred
```

## Key Points

- Deferred calls are executed in **Last-In-First-Out (LIFO)** order.
- Useful for resource management (e.g., closing files, unlocking mutexes).
- Arguments to deferred functions are evaluated immediately, but the function call is executed later.

## Example: File Handling

```go
file, err := os.Open("example.txt")
if err != nil {
    log.Fatal(err)
}
defer file.Close()
// Use file
```

## Multiple Defers

```go
defer fmt.Println("First")
defer fmt.Println("Second")
// "Second" will print before "First"
```

## When to Use

- Releasing resources (files, network connections)
- Unlocking mutexes
- Logging function exit

## References

- [Go Blog: Defer, Panic, and Recover](https://blog.golang.org/defer-panic-and-recover)
- [Official Documentation](https://golang.org/doc/effective_go#defer)