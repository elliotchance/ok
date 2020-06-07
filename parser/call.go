package parser

import (
	"ok/ast"
	"ok/lexer"
)

func consumeCall(f *File, offset int) (*ast.Call, int, error) {
	originalOffset := offset

	var err error
	offset, err = consume(f, offset, []string{
		lexer.TokenIdentifier,
		lexer.TokenParenOpen,
	})
	if err != nil {
		return nil, originalOffset, err
	}

	var expr ast.Node
	expr, offset, err = consumeExpr(f, offset)
	if err != nil {
		return nil, originalOffset, err
	}

	offset, err = consume(f, offset, []string{lexer.TokenParenClose})
	if err != nil {
		return nil, originalOffset, err
	}

	return &ast.Call{
		FunctionName: f.Tokens[originalOffset].Value,
		Arguments:    []ast.Node{expr},
	}, offset, nil
}
