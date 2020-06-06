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
   * [Language Specification](#language-specification)
      * [Comments](#comments)
      * [Data Types](#data-types)
      * [Literals](#literals)
         * [Booleans](#booleans)
         * [Characters](#characters)
         * [Data](#data)
         * [Numbers](#numbers)
         * [Strings](#strings)

<!-- Added by: elliot, at: Sat Jun  6 14:24:39 EDT 2020 -->

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


Language Specification
======================

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

Literals
--------

### Booleans

A boolean can be either `true` or `false`.

### Characters

A character represents a single symbol and is wrapped in single-quotes (`'`').
The symbol must be a UTF-8 code point.

Examples:

- `'a'` - valid ASCII and UTF-8.
- `'😃'` - valid UTF-8.
- `'hi'` - not valid.

### Data

A data value represents zero or more bytes. Data can be created with literals by
using backticks (`\``).

Examples:

- `\`\`` - zero bytes.
- `\`hello\`` - 5 bytes.
- `\`😃\`` - 4 bytes.

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

### Strings

Strings can be any length (including zero) and must be wrapped in double-quotes
(`"`).

Examples:

- `""` - an empty string (zero length).
- `"hello world"` - a string containing 11 characters.
