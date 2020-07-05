package parser

import (
	"fmt"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/lexer"
)

func consumeStatement(parser *Parser, offset int) (_ ast.Node, _ int, finalErr error) {
	originalOffset := offset
	var err error

	offset, err = consume(parser.File, offset, []string{lexer.TokenBreak})
	if err == nil {
		return &ast.Break{
			Pos: parser.File.Pos(originalOffset),
		}, offset, nil
	}

	offset, err = consume(parser.File, offset, []string{lexer.TokenContinue})
	if err == nil {
		return &ast.Continue{
			Pos: parser.File.Pos(originalOffset),
		}, offset, nil
	}

	var assert *ast.Assert
	assert, offset, err = consumeAssert(parser, offset)
	if err == nil {
		return assert, offset, nil
	}

	var rtn *ast.Return
	rtn, offset, err = consumeReturn(parser, offset)
	if err == nil {
		return rtn, offset, nil
	}

	var assign *ast.Assign
	assign, offset, err = consumeAssign(parser, offset)
	if err == nil {
		return assign, offset, nil
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

	var expr ast.Node
	expr, offset, err = consumeExpr(parser, offset, unlimitedTokens)
	if err == nil {
		return expr, offset, nil
	}

	return nil, originalOffset, fmt.Errorf("expecting statement")
}
