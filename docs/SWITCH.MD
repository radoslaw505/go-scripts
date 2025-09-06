# `switch` Statement in Go

The `switch` statement in Go is a concise way to perform multi-way branching based on the value of an expression.

## Basic Syntax

```go
switch expression {
case value1:
    // code block
case value2:
    // code block
default:
    // code block
}
```

## Example

```go
package main

import "fmt"

func main() {
    day := 3
    switch day {
    case 1:
        fmt.Println("Monday")
    case 2:
        fmt.Println("Tuesday")
    case 3:
        fmt.Println("Wednesday")
    default:
        fmt.Println("Another day")
    }
}
```

## Features

- **No fallthrough by default:** Only the matching case executes.
- **Multiple expressions per case:** Separate with commas.
- **Switch without an expression:** Acts like a series of `if` statements.
- **Type switch:** Used to compare types in interfaces.

## Type Switch Example

```go
var x interface{} = 7
switch v := x.(type) {
case int:
    fmt.Println("int:", v)
case string:
    fmt.Println("string:", v)
default:
    fmt.Println("unknown type")
}
```

## Resources

- [Go Tour: Switch](https://tour.golang.org/flowcontrol/9)
- [Go Blog: Switch](https://blog.golang.org/switch)
