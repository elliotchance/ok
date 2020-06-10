package parser

import (
	"fmt"
	"ok/ast"
)

// statement := [ assign | call ]
func consumeStatement(parser *Parser, offset int) (ast.Node, int, error) {
	originalOffset := offset
	var err error

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

	return nil, originalOffset, fmt.Errorf("expecting statement")
}
