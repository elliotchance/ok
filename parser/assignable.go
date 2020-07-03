package parser

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/lexer"
)

func consumeAssignable(parser *Parser, offset int) (ast.Node, int, error) {
	originalOffset := offset
	var err error

	var identifier *ast.Identifier
	identifier, offset, err = consumeIdentifier(parser, offset)
	if err != nil {
		return nil, originalOffset, err
	}

	// Read ahead for key expression.
	if parser.File.Tokens[offset].Kind == lexer.TokenSquareOpen {
		offset++ // skip "["

		var key ast.Node
		key, offset, err = consumeExpr(parser, offset)
		if err != nil {
			return nil, originalOffset, err
		}

		offset, err = consume(parser.File, offset, []string{lexer.TokenSquareClose})
		if err != nil {
			return nil, originalOffset, err
		}

		return &ast.Key{
			Expr: identifier,
			Key:  key,
			Pos:  parser.File.Pos(originalOffset),
		}, offset, nil
	}

	return identifier, offset, nil
}
