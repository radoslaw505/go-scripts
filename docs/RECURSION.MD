# Recursion in Go

Recursion is a programming technique where a function calls itself to solve a problem by breaking it down into smaller subproblems. In Go, recursion is supported like in most modern languages.

## Basic Structure

A recursive function typically has:
- **Base case:** Stops the recursion.
- **Recursive case:** Calls itself with modified arguments.

```go
func recursiveFunction(params) returnType {
    if baseCondition {
        return baseValue
    }
    // Recursive call
    return recursiveFunction(modifiedParams)
}
```

## Example: Factorial

```go
func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}
```

## Example: Fibonacci

```go
func fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}
```

## Considerations

- **Stack Overflow:** Deep recursion can cause stack overflow.
- **Performance:** Recursive solutions may be less efficient than iterative ones.
- **Tail Recursion:** Go does not optimize tail recursion.

## When to Use Recursion

- Problems that can be divided into similar subproblems (e.g., tree traversal, divide and conquer algorithms).
- When code clarity is more important than raw performance.

## References

- [Go by Example: Recursion](https://gobyexample.com/recursion)
- [Effective Go: Recursion](https://golang.org/doc/effective_go#recursion)