package parser

import (
	"ok/ast"
	"ok/lexer"
)

func consumeFor(parser *Parser, offset int) (*ast.For, int, error) {
	var err error

	offset, err = consume(parser.File, offset, []string{lexer.TokenFor})
	if err != nil {
		return nil, offset, err
	}

	node := &ast.For{}

	// Condition is optional.
	if parser.File.Tokens[offset].Kind != lexer.TokenCurlyOpen {
		node.Condition, offset, err = consumeExpr(parser, offset)
		if err != nil {
			return nil, offset, err
		}
	}

	node.Statements, offset, err = consumeBlock(parser, offset)
	if err != nil {
		return nil, offset, err
	}

	return node, offset, nil
}
