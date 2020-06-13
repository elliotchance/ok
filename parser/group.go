package parser

import (
	"ok/ast"
	"ok/lexer"
)

// group := "(" expr ")"
func consumeGroup(parser *Parser, offset int) (*ast.Group, int, error) {
	originalOffset := offset

	var err error
	offset, err = consume(parser.File, offset, []string{lexer.TokenParenOpen})
	if err != nil {
		return nil, originalOffset, err
	}

	var expr ast.Node
	expr, offset, err = consumeExpr(parser, offset)
	if err != nil {
		return nil, originalOffset, err
	}

	offset, err = consume(parser.File, offset, []string{lexer.TokenParenClose})
	if err != nil {
		return nil, originalOffset, err
	}

	return &ast.Group{
		Expr: expr,
	}, offset, nil
}
