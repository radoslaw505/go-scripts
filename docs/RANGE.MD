# `range` in Go

The `range` keyword in Go is used to iterate over elements in various data structures such as arrays, slices, maps, strings, and channels.

## Usage Examples

### Arrays and Slices

```go
nums := []int{2, 4, 6}
for i, v := range nums {
    fmt.Println(i, v)
}
```
- `i` is the index, `v` is the value.

### Maps

```go
m := map[string]int{"a": 1, "b": 2}
for k, v := range m {
    fmt.Println(k, v)
}
```
- `k` is the key, `v` is the value.

### Strings

```go
s := "hello"
for i, c := range s {
    fmt.Printf("%d: %c\n", i, c)
}
```
- Iterates over Unicode code points.

### Channels

```go
ch := make(chan int)
go func() {
    for i := 0; i < 3; i++ {
        ch <- i
    }
    close(ch)
}()
for v := range ch {
    fmt.Println(v)
}
```
- Iterates until the channel is closed.

## Notes

- You can ignore the index or value using `_`.
- For slices and arrays, `range` returns index and value.
- For maps, it returns key and value.
- For strings, it returns index and rune.
- For channels, it returns values sent on the channel.

## References

- [Go Language Specification: For statements](https://golang.org/ref/spec#For_statements)
- [Go Tour: Range](https://tour.golang.org/moretypes/16)