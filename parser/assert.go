package parser

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/lexer"
)

func consumeAssert(parser *Parser, offset int) (*ast.Assert, int, error) {
	originalOffset := offset
	var err error
	var expr ast.Node

	offset, err = consume(parser.File, offset, []string{
		lexer.TokenAssert, lexer.TokenParenOpen})
	if err != nil {
		return nil, originalOffset, err
	}

	expr, offset, err = consumeExpr(parser, offset)
	if err != nil {
		return nil, originalOffset, err
	}

	offset, err = consume(parser.File, offset, []string{lexer.TokenParenClose})
	if err != nil {
		return nil, originalOffset, err
	}

	assert := &ast.Assert{
		Pos: parser.File.Pos(originalOffset),
	}
	if e, ok := expr.(*ast.Binary); ok {
		assert.Expr = e
	} else {
		parser.AppendError(assert,
			"only binary expressions are permitted in assertions")
	}

	return assert, offset, nil
}
