# Go Naming Conventions

This document outlines the recommended naming conventions for Go (Golang) codebases.

## General Guidelines

- **Be concise, but descriptive.**
- **Use mixedCaps (camelCase) or CapitalizedWords (PascalCase).**
- **Avoid underscores (`_`) in names.**
- **Use full words instead of abbreviations, unless the abbreviation is widely known (e.g., `URL`, `ID`).**

## Packages

- Package names should be short, lowercase, and meaningful.
- Avoid plurals and underscores.
- Example: `http`, `json`, `math`

## Variables

- Use short, meaningful names.
- Use `camelCase` for local variables.
- Example: `userName`, `count`, `err`

## Constants

- Use `CamelCase` or `ALL_CAPS` for acronyms.
- Example: `MaxLength`, `DefaultTimeout`, `HTTPStatusOK`

## Functions

- Use `camelCase` for unexported functions.
- Use `PascalCase` for exported functions.
- Function names should be verbs or verb phrases.
- Example: `readFile()`, `WriteToBuffer()`

## Types

- Use `PascalCase` for type names.
- Example: `User`, `Config`, `HTTPClient`

## Interfaces

- Use `PascalCase`.
- Interface names ending with `-er` are common for single-method interfaces.
- Example: `Reader`, `Writer`, `Closer`

## Struct Fields

- Use `PascalCase` for exported fields.
- Use `camelCase` for unexported fields.
- Example: `Name string`, `age int`

## Receivers

- Use short, meaningful names (often one or two letters).
- Example: `func (u *User) Save() error`

## Acronyms

- Capitalize common acronyms (e.g., `HTTP`, `URL`, `ID`).
- Example: `userID`, `ParseURL()`

## Example

```go
package user

type User struct {
    ID   int
    Name string
}

func (u *User) Save() error {
    // implementation
}
```

## References

- [Effective Go - Naming](https://go.dev/doc/effective_go#names)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments#naming)