# Package json

## Index

- [func Decode(json string) any](#Decode)
- [func Encode(value any) string](#Encode)
- [func Null() Null](#Null)

### Decode

```
func Decode(json string) any
```

No documentation.

### Encode

```
func Encode(value any) string
```

Encode transforms a value into a JSON string.

### Null

```
func Null() Null
```

Null provides a value and interface for JSON "null". It can be used as a
literal:

```
assert(json.Encode(json.Null()) == "null")
```

Or, it can be used as an interface for existing types:

```
func Person(name string) {
func IsJSONNull() bool {
return ^name == ""
}
}
```

