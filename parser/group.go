package parser

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/lexer"
)

// group := "(" expr ")"
func consumeGroup(parser *Parser, offset int) (*ast.Group, int, error) {
	originalOffset := offset

	var err error
	offset, err = consume(parser, offset, []string{lexer.TokenParenOpen})
	if err != nil {
		return nil, originalOffset, err
	}

	var expr ast.Node
	expr, offset, err = consumeExpr(parser, offset, unlimitedTokens)
	if err != nil {
		return nil, originalOffset, err
	}

	offset, err = consume(parser, offset, []string{lexer.TokenParenClose})
	if err != nil {
		return nil, originalOffset, err
	}

	return &ast.Group{
		Expr: expr,
		Pos:  parser.pos(originalOffset),
	}, offset, nil
}
