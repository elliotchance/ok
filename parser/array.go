package parser

import (
	"ok/ast"
	"ok/lexer"
)

func consumeArray(parser *Parser, offset int) (*ast.Array, int, error) {
	var err error
	node := &ast.Array{}

	var ty string
	ty, offset, err = consumeType(parser, offset)
	if err == nil {
		node.Kind = ty
	}

	offset, err = consume(parser.File, offset, []string{lexer.TokenSquareOpen})
	if err != nil {
		return nil, offset, err
	}

	// Detect zero elements because consumeExprs will require at least one
	// expression.
	if parser.File.Tokens[offset].Kind == lexer.TokenSquareClose {
		return node, offset + 1, nil
	}

	node.Elements, offset, err = consumeExprs(parser, offset)
	if err != nil {
		return nil, offset, err
	}

	offset, err = consume(parser.File, offset, []string{lexer.TokenSquareClose})
	if err != nil {
		return nil, offset, err
	}

	return node, offset, nil
}
