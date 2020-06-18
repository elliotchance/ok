package ast

import "strconv"

// Literal represents a literal of any type.
type Literal struct {
	Kind  string
	Value string
	Array []*Literal
}

// NewLiteralData create a new literal representing a data value.
func NewLiteralData(data []byte) *Literal {
	return &Literal{
		Kind:  "data",
		Value: string(data),
	}
}

// NewLiteralNumber create a new literal representing a number value.
func NewLiteralNumber(number string) *Literal {
	return &Literal{
		Kind:  "number",
		Value: number,
	}
}

// NewLiteralString create a new literal representing a string value.
func NewLiteralString(str string) *Literal {
	return &Literal{
		Kind:  "string",
		Value: str,
	}
}

// NewLiteralBool create a new literal representing a boolean value.
func NewLiteralBool(b bool) *Literal {
	return &Literal{
		Kind:  "bool",
		Value: strconv.FormatBool(b),
	}
}

// NewLiteralChar create a new literal representing a character value.
func NewLiteralChar(c rune) *Literal {
	return &Literal{
		Kind:  "char",
		Value: string(c),
	}
}
