package parser

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/lexer"
)

func consumeArray(parser *Parser, offset int) (*ast.Array, int, error) {
	originalOffset := offset
	var err error
	node := &ast.Array{
		Pos: parser.File.Pos(offset),
	}

	var ty string
	ty, offset, err = consumeType(parser, offset)
	if err == nil {
		node.Kind = ty
	}

	offset, err = consume(parser.File, offset, []string{lexer.TokenSquareOpen})
	if err != nil {
		return nil, originalOffset, err
	}

	// Detect zero elements because consumeExprs will require at least one
	// expression.
	if parser.File.Tokens[offset].Kind == lexer.TokenSquareClose {
		return node, offset + 1, nil
	}

	node.Elements, offset, err = consumeExprs(parser, offset)
	if err != nil {
		return nil, originalOffset, err
	}

	offset, err = consume(parser.File, offset, []string{lexer.TokenSquareClose})
	if err != nil {
		return nil, originalOffset, err
	}

	return node, offset, nil
}
