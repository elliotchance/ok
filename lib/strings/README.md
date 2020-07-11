# strings

- [Contains](#Contains)
- [HasPrefix](#HasPrefix)
- [HasSuffix](#HasSuffix)
- [Index](#Index)
- [IndexAfter](#IndexAfter)
- [Join](#Join)
- [LastIndex](#LastIndex)
- [LastIndexBefore](#LastIndexBefore)
- [Repeat](#Repeat)
- [ReplaceAll](#ReplaceAll)
- [Reverse](#Reverse)
- [Split](#Split)
- [ToLower](#ToLower)
- [ToUpper](#ToUpper)
- [Trim](#Trim)
- [TrimLeft](#TrimLeft)
- [TrimPrefix](#TrimPrefix)
- [TrimRight](#TrimRight)
- [TrimSuffix](#TrimSuffix)

## Contains

```
func Contains(s string, substr string) bool
```

Contains checks whether substr exists in s.

## HasPrefix

```
func HasPrefix(s string, prefix string) bool
```

HasPrefix will return true if s starts with prefix.

## HasSuffix

```
func HasSuffix(s string, suffix string) bool
```

HasSuffix will return true if s ends with suffix.

## Index

```
func Index(s string, substr string) number
```

Index returns the index of the first occurance of substr; or -1 if not found.

## IndexAfter

```
func IndexAfter(s string, substr string, offset number) number
```

IndexAfter returns the index of the first occurance of substr at or after
offset; or -1 if not found. If the offset is zero or negative then this
function will perform as Index.

Example: Iterating through all matches

substr = "foo"
for i = Index(s, substr); i != -1; i = IndexAfter(s, substr, i) {
print(i)
}


## Join

```
func Join(strings []string, glue string) string
```

Join returns a string containing all elements joined with glue. The glue will
only appear between elements, even if those elements are empty strings
themselves. Glue is allowed to be empty.

## LastIndex

```
func LastIndex(s string, substr string) number
```

LastIndex returns the start of the last occurence of substr in s.

## LastIndexBefore

```
func LastIndexBefore(s string, substr string, offset number) number
```

LastIndexBefore returns the start of the last occurence of substr that starts
before offset.

Example: Iterating through all matches

substr = "foo"
for i = LastIndex(s, substr); i != -1; i = LastIndexBefore(s, substr, i) {
print(i)
}


## Repeat

```
func Repeat(str string, times number) string
```

Repeat returns str repeated times. Any values for times that is 0 or below
will result in an empty string.

## ReplaceAll

```
func ReplaceAll(s string, find string, replace string) string
```

ReplaceAll return a string with each of find substituted with replace.

If `find` is empty, then `replace` will be inserted between every character.
It will not be insert before the first character or after the last.

## Reverse

```
func Reverse(s string) string
```

Reverse creates a string with all characters in the opposite order.

## Split

```
func Split(s string, delimiter string) []string
```

Split will always returns one or more elements. If the glue is empty the
string will be split into characters.

TODO(elliot): This is a horribly inefficient algorithm. This was very early
on in the language when there we're barely any features, please clean this
up if you see it.

## ToLower

```
func ToLower(s string) string
```

ToLower returns a lower case version of s.

TODO(elliot): This only works for ASCII characters.

## ToUpper

```
func ToUpper(s string) string
```

ToUpper returns a upper case version of s.

TODO(elliot): This only works for ASCII characters.

## Trim

```
func Trim(s string, cutset string) string
```

Trim returns string with any of the cutset characters removed from the left
(start) and right (end).

## TrimLeft

```
func TrimLeft(s string, cutset string) string
```

TrimLeft returns string with any of the cutset characters removed from the
left (start).

## TrimPrefix

```
func TrimPrefix(s string, prefix string) string
```

TrimPrefix will trim the prefix if it exists at the start of the string, or
return the original string.

If prefix appears more then once at the start of the string only the first
prefix will be removed.

If prefix is equal to s then an empty result will be returned.

## TrimRight

```
func TrimRight(s string, cutset string) string
```

TrimRight returns string with any of the cutset characters removed from the
right (end).

## TrimSuffix

```
func TrimSuffix(s string, suffix string) string
```

TrimSuffix will trim the suffix if it exists at the end of the string, or
return the original string.

If suffix appears more then once at the end of the string only the last
suffix will be removed.

If suffix is equal to s then an empty result will be returned.

