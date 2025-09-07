# Arrays and Blank Identifier in Go

## Arrays

An **array** in Go is a fixed-size, ordered collection of elements of the same type.

### Declaration

```go
var a [5]int // array of 5 integers, initialized to zero values
b := [3]string{"Go", "Rust", "Python"}
```

### Accessing Elements

```go
a[0] = 10
fmt.Println(a[0]) // Output: 10
```

### Array Properties

- Fixed length (size is part of the type)
- Elements are zero-initialized
- Arrays are value types (assignments copy the entire array)

### Example

```go
var nums [4]int = [4]int{1, 2, 3, 4}
for i, v := range nums {
    fmt.Printf("Index %d: %d\n", i, v)
}
```

## Blank Identifier (`_`)

The **blank identifier** (`_`) is used to ignore values in assignments or imports.

### Use Cases

#### Ignoring Return Values

```go
value, _ := someFunc() // ignore the second return value
```

#### Ignoring Index or Value in Loops

```go
for _, v := range nums {
    fmt.Println(v) // ignore index
}
```

#### Importing for Side Effects

```go
import _ "net/http/pprof" // import only for initialization side effects
```

## Summary

- Arrays are fixed-size, value-type collections.
- The blank identifier (`_`) is used to ignore unwanted values or imports.
