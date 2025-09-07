# Using `init` Functions in Go

The `init` function in Go is a special function that is executed automatically when a package is initialized. It is commonly used for setup tasks such as initializing variables, registering types, or performing checks before the main execution starts.

## Key Points

- Each Go source file can have multiple `init` functions.
- `init` functions have no arguments and no return values.
- They are executed after all variable declarations in the package and before the `main` function.
- The order of execution for multiple `init` functions in the same file is top-down as they appear in the file.
- Across multiple files in a package, the order is not specified.

## Example

```go
package main

import "fmt"

var message string

func init() {
    message = "Hello from init!"
    fmt.Println("init function executed")
}

func main() {
    fmt.Println(message)
}
```

**Output:**
```
init function executed
Hello from init!
```

## Best Practices

- Use `init` only for package-level initialization.
- Avoid complex logic in `init` functions.
- Prefer explicit initialization in `main` or exported functions when possible for clarity and testability.

## References

- [Go Blog: Package Initialization](https://blog.golang.org/init)
- [Go Specification: Package Initialization](https://golang.org/ref/spec#Package_initialization)