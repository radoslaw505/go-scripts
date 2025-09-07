# Functions in Go

Go functions are fundamental building blocks for code organization and reuse. This document covers the basics of defining and using functions in Go.

## Defining a Function

```go
func functionName(parameter1 type1, parameter2 type2) returnType {
    // function body
}
```

**Example:**
```go
func add(a int, b int) int {
    return a + b
}
```

## Multiple Return Values

Go functions can return multiple values.

```go
func swap(x, y string) (string, string) {
    return y, x
}
```

## Named Return Values

You can name return values for clarity.

```go
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}
```

## Variadic Functions

Functions can accept a variable number of arguments.

```go
func sum(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}
```

## Anonymous Functions and Closures

Functions can be defined inline and assigned to variables.

```go
add := func(a, b int) int {
    return a + b
}
result := add(2, 3)
```

## Passing Functions as Arguments

Functions are first-class citizens and can be passed as arguments.

```go
func compute(fn func(int, int) int) int {
    return fn(3, 4)
}
```

## Summary

- Functions are defined with `func`.
- Support multiple and named return values.
- Variadic parameters are allowed.
- Functions can be anonymous and passed as values.

For more details, see the [Go documentation](https://golang.org/doc/effective_go.html#functions).