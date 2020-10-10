package parser

import (
	"fmt"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/lexer"
)

func consumeStatement(parser *Parser, offset int) (_ ast.Node, _ int, hoist bool, finalErr error) {
	originalOffset := offset
	var err error

	offset, err = consume(parser, offset, []string{lexer.TokenBreak})
	if err == nil {
		return &ast.Break{
			Pos: parser.pos(originalOffset),
		}, offset, hoist, nil
	}

	offset, err = consume(parser, offset, []string{lexer.TokenContinue})
	if err == nil {
		return &ast.Continue{
			Pos: parser.pos(originalOffset),
		}, offset, hoist, nil
	}

	var fn *ast.Func
	fn, offset, _, err = consumeFunc(parser, offset)
	if err == nil {
		assign := &ast.Assign{
			Lefts: []ast.Node{
				&ast.Identifier{Name: fn.Name},
			},
			Rights: []ast.Node{
				fn,
			},
		}

		// It's not enough to simply translate a nested function into a variable
		// assignment because we have to to ensure that the assignment happens
		// before any other code in this scope runs, so it's possible to call a
		// function that might be defined below the invocation.
		//
		// Setting hoist to true will move this instruction to the top of the
		// statements instead of appending on the end.
		hoist = true

		return assign, offset, hoist, nil
	}

	var assertRaise *ast.AssertRaise
	assertRaise, offset, err = consumeAssertRaise(parser, offset)
	if err == nil {
		return assertRaise, offset, hoist, nil
	}

	var assert *ast.Assert
	assert, offset, err = consumeAssert(parser, offset)
	if err == nil {
		return assert, offset, hoist, nil
	}

	var rtn *ast.Return
	rtn, offset, err = consumeReturn(parser, offset)
	if err == nil {
		return rtn, offset, hoist, nil
	}

	var raise *ast.Raise
	raise, offset, err = consumeRaise(parser, offset)
	if err == nil {
		return raise, offset, hoist, nil
	}

	var assign *ast.Assign
	assign, offset, err = consumeAssign(parser, offset)
	if err == nil {
		return assign, offset, hoist, nil
	}

	var errorScope *ast.ErrorScope
	errorScope, offset, err = consumeErrorScope(parser, offset)
	if err == nil {
		return errorScope, offset, hoist, nil
	}

	var forStmt *ast.For
	forStmt, offset, err = consumeFor(parser, offset)
	if err == nil {
		return forStmt, offset, hoist, nil
	}

	var ifStmt *ast.If
	ifStmt, offset, err = consumeIf(parser, offset)
	if err == nil {
		return ifStmt, offset, hoist, nil
	}

	var switchStmt *ast.Switch
	switchStmt, offset, err = consumeSwitch(parser, offset)
	if err == nil {
		return switchStmt, offset, hoist, nil
	}

	var expr ast.Node
	expr, offset, err = consumeExpr(parser, offset, unlimitedTokens)
	if err == nil {
		return expr, offset, hoist, nil
	}

	return nil, originalOffset, hoist, fmt.Errorf("expecting statement")
}
