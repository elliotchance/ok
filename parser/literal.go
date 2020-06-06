package parser

import (
	"errors"
	"ok/ast"
	"ok/lexer"
)

func consumeLiteral(f *File, offset int) (*ast.Literal, int, error) {
	types := []string{
		lexer.TokenBool,
		lexer.TokenCharacter,
		lexer.TokenData,
		lexer.TokenNumber,
		lexer.TokenString,
	}

	for _, ty := range types {
		if t := f.Tokens[offset]; t.Kind == ty {
			literal := &ast.Literal{
				Kind:  ty,
				Value: t.Value,
			}

			return literal, offset + 1, nil
		}
	}

	return nil, offset, errors.New("no literal found")
}

func validateLiteral(literal *ast.Literal) error {
	switch literal.Kind {
	case lexer.TokenCharacter:
		if len(literal.Value) == 0 {
			return errors.New("character literal cannot be empty")
		}
	}

	return nil
}
