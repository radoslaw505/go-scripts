# Arithmetic Operators in Go

Go provides several arithmetic operators for performing basic mathematical operations on numeric types.

## List of Arithmetic Operators

| Operator | Description        | Example      | Result |
|----------|-------------------|--------------|--------|
| `+`      | Addition          | `5 + 3`      | `8`    |
| `-`      | Subtraction       | `5 - 3`      | `2`    |
| `*`      | Multiplication    | `5 * 3`      | `15`   |
| `/`      | Division          | `5 / 3`      | `1`    |
| `%`      | Modulus (remainder) | `5 % 3`   | `2`    |
| `++`     | Increment         | `a++`        | `a = a + 1` |
| `--`     | Decrement         | `a--`        | `a = a - 1` |

> **Note:** Division of integers truncates towards zero.

## Example Usage

```go
package main

import "fmt"

func main() {
    a, b := 10, 3

    fmt.Println("a + b =", a + b) // 13
    fmt.Println("a - b =", a - b) // 7
    fmt.Println("a * b =", a * b) // 30
    fmt.Println("a / b =", a / b) // 3
    fmt.Println("a % b =", a % b) // 1

    a++
    fmt.Println("a after increment:", a) // 11

    b--
    fmt.Println("b after decrement:", b) // 2
}
```

## References

- [Go Operators Documentation](https://go.dev/ref/spec#Operators)