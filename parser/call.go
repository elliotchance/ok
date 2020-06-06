package parser

import (
	"ok/ast"
	"ok/lexer"
)

func consumeCall(f *File, offset int) (*ast.Call, int, error) {
	var err error
	offset, err = consume(f, offset, []string{
		lexer.TokenIdentifier,
		lexer.TokenParenOpen,
		lexer.TokenString,
		lexer.TokenParenClose,
	})
	if err != nil {
		return nil, offset, err
	}

	return &ast.Call{
		FunctionName: f.Tokens[offset-4].Value,
		Arguments:    []string{f.Tokens[offset-2].Value},
	}, offset, nil
}
