# Pointers in Go

## What is a Pointer?

A pointer is a variable that stores the memory address of another variable. In Go, pointers provide a way to share and modify data efficiently.

## Declaring Pointers

```go
var p *int // p is a pointer to an int
```

## Getting the Address of a Variable

Use the `&` operator to get the address:

```go
x := 10
p := &x // p now holds the address of x
```

## Dereferencing a Pointer

Use the `*` operator to access the value pointed to:

```go
fmt.Println(*p) // prints 10
*p = 20         // sets x to 20
```

## Zero Value of a Pointer

The zero value of a pointer is `nil`:

```go
var p *int
fmt.Println(p) // prints <nil>
```

## Pointers and Functions

Pointers allow functions to modify variables outside their scope:

```go
func increment(n *int) {
    *n++
}

x := 5
increment(&x)
fmt.Println(x) // prints 6
```

## No Pointer Arithmetic

Go does not support pointer arithmetic like C/C++.

## Pointers to Structs

```go
type Point struct { X, Y int }
p := &Point{1, 2}
p.X = 10 // shorthand for (*p).X = 10
```

## When to Use Pointers

- To modify a variable inside a function
- To avoid copying large structs or arrays
- To represent optional values (`nil`)

## Best Practices

- Avoid using pointers unless necessary
- Do not use pointers to interface types

## Further Reading

- [Go Blog: Pointers](https://blog.golang.org/pointers)
- [Go Tour: Pointers](https://tour.golang.org/moretypes/1)