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
      * [For](#for)
      * [If/Else](#ifelse)
      * [Switch](#switch)
      * [Arrays](#arrays)
      * [Maps](#maps)

<!-- Added by: elliot, at: Fri Jun 19 14:38:08 EDT 2020 -->

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
    // Strings, which can be added together with +.
    print("he" + "llo")

    // Numbers.
    print("1+1 =", 1+1)
    print("7.0/3.0 =", 7.0/3.0)

    // Booleans.
    print(true and false)
    print(true or false)
    print(not true)
}
```

```
$ ok run values
hello
1+1 = 2
7.0/3.0 = 2.3
false
true
false
````

Variables
---------

In ok, variables are explicitly declared and used by the compiler to e.g. check
type-correctness of function calls.

Every variable has a type, but that type is inferred from the expression or
value being assigned to it.

```
func main() {

    // Declare a variable.
    a = "initial"
    print(a)

    // ok will infer the type of initialized variables.
    b = true
    print(b)

    c = 1.23
    print(c)
}
```

```
$ ok run variables
initial
true
1.23
```

For
---

for is ok's only looping construct. Here are some basic types of for loops.

```
func main() {

    // The most basic type, with a single condition.
    i = 1
    for i <= 3 {
        print(i)
        i = i + 1
    }

    // A classic initial/condition/after for loop.
    for j = 7; j <= 9; ++j {
        print(j)
    }

    // for without a condition will loop repeatedly until you break out of the
    // loop or return from the enclosing function.
    for {
        print("loop")
        break
    }

    // You can also continue to the next iteration of the loop.
    for n = 0; n <= 5; ++n {
        if n%2 == 0 {
            continue
        }
        print(n)
    }
}
```

```
$ ok run for
1
2
3
7
8
9
loop
1
3
5
```

If/Else
-------

Branching with if and else in ok is straight-forward. In ok, there is no
"if else". Instead you should use a switch if there are multiple cases.

```
func main() {

    // Here's a basic example.
    if 7%2 == 0 {
        print("7 is even")
    } else {
        print("7 is odd")
    }

    // You can have an if statement without an else.
    if 8%4 == 0 {
        print("8 is divisible by 4")
    }

    // Note that you don't need parentheses around conditions in ok, but that
    // the braces are required.
    num = 9
    if num < 10 {
        print(num, "has 1 digit")
    } else {
        print(num, "has multiple digits")
    }
}
```

```
$ ok run if-else
7 is odd
8 is divisible by 4
9 has 1 digit
```

Switch
------

Switch statements express conditionals across many branches.

```
func main() {

    // Here's a basic switch.
    i = 2
    print("Write", i, "as")
    switch i {
        case 1 {
            print("one")
        }
        case 2 {
            print("two")
        }
        case 3 {
            print("three")
        }
    }

    // You can use commas to separate multiple expressions in the same case
    // statement. We use the optional else case in this example as well.
    weekday = "sunday"
    switch weekday {
        case "saturday", "sunday" {
            print("It's the weekend")
        }
        else {
            print("It's a weekday")
        }
    }

    // switch without an expression is an alternate way to express if/else
    // logic. Here we also show how the case expressions can be non-constants.
    hour = 11
    switch {
        case hour < 12 {
            print("It's before noon")
        }
        else {
            print("It's after noon")
        }
    }
}
```

```
$ ok run switch
Write 2 as
two
It's the weekend
It's before noon
```

Arrays
------

Arrays are an ordered sequence of values.

```
func main() {
    // Arrays are defined within square brackets. If all the elements are the
    // same type, then the type of the array is inferred. In this case the type
    // of "a" will be "[]number".
    a = [1, 2, 3]
    print(a)

    // Arrays can contain mixed types but the type of the array must be given
    // explicitly to indicate this.
    b = []any [123, "foo", true]

    // Arrays are always printed as valid JSON. This makes it easier and more
    // natural to pass data to other systems.
    print(b)

    // Access an element in an array by using its index. Remember, 0 is the
    // index for the first element.
    print(a[0], b[1])

    // Assigning an element works the same way with the index.
    a[1] = 7
    print(a)
}
```

```
$ ok run arrays
[1, 2, 3]
[123, "foo", true]
1 foo
[1, 7, 3]
```

Maps
----

Maps store an unordered collection of key-value pairs. The keys are strings and
must be unique within the same map. The values can be of any type and there is
no restriction on duplicate values.

```
func main() {
    // Maps are defined within curly brackets. If all the elements are the
    // same type, then the type of the map is inferred. In this case the type
    // of "a" will be "{}number".
    a = {"a": 1, "b": 2, "c": 3}
    print(a)

    // Like arrays, maps can contain mixed values by it has to be declared
    // explicitly.
    b = {}any {"a": 123, "b": "foo", "c": true}

    // Maps are always printed as valid JSON with their keys sorted. This makes
    // it easier and more natural to pass data to other systems.
    print(b)

    // Access an element in an map by using its key. Maps only support strings
    // as the key.
    print(a["b"], b["b"])

    // Assigning an element works the same way with the key.
    a["b"] = 7
    print(a)
}
```

```
{"a": 1, "b": 2, "c": 3}
{"a": 123, "b": "foo", "c": true}
2 foo
{"a": 1, "b": 7, "c": 3}
```
