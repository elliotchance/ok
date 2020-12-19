# reflect

- [func Call(fn any, args []any) []any](#Call)
- [func Get(obj any, prop any) any](#Get)
- [func Interface(value any) string](#Interface)
- [func Kind(value any) string](#Kind)
- [func Len(value any) number](#Len)
- [func Properties(obj any) []string](#Properties)
- [func Set(obj any, prop any, value any) any](#Set)
- [func Type(value any) string](#Type)

## Call

```
func Call(fn any, args []any) []any
```

Call can be used to call a function variable without knowing the type.

`fn` must be a callable function and args must have the correct number of
arguments for the function. The number of return values will depend on the
function being called.

Example:

```
fn = func(a, b number) number {
return a + b
}

print(reflect.Call(fn, []any [1.2, 3.4]))
```

## Get

```
func Get(obj any, prop any) any
```

Get performs one of the following (depending on the input type):

- Arrays: `prop` must be an number that represents the index in the array. If
it is not a number or is out of bound then an error is raised.

- Maps: `prop` must be a string that represents the key in the map. If prop
is not a string of the key does not exist in the map an error will be raised.

- Objects: `prop` must be a string representing the property name. You may
only access public properties (starting with a capital letter). Trying to
access a property that does not exist or is not public will result in an
error.

Any other type will always result in an error.

## Interface

```
func Interface(value any) string
```

Interface returns the canonical interface as a string.

Examples:

```
{ }
{ Name string }
{ Greet(string) bool; Name string }
```

## Kind

```
func Kind(value any) string
```

Kind returns the runtime type of a value. One of; "bool", "char", "data",
"number", "string", "array", "map" or "func".

## Len

```
func Len(value any) number
```

Len returns the number of elements in an array or map. If the value is not an
array or map then an error is raised.

## Properties

```
func Properties(obj any) []string
```

Properties returns the public properties of an object, or the keys in a map.
If the input is not an object or map, an error is raised. The values returned
will be sorted in either case.

## Set

```
func Set(obj any, prop any, value any) any
```

Set performs one of the following (depending on the input type):

- Arrays: `prop` must be an number that represents the index in the array. If
it is not a number or is out of bound then an error is raised.

- Maps: `prop` must be a string that represents the key in the map. If prop
is not a string an error will be raised. However, if the key does not exist
it will be created. If the key already exists its value will be replaced.

- Objects: `prop` must be a string representing the property name. You may
only access public properties (starting with a capital letter). Trying to
access a property that does not exist or is not public will result in an
error.

Any other type will always result in an error.

## Type

```
func Type(value any) string
```

Type returns the runtime type of a value, some examples would be:

```
reflect.Type(3)                      == "number"
reflect.Type(["foo", "bar"])         == "[]string"
reflect.Type(func hello(ok bool) {}) == "func(bool)"
```

