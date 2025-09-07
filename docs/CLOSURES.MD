# Closures in Go

Closures are functions that reference variables from outside their own scope. In Go, functions are first-class citizens, which means you can assign them to variables, pass them as arguments, and return them from other functions. This enables the use of closures.

## Basic Example

```go
func adder() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x
        return sum
    }
}

func main() {
    add := adder()
    fmt.Println(add(10)) // 10
    fmt.Println(add(5))  // 15
}
```

Here, the returned function "closes over" the `sum` variable, preserving its state between calls.

## Use Cases

- **Stateful functions:** Maintain state across multiple invocations.
- **Callbacks:** Pass logic as arguments to other functions.
- **Encapsulation:** Hide implementation details.

## Variable Capture

Closures capture variables by reference, not by value. Be careful when using closures inside loops:

```go
func main() {
    funcs := []func(){}
    for i := 0; i < 3; i++ {
        funcs = append(funcs, func() { fmt.Println(i) })
    }
    for _, f := range funcs {
        f() // Prints "3" three times
    }
}
```

To capture the current value, use a local variable:

```go
for i := 0; i < 3; i++ {
    v := i
    funcs = append(funcs, func() { fmt.Println(v) })
}
```

## Summary

Closures are a powerful feature in Go for creating flexible, stateful, and encapsulated functions. Use them thoughtfully, especially when capturing variables in loops.
