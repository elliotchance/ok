package parser

import (
	"ok/ast"
	"ok/lexer"
)

func consumeCall(tokens []lexer.Token, offset int) (*ast.Call, int, error) {
	var err error
	offset, err = consume(tokens, offset, []string{
		lexer.TokenIdentifier,
		lexer.TokenParenOpen,
		lexer.TokenString,
		lexer.TokenParenClose,
	})
	if err != nil {
		return nil, offset, err
	}

	return &ast.Call{
		FunctionName: tokens[offset-4].Value,
		Arguments:    []string{tokens[offset-2].Value},
	}, offset, nil
}
