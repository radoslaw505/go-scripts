# if-else Statements in Go

Go provides straightforward syntax for conditional branching using `if`, `else if`, and `else` statements.

## Basic Syntax

```go
if condition {
    // code to execute if condition is true
} else if anotherCondition {
    // code to execute if anotherCondition is true
} else {
    // code to execute if none of the above conditions are true
}
```

## Example

```go
package main

import "fmt"

func main() {
    x := 10

    if x > 0 {
        fmt.Println("x is positive")
    } else if x < 0 {
        fmt.Println("x is negative")
    } else {
        fmt.Println("x is zero")
    }
}
```

## Short Statement in if

You can include a short statement before the condition:

```go
if y := compute(); y > 5 {
    fmt.Println("y is greater than 5")
}
```

## Notes

- Parentheses around the condition are not required.
- Curly braces `{}` are mandatory, even for single statements.

For more details, see the [Go documentation](https://golang.org/doc/effective_go#control-structures).