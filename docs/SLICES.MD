# Slices in Go

Slices are a key data structure in Go, providing a flexible and powerful way to work with sequences of elements.

## What is a Slice?

A slice is a dynamically-sized, flexible view into the elements of an array. Unlike arrays, slices do not have a fixed size.

```go
var s []int // declares a slice of ints, initially nil
s = []int{1, 2, 3}
```

## Creating Slices

- **From an array:**
    ```go
    arr := [5]int{1, 2, 3, 4, 5}
    s := arr[1:4] // s == []int{2, 3, 4}
    ```
- **Using `make`:**
    ```go
    s := make([]int, 3) // length 3, capacity 3
    ```

## Slice Properties

- **Length:** Number of elements in the slice (`len(s)`)
- **Capacity:** Number of elements in the underlying array, counting from the start of the slice (`cap(s)`)

## Appending to Slices

Use `append` to add elements:
```go
s := []int{1, 2}
s = append(s, 3) // s == []int{1, 2, 3}
```

## Copying Slices

Use `copy` to copy elements:
```go
src := []int{1, 2, 3}
dst := make([]int, len(src))
copy(dst, src)
```

## Slices are References

Slices reference the underlying array. Modifying a slice may affect other slices referencing the same array.

```go
a := [3]int{1, 2, 3}
s1 := a[:]
s2 := a[:]
s1[0] = 100
// a[0] == 100, s2[0] == 100
```

## Nil vs. Empty Slices

- Nil slice: `var s []int` (len and cap are 0, s == nil)
- Empty slice: `s := []int{}` (len and cap are 0, s != nil)

## Common Pitfalls

- Slices share underlying arrays; be careful with modifications.
- Appending may allocate a new array if capacity is exceeded.

## Further Reading

- [Go Slices: usage and internals](https://blog.golang.org/slices-intro)
- [Go Slices Reference](https://golang.org/ref/spec#Slice_types)