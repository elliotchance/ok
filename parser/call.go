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
	for {
		var expr ast.Node
		expr, offset, err = consumeExpr(parser, offset)
		if err != nil {
			return nil, originalOffset, err
		}

		exprs = append(exprs, expr)

		offset, err = consume(parser.File, offset, []string{lexer.TokenComma})
		if err != nil {
			break
		}
	}

	offset, err = consume(parser.File, offset, []string{lexer.TokenParenClose})
	if err != nil {
		return nil, originalOffset, err
	}

	return &ast.Call{
		FunctionName: parser.File.Tokens[originalOffset].Value,
		Arguments:    exprs,
	}, offset, nil
}
