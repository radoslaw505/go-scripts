# 📘 Working with Strings and Runes in Go

Go (Golang) provides powerful tools for handling text through its `string` and `rune` types. Understanding how these types work is essential for building robust applications that deal with Unicode, text manipulation, and encoding.

---

## 🧵 Strings in Go

### What is a String?

A `string` in Go is a read-only slice of bytes. It represents a sequence of UTF-8 encoded characters.

```go
str := "Hello, 世界"
fmt.Println(str) // Output: Hello, 世界
```

### String Properties

- Immutable: Strings cannot be changed once created.
- UTF-8 encoded: Each character may take 1–4 bytes.

### Common String Operations

```go
s := "golang"

// Length in bytes
fmt.Println(len(s)) // 6

// Accessing individual bytes
fmt.Println(s[0]) // 103 ('g')

// String concatenation
s2 := s + " rocks"
fmt.Println(s2) // golang rocks
```

---

## 🔡 Runes in Go

### What is a Rune?

A `rune` is an alias for `int32` and represents a Unicode code point.

```go
var r rune = '世'
fmt.Println(r)        // 19990
fmt.Printf("%c\n", r) // 世
```

### Why Use Runes?

- Useful for working with Unicode characters.
- Allows iteration over characters rather than bytes.

### Iterating Over Runes

```go
str := "Hello, 世界"

for i, r := range str {
    fmt.Printf("Index: %d, Rune: %c\n", i, r)
}
```

---

## 🔁 Converting Between Strings and Runes

### String to Rune Slice

```go
str := "Go语言"
runes := []rune(str)
fmt.Println(runes) // [71 111 35821 35328]
```

### Rune Slice to String

```go
runes := []rune{71, 111, 35821, 35328}
str := string(runes)
fmt.Println(str) // Go语言
```

---

## 🧪 Unicode and Encoding

Go handles Unicode natively with UTF-8 encoding. This means:

- Multibyte characters are supported.
- You must be careful when slicing strings by index.

```go
str := "世界"
fmt.Println(str[0]) // 228 (not '世', but the first byte of UTF-8 encoding)
```

To safely work with characters, convert to `[]rune`.

---

## 🛠️ Useful Packages

- `strings`: Provides functions for string manipulation.
- `unicode/utf8`: Helps with UTF-8 encoding and decoding.
- `strconv`: Converts between strings and other types.

---

## ✅ Best Practices

- Use `rune` when dealing with characters, especially non-ASCII.
- Avoid slicing strings by index unless you're sure of byte boundaries.
- Prefer `range` loops for character-safe iteration.

---

## 📚 Example: Counting Characters

```go
str := "Hello, 世界"
fmt.Println("Byte length:", len(str))
fmt.Println("Rune count:", utf8.RuneCountInString(str))
```

---

Happy coding in Go! 🚀
```

Let me know if you'd like this turned into a tutorial, cheat sheet, or even a mini course.
