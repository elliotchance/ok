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
	switch parser.File.Tokens[offset].Kind {
	case lexer.TokenSquareOpen:
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

	case lexer.TokenDot:
		offset++ // skip "."

		var key *ast.Identifier
		key, offset, err = consumeIdentifier(parser, offset)
		if err != nil {
			return nil, originalOffset, err
		}

		return &ast.Key{
			Expr: identifier,

			// Converting the identifier to a literal string means that the
			// compiler need not even know that it is handling an object or map.
			// The only difference is the compiler will need to check the key
			// exists.
			Key: asttest.NewLiteralString(key.Name),

			Pos: parser.File.Pos(originalOffset),
		}, offset, nil
	}

	return identifier, offset, nil
}
