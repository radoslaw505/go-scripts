# Using `panic` and `recover` in Go

Go provides the `panic` function as a way to handle unexpected errors and exceptional situations. When `panic` is called, the normal execution of the current function stops, and the program begins to unwind the stack, running any deferred functions before the program crashes.

## When to Use `panic`

- For unrecoverable errors (e.g., out-of-bounds access, nil pointer dereference).
- When the program cannot continue safely.
- Not for regular error handling—use `error` values for expected errors.

## Basic Example

```go
package main

import "fmt"

func main() {
    defer fmt.Println("This will run even after panic")
    fmt.Println("About to panic!")
    panic("Something went wrong")
    fmt.Println("This will not be printed")
}
```

**Output:**
```
About to panic!
This will run even after panic
panic: Something went wrong
...
```

## Recovering from Panic

You can recover from a panic using the `recover` function, typically inside a deferred function.

```go
func safeDivide(a, b int) {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()
    fmt.Println(a / b)
}

func main() {
    safeDivide(10, 0)
    fmt.Println("Program continues...")
}
```

## Best Practices

- Use `panic` only for truly exceptional cases.
- Always clean up resources using `defer`.
- Prefer returning errors for expected error conditions.

For more details, see the [Go blog on errors are values](https://blog.golang.org/errors-are-values).