package parser

import (
	"fmt"
	"ok/ast"
	"ok/lexer"
)

func consumeStatement(parser *Parser, offset int) (ast.Node, int, error) {
	originalOffset := offset
	var err error

	offset, err = consume(parser.File, offset, []string{lexer.TokenBreak})
	if err == nil {
		return &ast.Break{}, offset, nil
	}

	offset, err = consume(parser.File, offset, []string{lexer.TokenContinue})
	if err == nil {
		return &ast.Continue{}, offset, nil
	}

	var call *ast.Call
	call, offset, err = consumeCall(parser, offset)
	if err == nil {
		return call, offset, nil
	}

	var forStmt *ast.For
	forStmt, offset, err = consumeFor(parser, offset)
	if err == nil {
		return forStmt, offset, nil
	}

	var ifStmt *ast.If
	ifStmt, offset, err = consumeIf(parser, offset)
	if err == nil {
		return ifStmt, offset, nil
	}

	var switchStmt *ast.Switch
	switchStmt, offset, err = consumeSwitch(parser, offset)
	if err == nil {
		return switchStmt, offset, nil
	}

	var unary *ast.Unary
	unary, offset, err = consumeUnary(parser, offset)
	if err == nil {
		return unary, offset, nil
	}

	var expr ast.Node
	expr, offset, err = consumeExpr(parser, offset)
	if err == nil {
		return expr, offset, nil
	}

	return nil, originalOffset, fmt.Errorf("expecting statement")
}
