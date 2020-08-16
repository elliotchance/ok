[![GitHub release](https://img.shields.io/github/release/elliotchance/ok.svg)](https://github.com/elliotchance/ok/releases/)
[![Build Status](https://travis-ci.org/elliotchance/ok.svg?branch=master)](https://travis-ci.org/elliotchance/ok)
[![codecov](https://codecov.io/gh/elliotchance/ok/branch/master/graph/badge.svg)](https://codecov.io/gh/elliotchance/ok)
[![License: MIT](https://img.shields.io/badge/license-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Join the chat at https://gitter.im/ok-lang/community](https://badges.gitter.im/ok-lang/community.svg)](https://gitter.im/ok-lang/community?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

- [Why OK?](https://github.com/elliotchance/ok/wiki/Why-OK%3F)
- [Installation](https://github.com/elliotchance/ok/wiki/Installation)
- [Learn By Example](https://github.com/elliotchance/ok/wiki/Learn-By-Example)
- [IDE Support](https://github.com/elliotchance/ok/wiki/IDE-Support)
- [Frequently Asked Questions (FAQ)](https://github.com/elliotchance/ok/wiki/FAQ)
- [CLI Commands](https://github.com/elliotchance/ok/wiki/CLI-Commands)
- [Help](https://github.com/elliotchance/ok/wiki/Help)

**ok** is a strongly-duck-typed language. It has the rigidness of a strongly
typed language with more runtime flexibility and simplified . See
[Why Go?](https://github.com/elliotchance/ok/wiki/Why-OK%3F) for language
highlights, or if you just want to browse some examples head over to
[Learn By Example](https://github.com/elliotchance/ok/wiki/Learn-By-Example).

```
func main() {
    print("hello world")
}
```

It is heavily inspired by Go, especially around how interfaces are implicit.
However, ok takes this one step further by merging the concepts of a `struct`
and `interface` into a `func` declaration. This idea is not new, but it means
that all cases where a instance is expected actually is an implicit interface.

```
func Person(Name string) Person {
    func Greeting() string {
        return "Hello, " + ^Name
    }
}

func Impersonator() Impersonator {
    Name = "Secret Agent"

    func Greeting() string {
        return "Hello, Joe Bloggs"
    }
}

func SayHello(p Person) {
    print(p.Greeting())
}

func main() {
    SayHello(Person("Joe Bloggs"))

    // Impersonator satisfies the interface of a Person, so we can use that
    // implicitly instead:
    SayHello(Impersonator())
}
```

Having all inputs being implicit interfaces makes testing really easy and
flexible. Also, ok treats testing as a first-class citizen with syntax:

```
// add.ok
func add(a, b number) number {
    return a + b
}
```

```
// add.okt
test "adding some numbers" {
    assert(add(3, 5) == 8)
    assert(add(3, 5) == 10)
}
```

```
$ ok test
add: add.okt:2:4: adding some numbers: assert(8 == 10) failed
add: 1 failed 0 passed 2 asserts (0ms)
```

For more language features and examples, check out
[Learn By Example](https://github.com/elliotchance/ok/wiki/Learn-By-Example).