# Exiting a Program in Go

In Go, you can terminate a program using the `os.Exit()` function from the `os` package. This function immediately stops program execution and returns a specified exit code to the operating system.

## Usage

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Println("This will be printed.")
    os.Exit(1)
    fmt.Println("This will NOT be printed.")
}
```

- Any code after `os.Exit()` will **not** be executed.
- The argument to `os.Exit()` is the exit code:
  - `0` indicates success.
  - Non-zero values indicate an error.

## Notes

- Deferred functions are **not** run when `os.Exit()` is called.
- For graceful shutdowns, ensure cleanup is done **before** calling `os.Exit()`.

## Example: Exiting on Error

```go
if err != nil {
    fmt.Fprintln(os.Stderr, "Error:", err)
    os.Exit(1)
}
```

## References

- [os.Exit documentation](https://pkg.go.dev/os#Exit)
- [Effective Go: Program Exit](https://go.dev/doc/effective_go#exit)