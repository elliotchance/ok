# Package unicode

## Index

- [func IsControl(c char) bool](#IsControl)
- [func IsDigit(c char) bool](#IsDigit)
- [func IsGraphic(c char) bool](#IsGraphic)
- [func IsLetter(c char) bool](#IsLetter)
- [func IsLower(c char) bool](#IsLower)
- [func IsMark(c char) bool](#IsMark)
- [func IsNumber(c char) bool](#IsNumber)
- [func IsPrint(c char) bool](#IsPrint)
- [func IsPunct(c char) bool](#IsPunct)
- [func IsSpace(c char) bool](#IsSpace)
- [func IsSymbol(c char) bool](#IsSymbol)
- [func IsTitle(c char) bool](#IsTitle)
- [func IsUpper(c char) bool](#IsUpper)
- [func ToLower(c char) char](#ToLower)
- [func ToTitle(c char) char](#ToTitle)
- [func ToUpper(c char) char](#ToUpper)

### IsControl

```
func IsControl(c char) bool
```

IsControl reports whether the character is a control character.

### IsDigit

```
func IsDigit(c char) bool
```

IsDigit reports whether the character is a decimal digit.

### IsGraphic

```
func IsGraphic(c char) bool
```

IsGraphic reports whether the character is defined as a Graphic by Unicode.
Such characters include letters, marks, numbers, punctuation, symbols, and
spaces, from categories L, M, N, P, S, Zs.

### IsLetter

```
func IsLetter(c char) bool
```

IsLetter reports whether the character is a letter (category L).

### IsLower

```
func IsLower(c char) bool
```

IsLower reports whether the character is a lower case letter.

### IsMark

```
func IsMark(c char) bool
```

IsMark reports whether the character is a mark character (category M).

### IsNumber

```
func IsNumber(c char) bool
```

IsNumber reports whether the character is a number (category N).

### IsPrint

```
func IsPrint(c char) bool
```

IsPrint reports whether the character is defined as printable. Such
characters include letters, marks, numbers, punctuation, symbols, and the
ASCII space character, from categories L, M, N, P, S and the ASCII space
character.

This categorization is the same as IsGraphic except that the only spacing
character is ASCII space, U+0020.

### IsPunct

```
func IsPunct(c char) bool
```

IsPunct reports whether the character is a Unicode punctuation character
(category P).

### IsSpace

```
func IsSpace(c char) bool
```

IsSpace reports whether the character is a space character as defined by
Unicode's White Space property; in the Latin-1 space this is '\t', '\n',
'\v', '\f', '\r', ' ', U+0085 (NEL), U+00A0 (NBSP).

Other definitions of spacing characters are set by category Z and property
Pattern_White_Space.

### IsSymbol

```
func IsSymbol(c char) bool
```

IsSymbol reports whether the character is a symbolic character.

### IsTitle

```
func IsTitle(c char) bool
```

IsTitle reports whether the character is a title case letter.

### IsUpper

```
func IsUpper(c char) bool
```

IsUpper reports whether the character is an upper case letter.

### ToLower

```
func ToLower(c char) char
```

ToLower maps the character to lower case.

### ToTitle

```
func ToTitle(c char) char
```

ToTitle maps the character to title case.

### ToUpper

```
func ToUpper(c char) char
```

ToUpper maps the character to upper case.

