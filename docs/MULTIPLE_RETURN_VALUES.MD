# Multiple Return Values in Go

Go supports functions that return multiple values, a feature that enables concise error handling and more expressive APIs.

## Basic Syntax

```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}
```

## Usage Example

```go
result, err := divide(10, 2)
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Println("Result:", result)
}
```

## Named Return Values

You can name return values for clarity:

```go
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}
```

## Ignoring Values

Use `_` to ignore unwanted return values:

```go
value, _ := divide(10, 2)
```

## Common Use Cases

- Error handling (`value, err := ...`)
- Returning multiple related results

## Summary

Multiple return values make Go functions flexible and help simplify error handling patterns.