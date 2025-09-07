# 🧾 Go Formatting Verbs Cheat Sheet

Go’s `fmt` package provides a rich set of **formatting verbs** that control how values are printed. These verbs are used with functions like `fmt.Printf`, `fmt.Sprintf`, and `fmt.Fprintf`.

---

## 📦 Importing the Package

```go
import "fmt"
```

---

## 🔤 General Verbs

| Verb | Description                          | Example Output |
|------|--------------------------------------|----------------|
| `%v` | Default format                       | `true`, `123`, `map[x:1]` |
| `%+v`| Adds field names to structs          | `{Name:John Age:30}` |
| `%#v`| Go syntax representation             | `main.Person{Name:"John", Age:30}` |
| `%T` | Type of the value                    | `int`, `string`, `[]float64` |
| `%%` | Literal percent sign                 | `%`            |

---

## 🔢 Integer Verbs

| Verb | Description         | Example Output |
|------|---------------------|----------------|
| `%d` | Decimal             | `42`           |
| `%b` | Binary              | `101010`       |
| `%o` | Octal               | `52`           |
| `%x` | Hex (lowercase)     | `2a`           |
| `%X` | Hex (uppercase)     | `2A`           |
| `%c` | Character (Unicode) | `*`            |
| `%q` | Quoted character    | `'A'`          |

---

## 🔡 String and Byte Slice Verbs

| Verb | Description         | Example Output |
|------|---------------------|----------------|
| `%s` | Plain string        | `hello`        |
| `%q` | Quoted string       | `"hello"`      |
| `%x` | Hex (lowercase)     | `68656c6c6f`   |
| `%X` | Hex (uppercase)     | `68656C6C6F`   |

---

## 🔣 Floating-Point Verbs

| Verb | Description         | Example Output |
|------|---------------------|----------------|
| `%f` | Decimal point       | `3.141593`     |
| `%e` | Scientific notation | `3.141593e+00` |
| `%E` | Scientific (caps)   | `3.141593E+00` |
| `%g` | Compact format      | `3.14159`      |
| `%G` | Compact (caps)      | `3.14159`      |

---

## 🧪 Boolean Verbs

| Verb | Description         | Example Output |
|------|---------------------|----------------|
| `%t` | Boolean value       | `true` / `false` |

---

## 🧰 Pointer Verbs

| Verb | Description         | Example Output |
|------|---------------------|----------------|
| `%p` | Pointer address     | `0xc000010230` |

---

## 🧵 Width and Precision

You can control **width** and **precision** using flags:

```go
fmt.Printf("%6.2f", 3.14159) // Output: "  3.14"
```

| Flag | Meaning                  |
|------|--------------------------|
| `-`  | Left-align               |
| `+`  | Always show sign         |
| `0`  | Pad with zeros           |
| `#`  | Alternate format         |
| `.`  | Precision for floats     |

---

## 🧪 Example Usage

```go
name := "Radosław"
age := 30
fmt.Printf("Name: %s, Age: %d\n", name, age)
```

```go
pi := 3.14159
fmt.Printf("Pi: %6.2f\n", pi) // Output: "  3.14"
```

---

## 🧠 Pro Tip

Use `fmt.Sprintf` to format strings without printing:

```go
formatted := fmt.Sprintf("Hello, %s!", "world")
fmt.Println(formatted)
```

---

## 📚 Resources

- [`fmt` package documentation](https://pkg.go.dev/fmt)
- Try it live: [Go Playground](https://go.dev/play)

---
