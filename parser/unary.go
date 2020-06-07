package parser

import (
	"ok/ast"
	"ok/lexer"
)

// group := [ "not" | "-" ] expr
func consumeUnary(f *File, offset int) (*ast.Unary, int, error) {
	originalOffset := offset

	var err error
	var t lexer.Token
	t, offset, err = consumeOneOf(f, offset, []string{
		lexer.TokenNot,
		lexer.TokenMinus,
	})
	if err != nil {
		return nil, originalOffset, err
	}

	var expr ast.Node
	expr, offset, err = consumeExpr(f, offset)
	if err != nil {
		return nil, originalOffset, err
	}

	return &ast.Unary{
		Op:   t.Kind,
		Expr: expr,
	}, offset, nil
}
