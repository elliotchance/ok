package parser

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/lexer"
)

// An assignable is a variable, array element, map value or object property.
func consumeAssignable(parser *Parser, offset int) (ast.Node, int, error) {
	originalOffset := offset
	var err error

	var identifier *ast.Identifier
	identifier, offset, err = consumeIdentifier(parser, offset)
	if err != nil {
		return nil, originalOffset, err
	}

	// Read ahead for key expression.
	switch parser.tokens[offset].Kind {
	case lexer.TokenSquareOpen:
		offset++ // skip "["

		var key ast.Node
		key, offset, err = consumeExpr(parser, offset, unlimitedTokens)
		if err != nil {
			return nil, originalOffset, err
		}

		offset, err = consume(parser, offset, []string{lexer.TokenSquareClose})
		if err != nil {
			return nil, originalOffset, err
		}

		return &ast.Key{
			Expr: identifier,
			Key:  key,
			Pos:  parser.pos(originalOffset),
		}, offset, nil

	case lexer.TokenDot:
		offset++ // skip "."

		var key *ast.Identifier
		key, offset, err = consumeIdentifier(parser, offset)
		if err != nil {
			return nil, originalOffset, err
		}

		return &ast.Key{
			Expr: identifier,
			Pos:  parser.pos(originalOffset),

			// Even though this is an identifier it's not actually a variable
			// and needs to be translated into a literal.
			Key: asttest.NewLiteralString(key.Name),
		}, offset, nil
	}

	return identifier, offset, nil
}
