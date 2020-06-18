package parser

import (
	"ok/ast"
	"ok/lexer"
)

func consumeCall(parser *Parser, offset int) (*ast.Call, int, error) {
	originalOffset := offset

	var err error
	offset, err = consume(parser.File, offset, []string{
		lexer.TokenIdentifier,
		lexer.TokenParenOpen,
	})
	if err != nil {
		return nil, originalOffset, err
	}

	var exprs []ast.Node
	exprs, offset, err = consumeExprs(parser, offset)

	offset, err = consume(parser.File, offset, []string{lexer.TokenParenClose})
	if err != nil {
		return nil, originalOffset, err
	}

	return &ast.Call{
		FunctionName: parser.File.Tokens[originalOffset].Value,
		Arguments:    exprs,
	}, offset, nil
}
