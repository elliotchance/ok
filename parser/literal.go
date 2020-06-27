package parser

import (
	"errors"
	"strings"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/lexer"
)

func consumeLiteral(f *File, offset int) (*ast.Literal, int, error) {
	types := []string{
		lexer.TokenBoolLiteral,
		lexer.TokenCharLiteral,
		lexer.TokenDataLiteral,
		lexer.TokenNumberLiteral,
		lexer.TokenStringLiteral,
	}

	for _, ty := range types {
		if t := f.Tokens[offset]; t.Kind == ty {
			literal := &ast.Literal{
				Kind:  strings.Split(ty, " ")[0],
				Value: t.Value,
			}

			return literal, offset + 1, nil
		}
	}

	return nil, offset, errors.New("no literal found")
}

func validateLiteral(literal *ast.Literal) error {
	switch literal.Kind {
	case "char":
		if len(literal.Value) == 0 {
			return errors.New("character literal cannot be empty")
		}
	}

	return nil
}
