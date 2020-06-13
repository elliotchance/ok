[![GitHub release](https://img.shields.io/github/release/elliotchance/ok.svg)](https://github.com/elliotchance/ok/releases/)
[![Build Status](https://travis-ci.org/elliotchance/ok.svg?branch=master)](https://travis-ci.org/elliotchance/ok)
[![codecov](https://codecov.io/gh/elliotchance/ok/branch/master/graph/badge.svg)](https://codecov.io/gh/elliotchance/ok)
[![License: MIT](https://img.shields.io/badge/license-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

**ok** is a strongly-duck-typed language, heavily influenced by Go. The goals
are:

1. **Strongly-typed**, but can use type inference in most places.
2. **Only decimal math** for absolute precision in all calculations.
3. **Syntax that is extremely simple** and very fast to parse for compilation.
4. TODO: **Avoids code that can lead to common runtime bugs** - no global
variables, nils, dereferencing or variables/arguments that have defaults.
5. TODO: **Functions are also objects and interfaces as themselves**.
6. TODO: **Testing is first-class**.

<!--ts-->
   * [Installation](#installation)
      * [Precompiled Binaries](#precompiled-binaries)
      * [go get](#go-get)
      * [From Source](#from-source)
   * [Command Line Interface](#command-line-interface)
      * [run](#run)
      * [version](#version)
   * [Learn By Example](#learn-by-example)
      * [Hello World](#hello-world)
      * [Values](#values)
      * [Variables](#variables)
   * [Language Specification](#language-specification)
      * [Built-in Functions](#built-in-functions)
      * [Comments](#comments)
      * [Data Types](#data-types)
      * [Expressions](#expressions)
      * [Literals](#literals)
         * [Booleans](#booleans)
         * [Characters](#characters)
         * [Data](#data)
         * [Numbers](#numbers)
         * [Strings](#strings)
      * [Operators](#operators)
      * [Variables](#variables-1)

<!-- Added by: elliot, at: Sat Jun 13 18:14:07 EDT 2020 -->

<!--te-->


Installation
============


Precompiled Binaries
--------------------

You can find ready to go binaries for Mac, Windows and Linux on the
[Releases page](https://github.com/elliotchance/ok/releases).

These do not require any dependencies.


go get
------

If you have Go installed, you can install or update the latest version of ok
with:

```bash
go get -u github.com/elliotchance/ok
```


From Source
-----------

You will need to have Go 1.14+ installed to build from source.


Command Line Interface
======================

run
---

Programs in ok are directories containing one or more `.ok` files. You can run a
program by specifying the directory (or see the included tests/):

```bash
ok run my/program
```

version
-------

`ok version` will show the current version and the date it was built.


Learn By Example
================

These have been heavily influenced (copied) from
[gobyexample](https://gobyexample.com) because it's such a great source!


Hello World
-----------

Our first program will print the classic "hello world" message. Here’s the full
source code.

```
func main() {
    print("hello world")
}
```

To run the program, put the code in `hello-world/main.ok` and use `ok run`.

```bash
$ ok run hello-world
hello world
```

Now that we can run and build basic ok programs, let’s learn more about the
language.

Values
------

ok has various value types including strings, numbers, booleans, etc. Here are a
few basic examples.

```
func main() {
    print("go" + "lang")

    print("1+1 =", 1+1)
    print("7.0/3.0 =", 7.0/3.0)

    print(true and false)
    print(true or false)
    print(not true)
}
```

Variables
---------

In ok, variables are explicitly declared and used by the compiler to e.g. check
type-correctness of function calls.

Every variable has a type, but that type is inferred from the expression or
value being assigned to it.

```
func main() {

    a = "initial"
    print(a)

    b = true
    print(b)

    c = 1.23
    print(c)
}
```


Language Specification
======================

Built-in Functions
------------------

- `print(...any)` - Prints a single line with each argument separated by a
space.

Comments
--------

Comments begin with `//` and are terminated by a new line. Comments can exist on
the same line as code:

```
// This is a comment.
print("Hi") // Also a comment
```

Data Types
----------

ok supports these data types:

- `bool`: [Booleans](#booleans)
- `char`: [Characters](#characters)
- `data`: [Data](#data)
- `number`: [Numbers](#numbers)
- `string`: [Strings](#strings)

Expressions
-----------

An expression resolves to a single value and can be recursive. An expression is
one of:

1. A [literal](#Literals).
2. `(` [expression](#Expressions) `)`.
3. [expression](#Expressions) [binary operator](#Operators) [expression](#Expressions).
4. [unary operator](#Operators) [expression](#Expressions).

Literals
--------

### Booleans

A boolean can be either `true` or `false`.

### Characters

A character represents a single symbol and is wrapped in single-quotes (`'`).
The symbol must be a UTF-8 code point.

Examples:

- `'a'` - valid ASCII and UTF-8.
- `'😃'` - valid UTF-8.
- `'hi'` - not valid.

### Data

A data value represents zero or more bytes. Data can be created with literals by
using backticks (`` ` ``).

Examples:

- ``` `` ``` - zero bytes.
- `` `hello` `` - 5 bytes.
- `` `😃` `` - 4 bytes.

### Numbers

Numbers can be represented in two forms. Each allows an optional proceeding
negative sign:

- Integers: `1234`, `0`, `-230`, etc.
- Decimals: `1.23`, `17.0`, `0.0001`, etc.

Numbers are exact (unlike IEEE 754 floating point approximations) and have no
practical limitation in magnitude or precision. The precision of a number can be
defined explicitly in the literal based on the number of digits after the `.`,
including zeros:

- `12` - precision is 0.
- `12.3` - precision is 1.
- `12.00` - precision is 2.

Zeros on the left will always be removed. For example `0012.3` will be reduced
to `12.3`. However, zeros on the right will never be removed because that would
modify the precision. For example `0.1300` will remain as such. This represents
a decimal with a precision of 4.

Numbers cannot start with a `.`. For example `.1` is not valid. However, `0.1`
is valid.

When printing numbers the precision is preserved. So `1.300` will also display
as `1.300`.

Number have the value of a negative zero (`-0`). This is always normalised as
`0`.

Divide and remainder with a denominator of zero (ie. division by zero) will
raise an error.

### Strings

Strings can be any length (including zero) and must be wrapped in double-quotes
(`"`).

Examples:

- `""` - an empty string (zero length).
- `"hello world"` - a string containing 11 characters.


Operators
---------

The following table describes the supported binary operators. All binary
operations require the same type on the left and right.

|       | `bool` | `char` | `data` | `number` | `string` |
| ----- | ------ | ------ | ------ | -------- | -------- |
| `+`   | No     | No     | Yes    | Yes      | Yes      |
| `-`   | No     | No     | No     | Yes      | No       |
| `*`   | No     | No     | No     | Yes      | No       |
| `/`   | No     | No     | No     | Yes      | No       |
| `%`   | No     | No     | No     | Yes      | No       |
| `and` | Yes    | No     | No     | No       | No       |
| `or`  | Yes    | No     | No     | No       | No       |
| `==`  | Yes    | Yes    | Yes    | Yes      | Yes      |
| `!=`  | Yes    | Yes    | Yes    | Yes      | Yes      |
| `>`   | No     | Yes    | No     | Yes      | Yes      |
| `>=`  | No     | Yes    | No     | Yes      | Yes      |
| `<`   | No     | Yes    | No     | Yes      | Yes      |
| `<=`  | No     | Yes    | No     | Yes      | Yes      |

The following table describes the unary operators.

|       | `bool` | `char` | `data` | `number` | `string` |
| ----- | ------ | ------ | ------ | -------- | -------- |
| `-`   | No     | No     | No     | Yes      | No       |
| `not` | Yes    | No     | No     | No       | No       |

When evaluating expressions the order of operations is influenced by the
precedence of the operator. The precedence from most to least important:

1. `*`, `/`, `%`
2. `+`, `-`
3. `==`, `!=`, `>`, `>=`, `<`, `<=`
4. `and`
5. `or`

Examples:

- `1 + 2 + 3` is evaluated as `(1 + 2) + 3`
- `1 + 2 * 3` is evaluated as `1 + (2 * 3)`
- `1 * 2 + 3` is evaluated as `(1 * 2) + 3`

Other notes:

1. All arithmetic operations return a number that has the same precision as the
greatest precision input.

2. The remainder operator (`%`) is not the same as a modulo operator. Whereas as
modulo operation will always return a positive number, a remainder may be
negative if one of the inputs is also negative.


Variables
---------

A variable holds a single value.

Variables must be defined (by assigning) before they can be referenced and may
be reassigned as long as the new value is of the same type.

A variable name is case-sensitive, must start with a letter and can contain
letters, digits and an underscore (`_`).
