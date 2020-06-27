package parser

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/lexer"
)

func consumeMap(parser *Parser, offset int) (*ast.Map, int, error) {
	originalOffset := offset

	var err error
	node := &ast.Map{}

	var ty string
	ty, offset, err = consumeType(parser, offset)
	if err == nil {
		node.Kind = ty
	}

	offset, err = consume(parser.File, offset, []string{lexer.TokenCurlyOpen})
	if err != nil {
		return nil, originalOffset, err
	}

	// Detect zero elements because consumeExprs will require at least one
	// expression.
	//
	// We are only allowed to consume zero elements if a type was supplied,
	// otherwise an empty map would be confused with a block.
	if parser.File.Tokens[offset].Kind == lexer.TokenCurlyClose && ty != "" {
		return node, offset + 1, nil
	}

	node.Elements, offset, err = consumeKeyValues(parser, offset)
	if err != nil {
		return nil, originalOffset, err
	}

	offset, err = consume(parser.File, offset, []string{lexer.TokenCurlyClose})
	if err != nil {
		return nil, originalOffset, err
	}

	return node, offset, nil
}
