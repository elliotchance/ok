[![GitHub release](https://img.shields.io/github/release/elliotchance/ok.svg)](https://github.com/elliotchance/ok/releases/)
[![Build Status](https://travis-ci.org/elliotchance/ok.svg?branch=master)](https://travis-ci.org/elliotchance/ok)
[![codecov](https://codecov.io/gh/elliotchance/ok/branch/master/graph/badge.svg)](https://codecov.io/gh/elliotchance/ok)
[![License: MIT](https://img.shields.io/badge/license-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Join the chat at https://gitter.im/ok-lang/community](https://badges.gitter.im/ok-lang/community.svg)](https://gitter.im/ok-lang/community?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

**ok** is a strongly-duck-typed language. If you notice a lot of similarities
with Go, that is no accident. There are a lot of things I liked about Go that I
brought to OK. Also, OK is written in Go until which time it can be written in
itself.

<!--ts-->
   * [Why OK?](#why-ok)
      * [Strongly Typed](#strongly-typed)
      * [First-class Testing](#first-class-testing)
      * [Everything in a Single Binary](#everything-in-a-single-binary)
   * [Installation](#installation)
      * [Precompiled Binaries](#precompiled-binaries)
      * [go get](#go-get)
      * [From Source](#from-source)
   * [IDE Support](#ide-support)
      * [Visual Studio Code](#visual-studio-code)
   * [Command Line Interface](#command-line-interface)
      * [asm](#asm)
      * [build](#build)
      * [doc](#doc)
      * [run](#run)
      * [test](#test)
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
      * [Iteration](#iteration)
      * [Functions](#functions)
      * [Multiple Return Values](#multiple-return-values)
      * [Interpolation](#interpolation)
      * [Stateful Functions (Objects)](#stateful-functions-objects)
      * [Errors](#errors)
      * [Finally](#finally)


<!--te-->


Why OK?
=======

Strongly Typed
--------------

OK is strongly typed. However, it has less types that most strongly typed
languages because it doesn't directly expose underlying CPU types. These types
are:

* [`bool`](https://github.com/elliotchance/ok/wiki/Language-Specification#booleans)
is a `true` or `false` value.
* [`char`](https://github.com/elliotchance/ok/wiki/Language-Specification#characters)
is a single character. Supports all UTF-8 up to 4 bytes.
* [`data`](https://github.com/elliotchance/ok/wiki/Language-Specification#data)
is raw binary data.
* [`number`](https://github.com/elliotchance/ok/wiki/Language-Specification#numbers)
is a decimal value with an exact value (this is not an approximation like
floating-point types) so all arithmetic and math operation are exact. An
internal maximum precision is used for non-exact calculations.
* [`string`](https://github.com/elliotchance/ok/wiki/Language-Specification#strings)
is a UTF-8 string of any length (including zero).

These five fundamental types can be used in:

* [arrays](https://github.com/elliotchance/ok/wiki/Language-Specification#arrays) -
an ordered sequence of values.
* [maps](https://github.com/elliotchance/ok#maps) -
an unordered collection of key-value pairs


First-class Testing
-------------------

Testing is not just an added library of helpers, but rather it is built into the
language itself.

Here is an example of some OK code we may want to test:

```
// main.ok

func add(a, b number) number {
    return a + b
}
```

Tests are placed in the same package (directory) as the source code, but with a
different extension (`.okt`):

```
// main.okt

test "adding some numbers" {
    assert(add(3, 5) == 8)
    assert(add(3, 5) == 10)
}
```

The second assertion will fail. Here is how we run the tests with the output:

```
$ ok test tests/example-tests
tests/example-tests: adding some numbers: assert(8 == 10) failed
tests/example-tests: 1 failed 0 passed 2 asserts (60.638µs)
```

Everything in a Single Binary
-----------------------------

The binary for the release (you can
[download it from the Releases page](https://github.com/elliotchance/ok/releases))
contains all of the tools and
[standard library](https://github.com/elliotchance/ok/tree/master/lib) so there
is zero installation or configuration. Just place the binary anywhere and start
using it.


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


IDE Support
===========

Visual Studio Code
------------------

You can install the
[ok](https://marketplace.visualstudio.com/items?itemName=elliotchance.ok)
extension inside of VSCode.


Command Line Interface
======================

asm
---

The `asm` tool is useful for debugging. It will print out human-readable
descriptions of each of the compiled instructions for one or more
functions. Usage:

```bash
ok asm package-name function1 function2 ...
```

Without providing any arguments it will use the current package and
output the `main` function, this is equivalent to running:

```bash
ok asm . main
```

Here is the output from the `hello-world` example:

```
func main():
    1 Assign                 # $1 = "Hello, World!"
    2 Print                  # print($1)
```

build
-----

The build tool will compile a package and produce a binary. The binary
is completely standalone and can be run on another computer.

```bash
$ ok build tests/hello-world
$ ./hello-world
Hello, World!
```

doc
---

The doc tool will generate Github-friendly Markdown for a package. To
see what this looks like you can have a look at the
[standard library](https://github.com/elliotchance/ok/tree/master/lib).

```bash
ok doc my/package
```

run
---

Programs in ok are directories containing one or more `.ok` files. You can run a
program by specifying the directory (or see the included `tests/`):

```bash
ok run my/program
```

test
----

Use `test` to run all tests in a package. Tests are placed in the same package
(directory) as the code they will test, but with a different file extension
(`.okt`):

```bash
ok test my/program
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
7.0/3.0 = 2.3333333333333333333
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
$ ok run maps
{"a": 1, "b": 2, "c": 3}
{"a": 123, "b": "foo", "c": true}
2 foo
{"a": 1, "b": 7, "c": 3}
```

Iteration
---------

```
func main() {
    myArray = [7, 11, 13]

    // When iterating an array the first and second variable are assigned the
    // value and the index respectively.
    for v, i in myArray {
        print(i, v)
    }

    myMap = {"foo": 1.23, "bar": 4.56}

    // Maps work the same way but the second variable will be the key.
    for value, key in myMap {
        print(key, value)
    }

    // For both arrays and maps you may omit the second variable if you only
    // need to iterate the values.
    for value in myArray {
        print(value)
    }

    for value in myMap {
        print("value is", value)
    }

    // If you also need to keep a numeric iterator while iterating a map you can
    // use another form of for.
    for i = 0; value, key in myMap; ++i {
        print(i, key, value)
    }
}
```

```
$ ok run iteration
0 7
1 11
2 13
foo 1.23
bar 4.56
7
11
13
value is 1.23
value is 4.56
0 foo 1.23
1 bar 4.56
```

Functions
---------

Functions are central in ok. We'll learn about functions with a few different
examples.

```
// Here's a function that takes two numbers and returns their sum.
func plus(a number, b number) number {
    // ok requires explicit returns, i.e. it won’t automatically return the
    // value of the last expression.
    return a + b
}

// When you have multiple consecutive parameters of the same type, you may omit
// the type name for the like-typed parameters up to the final parameter that
// declares the type.
func plusPlus(a, b, c number) number {
    return a + b + c
}

func main() {
    // Call a function just as you'd expect, with name(args).

    res = plus(1, 2)
    print("1+2 =", res)

    res = plusPlus(1, 2, 3)
    print("1+2+3 =", res)
}
```

```
$ ok run functions
1+2 = 3
1+2+3 = 6
```

There are several other features to Go functions. One is multiple return values,
which we'll look at next.

Multiple Return Values
----------------------

OK has built-in support for multiple return values.

```
// The (number, number) in this function signature shows
// that the function returns 2 numbers.
func vals() (number, number) {
    return 3, 7
}

func main() {
    // Here we use the 2 different return values from the
    // call with multiple assignment.
    a, b = vals()
    print(a)
    print(b)

    // If you only want a subset of the returned values,
    // use the blank identifier _.
    _, c = vals()
    print(c)
}
```

```
$ ok run multiple-return-values
3
7
7
```

Interpolation
-------------

```
import "math"

func main() {
    // Variables can be placed in {}. This is called interpolation.
    name = "Bob"
    print("Hello, {name}. How are you?")

    // Interpolation can be used with any expression. The result doesn't
    // even have to be a string.
    print("3 + 5 = {3+5}")

    // The spaces added between {} are not required, but might be easier
    // to read.
    total = 3.557
    print("The total is ${ math.Round(total, 2) }")
}
```

```
$ ok run interpolation
Hello, Bob. How are you?
3 + 5 = 8
The total is $3.56
```

Stateful Functions (Objects)
----------------------------

```
// When a function has a single return type that is also the same name
// as the function, this is called a constructor. Constructors are used
// to create stateful functions. Sometimes called objects, records or
// structures in other languages.
func person(name string) person {
    // When you set variables here, they will be accessible outside.
    Name = name
    Age = 42

    // A constructor does not need to have a return statement. It will
    // always return itself.
}

func main() {
    // Create an instance by calling the constructor.
    p = person("Bob")

    // You can read and write properties using the dot notation.
    p.Name = "John"
    print(p.Age)

    // Objects are always printed as JSON objects.
    print(p)
}
```

```
$ ok run objects
42
{"Age": 42, "Name": "John"}
```

Errors
------

```
// Errors in ok work very similar to exceptions in some other languages. If you
// are not familar with exceptions, read on.

// Any function can raise an error. Raising an error prevents any more code from
// running, the code will jump to an error handler. If there is no error handler
// in this function it will be raised to the caller, and so on until it's
// handled, as we will see.
func f1(arg number) number {
    if arg == 42 {
        // Error is a build in type that takes a message. Since we do not handle
        // the error here, it will be passed up to the caller.
        raise Error("can't work with 42")
    }

    return arg + 3
}

// Sometimes just a message is not enough. We can create our own custom error
// type by implementing the Error interface. That is, any type that has the
// property Error of type string.
func argError(arg number, prob string) argError {
    // Save these properties for later.
    Arg = arg
    Prob = prob

    // This satisfies the Error interface.
    Error = "{arg} - {prob}"
}

// f2 performs the same logic as f1, but we raise our custom error instead.
func f2(arg number) number {
    if arg == 42 {
        // Our custom error will act like an Error, but we can also inspect it.
        raise argError(arg, "can't work with it")
    }

    return arg + 3
}

func main() {
    for i in [7, 42] {
        try {
            r = f1(i)

            // If f1 raised an error this will not run because the code will
            // jump immediately to the error handler below.
            print("f1 worked:", r)
        } on Error {
            // When an error is caught, a special "err" variable is provided.
            // Since the Error type must include an Error string property, we
            // can print it out now.
            print("f1 failed:", err.Error)
        }
    }

    // Let's try the same thing with f2, which has the same logic, but raises a
    // custom error type.
    for i in [7, 42] {
        try {
            r = f2(i)
            print("f2 worked:", r)
        } on Error {
            // The type of "err" is still an Error, so we cannot access Arg and
            // Prob. We can rely on the Error message through.
            print("f2 failed:", err.Error)
        }
    }

    try {
        f2(42)
    } on argError {
        // The "err" is now a type of argError, so we can inpect those extra
        // properties. Be careful, if the error raised was not of type argError,
        // then it will not be handled here. Instead it would be passed up to
        // the caller to handle.
        print(err.Arg)
        print(err.Prob)
    }

    // You can have multiple handlers. Only the first match will be executed.
    // This let's you provide more generic handlers if specific cases are not
    // raised.
    try {
        f1(42)
    } on argError {
        // f1 raises an Error, so it will not be caught here.
        print("This should not happen!")
    } on Error {
        // This will handle the error.
        print("f1 failed, all we have is the message:", err.Error)
    }
}
```

```
$ ok run errors
f1 worked: 10
f1 failed: can't work with 42
f2 worked: 10
f2 failed: 42 - can't work with it
42
can't work with it
f1 failed, all we have is the message: can't work with 42
```

Finally
-------

```
func f1(arg number) number {
    if arg == 42 {
        raise Error("can't work with 42")
    }

    return arg + 3
}

func main() {
    for i in [7, 42] {
        try {
            r = f1(i)
            print("f1 worked:", r)
        } on Error {
            print("f1 failed:", err.Error)
        } finally {
            // Finally will always run after the try block or the error
            // handler.
            print("finally f1")
        }
    }

    for i in [7, 42] {
        try {
            r = f2(i)
            print("f2 worked:", r)
        } on Error {
            print("f2 failed:", err.Error)
        }
    }
}

func f2(arg number) number {
    try {
        return f1(arg)
    } finally {
        // Finally will always run before the return, even if there is
        // an error.
        print("f2 done")
    }
}
```

```
$ ok run finally
f1 worked: 10
finally f1
f1 failed: can't work with 42
finally f1
f2 done
f2 worked: 10
f2 done
f2 failed: can't work with 42
```
