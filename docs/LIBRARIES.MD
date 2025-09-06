# Go Library Documentation

## Overview

This document provides an overview of the libraries used in this project, including their purpose and usage examples.

---

## Table of Contents

- [Standard Libraries](#standard-libraries)
- [Third-Party Libraries](#third-party-libraries)
- [How to Add a New Library](#how-to-add-a-new-library)

---

## Standard Libraries

| Package      | Description                       |
|--------------|-----------------------------------|
| `fmt`        | Formatting for I/O                |
| `os`         | OS functions and file operations  |
| `net/http`   | HTTP client and server            |
| `encoding/json` | JSON encoding and decoding     |

---

## Third-Party Libraries

| Package         | Version | Description                  | Usage Example                |
|-----------------|---------|------------------------------|------------------------------|
| `github.com/sirupsen/logrus` | v1.9.0 | Structured logging           | `logrus.Info("message")`     |
| `github.com/spf13/viper`     | v1.14.0| Configuration management     | `viper.GetString("key")`     |
| `github.com/stretchr/testify` | v1.8.0 | Testing utilities            | `assert.Equal(t, 1, 1)`      |

---

## How to Add a New Library

1. Run: `go get github.com/username/library`
2. Import the library in your Go code as needed.
3. Run: `go mod tidy` to update dependencies.
4. Build your script with: `go build path/to/your/script.go`

**Example workflow for using third-party libraries:**

```bash
# Initialize a Go module (if not already done)
go mod init <module-name>

# Add a third-party library
go get github.com/sirupsen/logrus

# Tidy up dependencies
go mod tidy

# Build your script
go build basics/hello.go
```

The Go toolchain will handle dependencies using the `go.mod` and `go.sum` files automatically.

---

## Example

```go
import (
    "github.com/sirupsen/logrus"
)

func main() {
    logrus.Info("Hello, world!")
}
```

---

## References

- [Go Modules Reference](https://blog.golang.org/using-go-modules)
- [pkg.go.dev](https://pkg.go.dev/)
