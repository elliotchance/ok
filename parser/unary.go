package parser

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/lexer"
)

// unary := [ "not" | "-" | "++" | "--" ] expr
func consumeUnary(parser *Parser, offset int) (*ast.Unary, int, error) {
	originalOffset := offset

	var err error
	var t lexer.Token
	t, offset, err = consumeOneOf(parser.File, offset, []string{
		lexer.TokenNot,
		lexer.TokenMinus,
		lexer.TokenIncrement,
		lexer.TokenDecrement,
	})
	if err != nil {
		return nil, originalOffset, err
	}

	var expr ast.Node
	expr, offset, err = consumeExpr(parser, offset)
	if err != nil {
		return nil, originalOffset, err
	}

	return &ast.Unary{
		Op:   t.Kind,
		Expr: expr,
	}, offset, nil
}
