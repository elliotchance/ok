package parser

import (
	"errors"
	"fmt"
	"strings"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/lexer"
	"github.com/elliotchance/ok/types"
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

	offset, err = consume(parser, offset, []string{lexer.TokenAssign})
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

	// Optional unary expression.
	var unary lexer.Token
	unary, offset, _ = consumeOneOf(parser, offset, []string{
		lexer.TokenNot,
		lexer.TokenMinus,
	})

	typeLiteralTokens := []string{
		lexer.TokenBoolLiteral,
		lexer.TokenCharLiteral,
		lexer.TokenDataLiteral,
		lexer.TokenNumberLiteral,
		lexer.TokenStringLiteral,
	}

	for _, ty := range typeLiteralTokens {
		if t := parser.tokens[offset]; t.Kind == ty {
			literal := &ast.Literal{
				Kind:  types.TypeFromString(strings.Split(ty, " ")[0]),
				Value: t.Value,
				Pos:   parser.pos(originalOffset),
			}

			err := validateLiteral(literal)
			if err != nil {
				// This kind of error should not stop the parsing.
				parser.appendError(literal, err.Error())
			}

			// Some literals can be reduced now. Other unary combinations will
			// result in an error.
			switch {
			case literal.Kind.Kind == types.KindNumber &&
				unary.Kind == lexer.TokenMinus:
				literal.Value = "-" + literal.Value

			case literal.Kind.Kind == types.KindBool &&
				unary.Kind == lexer.TokenNot:
				if literal.Value == "true" {
					literal.Value = "false"
				} else {
					literal.Value = "true"
				}

			case unary.Kind != "":
				return nil, originalOffset, fmt.Errorf(
					"%s bad unary operation: %s %s",
					literal.Position(),
					unary.Kind, literal.Kind.String())
			}

			return literal, offset + 1, nil
		}
	}

	if unary.Kind != "" {

	}

	return nil, originalOffset, errors.New("no literal found")
}

func validateLiteral(literal *ast.Literal) error {
	switch literal.Kind.Kind {
	case types.KindChar:
		if len(literal.Value) == 0 {
			return errors.New("character literal cannot be empty")
		}
	}

	return nil
}
