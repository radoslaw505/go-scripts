# Go Import Statement

The `import` statement in Go is used to include packages in your program, allowing you to use their exported functions, types, and variables.

## Basic Syntax

```go
import "package/path"
```

Example:
```go
import "fmt"
```

## Multiple Imports

You can import multiple packages using parentheses:

```go
import (
    "fmt"
    "math"
)
```

## Aliased Imports

You can assign an alias to a package:

```go
import io "io/ioutil"
```

## Blank Identifier

Use `_` to import a package solely for its side-effects (e.g., initialization):

```go
import _ "net/http/pprof"
```

## Grouping Imports

Standard library and third-party packages are often grouped separately for clarity.

```go
import (
    "fmt"
    "os"

    "github.com/some/module"
)
```

## Notes

- Imports must be at the top of the file, after the `package` declaration.
- Unused imports cause a compilation error.

For more details, see the [Go documentation](https://golang.org/doc/effective_go#imports).