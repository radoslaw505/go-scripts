# Variadic Functions in Go

Variadic functions are functions that accept a variable number of arguments. In Go, you can define a variadic function by using an ellipsis (`...`) before the type of the last parameter.

## Syntax

```go
func functionName(args ...Type) {
    // function body
}
```

## Example

```go
package main

import "fmt"

func sum(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}

func main() {
    fmt.Println(sum(1, 2, 3))        // Output: 6
    fmt.Println(sum(4, 5, 6, 7, 8))  // Output: 30
}
```

## Key Points

- The variadic parameter must be the last parameter in the function signature.
- Inside the function, the variadic parameter is a slice of the specified type.
- You can pass a slice to a variadic function using the `...` syntax: `sum(nums...)`.

## Use Cases

- Formatting functions like `fmt.Println` and `fmt.Printf`.
- Aggregation functions (e.g., sum, min, max).
- Functions that need to handle an arbitrary number of arguments.

## References

- [Go by Example: Variadic Functions](https://gobyexample.com/variadic-functions)
- [Official Go Documentation](https://golang.org/doc/)
