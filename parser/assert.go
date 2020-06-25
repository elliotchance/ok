package parser

import (
	"errors"
	"ok/ast"
	"ok/lexer"
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

	assert := &ast.Assert{}
	if e, ok := expr.(*ast.Binary); ok {
		assert.Expr = e
	} else {
		err = errors.New(
			"only binary expressions are permitted in assertions")
		parser.Errors = append(parser.Errors, err)
	}

	return assert, offset, nil
}
