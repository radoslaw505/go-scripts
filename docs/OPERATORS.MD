# Operators in Go

Go provides a variety of operators for performing operations on variables and values. Below is an overview of the main operator categories in Go.

## 1. Arithmetic Operators

| Operator | Description      | Example      |
|----------|------------------|-------------|
| `+`      | Addition         | `a + b`     |
| `-`      | Subtraction      | `a - b`     |
| `*`      | Multiplication   | `a * b`     |
| `/`      | Division         | `a / b`     |
| `%`      | Modulus          | `a % b`     |
| `++`     | Increment        | `a++`       |
| `--`     | Decrement        | `a--`       |

## 2. Relational Operators

| Operator | Description      | Example      |
|----------|------------------|-------------|
| `==`     | Equal to         | `a == b`    |
| `!=`     | Not equal to     | `a != b`    |
| `>`      | Greater than     | `a > b`     |
| `<`      | Less than        | `a < b`     |
| `>=`     | Greater or equal | `a >= b`    |
| `<=`     | Less or equal    | `a <= b`    |

## 3. Logical Operators

| Operator | Description      | Example      |
|----------|------------------|-------------|
| `&&`     | Logical AND      | `a && b`    |
| `||`     | Logical OR       | `a || b`    |
| `!`      | Logical NOT      | `!a`        |

## 4. Bitwise Operators

| Operator | Description      | Example      |
|----------|------------------|-------------|
| `&`      | AND              | `a & b`     |
| `|`      | OR               | `a | b`     |
| `^`      | XOR              | `a ^ b`     |
| `&^`     | AND NOT          | `a &^ b`    |
| `<<`     | Left shift       | `a << 2`    |
| `>>`     | Right shift      | `a >> 2`    |

## 5. Assignment Operators

| Operator | Description      | Example      |
|----------|------------------|-------------|
| `=`      | Assign           | `a = b`     |
| `+=`     | Add and assign   | `a += b`    |
| `-=`     | Subtract and assign| `a -= b`  |
| `*=`     | Multiply and assign| `a *= b`  |
| `/=`     | Divide and assign| `a /= b`    |
| `%=`     | Modulus and assign| `a %= b`   |
| `&=`     | AND and assign   | `a &= b`    |
| `|=`     | OR and assign    | `a |= b`    |
| `^=`     | XOR and assign   | `a ^= b`    |
| `<<=`    | Left shift assign| `a <<= 2`   |
| `>>=`    | Right shift assign| `a >>= 2`  |
| `&^=`    | AND NOT assign   | `a &^= b`   |

## 6. Miscellaneous Operators

- **Address-of**: `&a` (returns the address of variable `a`)
- **Pointer dereference**: `*a` (returns the value at address `a`)
- **Channel receive**: `<-ch` (receives value from channel `ch`)

## Example

```go
a := 10
b := 3

sum := a + b      // 13
isEqual := a == b // false
a++               // 11
bitwise := a & b  // 2
```

Refer to the [Go specification](https://golang.org/ref/spec#Operators) for more details.