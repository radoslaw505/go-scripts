# Maps in Go

Maps in Go are unordered collections of key-value pairs. They are similar to dictionaries in other languages.

## Declaration and Initialization

```go
// Declare a map with string keys and int values
var m map[string]int

// Initialize a map using make
m = make(map[string]int)

// Or use a map literal
n := map[string]int{
    "foo": 1,
    "bar": 2,
}
```

## Basic Operations

- **Insert or Update:**  
  `m["key"] = value`

- **Retrieve:**  
  `val := m["key"]`

- **Delete:**  
  `delete(m, "key")`

- **Check Existence:**  
  ```go
  val, ok := m["key"]
  if ok {
      // key exists
  }
  ```

## Iterating Over a Map

```go
for k, v := range m {
    fmt.Println(k, v)
}
```

## Notes

- The zero value of a map is `nil`. You must initialize it before use.
- Maps are not safe for concurrent use. Use synchronization primitives if needed.
- The order of iteration is not guaranteed.

## Example

```go
package main

import "fmt"

func main() {
    capitals := map[string]string{
        "Poland": "Warsaw",
        "France": "Paris",
    }
    capitals["Germany"] = "Berlin"
    for country, capital := range capitals {
        fmt.Printf("%s: %s\n", country, capital)
    }
}
```