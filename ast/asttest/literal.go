package asttest

import (
	"strconv"

	"github.com/elliotchance/ok/ast"
)

// NewLiteralData create a new literal representing a data value.
func NewLiteralData(data []byte) *ast.Literal {
	return &ast.Literal{
		Kind:  "data",
		Value: string(data),
	}
}

// NewLiteralNumber create a new literal representing a number value.
func NewLiteralNumber(number string) *ast.Literal {
	return &ast.Literal{
		Kind:  "number",
		Value: number,
	}
}

// NewLiteralString create a new literal representing a string value.
func NewLiteralString(str string) *ast.Literal {
	return &ast.Literal{
		Kind:  "string",
		Value: str,
	}
}

// NewLiteralBool create a new literal representing a boolean value.
func NewLiteralBool(b bool) *ast.Literal {
	return &ast.Literal{
		Kind:  "bool",
		Value: strconv.FormatBool(b),
	}
}

// NewLiteralChar create a new literal representing a character value.
func NewLiteralChar(c rune) *ast.Literal {
	return &ast.Literal{
		Kind:  "char",
		Value: string(c),
	}
}
