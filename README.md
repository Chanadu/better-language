# Better-Language

## Description

Better-Language is a new programming language built from scratch in Go using a tree-walking interpreter.

Try it in the [interactive playground](https://better-language-playground.pages.dev/), or view the [playground source code](https://github.com/Chanadu/better-language-playground).

## How to run

1. Clone the repository.
2. Open a terminal and navigate to the directory containing the repository.
3. Build the interpreter:

   ```bash
   go build -o gbpl .
   ```

4. Create a `.bpl` file, write your code, and run it:

   ```bash
   ./gbpl <filename>.bpl
   ```

   Alternatively, start the very limited REPL environment:

   ```bash
   ./gbpl
   ```

## Why

I created this project to learn more about how programming languages are created and how interpreters work. This was a personal project and is not connected to any other organization or activity.

## Syntax

The language uses syntax similar to the C programming language. It currently supports the following features.

### Variables

| Feature | Example |
| --- | --- |
| Declaration | `var x = 5` |
| Assignment | `x = 10` |

### Functions

| Feature | Example |
| --- | --- |
| Declaration with arguments | `function add(x, y) {}` |
| Call | `add(1, 2)` |
| Return value | `return x + y` |

Function recursion is also supported.

### Control flow

| Feature | Example |
| --- | --- |
| If and else statements | `if (x > 5) {} else {}` |
| For loops | `for (var i = 0; i < 5; i = i + 1) {}` |
| While loops | `while (x < 5) {}` |
| Ternary expressions | `x > 5 ? 10 : 20` |

### Output and scope

| Feature | Example |
| --- | --- |
| Print statement | `print(x)` |
| Scope | `var x = 5; { var x = 10; }` |

### Arithmetic operations

| Operation | Example |
| --- | --- |
| Addition | `x + y` |
| Subtraction | `x - y` |
| Multiplication | `x * y` |
| Division | `x / y` |
| Modulus | `x % y` |

### Logical operations

| Operation | Example |
| --- | --- |
| And | `x && y` |
| Or | `x \|\| y` |
| Not | `!x` |

### Comparison operations

| Operation | Example |
| --- | --- |
| Greater than | `x > y` |
| Greater than or equal to | `x >= y` |
| Less than | `x < y` |
| Less than or equal to | `x <= y` |
| Equal to | `x == y` |
| Not equal to | `x != y` |

### Bitwise operators

| Operator | Example |
| --- | --- |
| Bitwise AND | `x & y` |
| Bitwise OR | `x \| y` |
| Bitwise XOR | `x ^ y` |
| Bitwise NOT | `~x` |
| Bitwise left shift | `x << y` |
| Bitwise right shift | `x >> y` |

### Data types

| Type | Example |
| --- | --- |
| Integer | `var x = 5` |
| Boolean | `var x = true` |
| String | `var x = "Hello, World!"` |

### Comments

Single-line comments are supported:

```bpl
// This is a comment
```

### Built-in functions

- `clock()` returns the current Unix time in milliseconds.

## Features

- [x] Scanner
- [x] Parser
- [x] Interpreter
