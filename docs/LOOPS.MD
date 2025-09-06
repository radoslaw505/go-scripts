# Loops in Go

Go provides only one looping construct: the `for` loop. It can be used in several ways to achieve different looping behaviors.

## Basic `for` Loop

```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

## While-Style Loop

Go does not have a `while` keyword, but you can use `for` as a while loop:

```go
i := 0
for i < 5 {
    fmt.Println(i)
    i++
}
```

## Infinite Loop

Use `for` without any condition for an infinite loop:

```go
for {
    fmt.Println("Looping forever")
}
```

## Looping Over Collections

### Slices or Arrays

```go
numbers := []int{1, 2, 3}
for index, value := range numbers {
    fmt.Printf("Index: %d, Value: %d\n", index, value)
}
```

### Maps

```go
m := map[string]int{"a": 1, "b": 2}
for key, value := range m {
    fmt.Printf("Key: %s, Value: %d\n", key, value)
}
```

## Breaking and Continuing

- `break` exits the loop.
- `continue` skips to the next iteration.

```go
for i := 0; i < 5; i++ {
    if i == 2 {
        continue
    }
    if i == 4 {
        break
    }
    fmt.Println(i)
}
```

## Summary

- Go uses only the `for` loop for all looping needs.
- You can use `for` as a traditional loop, a while loop, or an infinite loop.
- Use `range` to iterate over slices, arrays, maps, and strings.
