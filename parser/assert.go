package parser

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/lexer"
)

func consumeAssert(parser *Parser, offset int) (*ast.Assert, int, error) {
	originalOffset := offset
	var err error
	var expr ast.Node

	offset, err = consume(parser, offset, []string{
		lexer.TokenAssert, lexer.TokenParenOpen})
	if err != nil {
		return nil, originalOffset, err
	}

	expr, offset, err = consumeExpr(parser, offset, unlimitedTokens)
	if err != nil {
		return nil, originalOffset, err
	}

	offset, err = consume(parser, offset, []string{lexer.TokenParenClose})
	if err != nil {
		return nil, originalOffset, err
	}

	assert := &ast.Assert{
		Pos: parser.pos(originalOffset),
	}
	if e, ok := expr.(*ast.Binary); ok {
		assert.Expr = e
	} else {
		parser.appendError(assert,
			"only binary expressions are permitted in assertions")
	}

	return assert, offset, nil
}
