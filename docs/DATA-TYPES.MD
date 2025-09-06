# Data Types in Go

Go (Golang) is a statically typed language, meaning every variable has a specific type. Understanding Go's data types is essential for effective programming.

## Basic Data Types

| Type      | Description                        | Example           |
|-----------|------------------------------------|-------------------|
| `bool`    | Boolean (true or false)            | `var b bool = true` |
| `string`  | Sequence of Unicode characters     | `var s string = "Go"` |
| `int`     | Signed integer (platform-dependent)| `var i int = 42`  |
| `int8`    | 8-bit signed integer               | `var i int8 = -128` |
| `int16`   | 16-bit signed integer              | `var i int16 = 32767` |
| `int32`   | 32-bit signed integer              | `var i int32 = 2147483647` |
| `int64`   | 64-bit signed integer              | `var i int64 = 9223372036854775807` |
| `uint`    | Unsigned integer                   | `var u uint = 42` |
| `uint8`   | 8-bit unsigned integer (byte)      | `var u uint8 = 255` |
| `uint16`  | 16-bit unsigned integer            | `var u uint16 = 65535` |
| `uint32`  | 32-bit unsigned integer            | `var u uint32 = 4294967295` |
| `uint64`  | 64-bit unsigned integer            | `var u uint64 = 18446744073709551615` |
| `float32` | 32-bit floating point              | `var f float32 = 3.14` |
| `float64` | 64-bit floating point              | `var f float64 = 3.1415926535` |
| `complex64` | Complex numbers with float32 real and imaginary parts | `var c complex64 = 1+2i` |
| `complex128`| Complex numbers with float64 real and imaginary parts | `var c complex128 = 1+2i` |
| `byte`    | Alias for `uint8`                  | `var b byte = 255` |
| `rune`    | Alias for `int32` (Unicode code point) | `var r rune = '世'` |

## Composite Types

- **Array**: Fixed-size sequence of elements of the same type.
    ```go
    var arr [3]int = [3]int{1, 2, 3}
    ```
- **Slice**: Dynamically-sized, flexible view into arrays.
    ```go
    var s []int = []int{1, 2, 3}
    ```
- **Map**: Unordered collection of key-value pairs.
    ```go
    var m map[string]int = map[string]int{"a": 1, "b": 2}
    ```
- **Struct**: Collection of fields.
    ```go
    type Person struct {
            Name string
            Age  int
    }
    ```
- **Pointer**: Stores the memory address of a value.
    ```go
    var p *int = &i
    ```

## Interface Type

- **interface{}**: Can hold values of any type.
    ```go
    var x interface{} = "hello"
    ```

## Function Type

Functions are first-class citizens and can be assigned to variables.
```go
var fn func(int) int = func(x int) int { return x * x }
```

## Type Conversion

Go requires explicit type conversion:
```go
var a int = 10
var b float64 = float64(a)
```

## Resources

- [Go Data Types - Official Documentation](https://golang.org/ref/spec#Types)
- [Tour of Go: Types](https://tour.golang.org/basics/11)