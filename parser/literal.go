package parser

import (
	"errors"
	"strings"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/lexer"
)

// constant := identifier "=" literal
func consumeConstant(parser *Parser, offset int) (string, *ast.Literal, int, error) {
	originalOffset := offset

	var name *ast.Identifier
	var err error
	name, offset, err = consumeIdentifier(parser, offset)
	if err != nil {
		return "", nil, originalOffset, err
	}

	offset, err = consume(parser.File, offset, []string{lexer.TokenAssign})
	if err != nil {
		return "", nil, originalOffset, err
	}

	var value *ast.Literal
	value, offset, err = consumeLiteral(parser, offset)
	if err != nil {
		return "", nil, originalOffset, err
	}

	return name.Name, value, offset, nil
}

func consumeLiteral(parser *Parser, offset int) (*ast.Literal, int, error) {
	originalOffset := offset

	types := []string{
		lexer.TokenBoolLiteral,
		lexer.TokenCharLiteral,
		lexer.TokenDataLiteral,
		lexer.TokenNumberLiteral,
		lexer.TokenStringLiteral,
	}

	for _, ty := range types {
		if t := parser.File.Tokens[offset]; t.Kind == ty {
			literal := &ast.Literal{
				Kind:  strings.Split(ty, " ")[0],
				Value: t.Value,
				Pos:   parser.File.Pos(originalOffset),
			}

			err := validateLiteral(literal)
			if err != nil {
				// This kind of error should not stop the parsing.
				parser.AppendError(literal, err.Error())
			}

			return literal, offset + 1, nil
		}
	}

	return nil, originalOffset, errors.New("no literal found")
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
