package parser

import (
	"fmt"
	"ok/ast"
	"ok/lexer"
)

// statement := [ assign | call | for | break | continue ]
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

	return nil, originalOffset, fmt.Errorf("expecting statement")
}
