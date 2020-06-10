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

	var exprs []ast.Node
	for {
		var expr ast.Node
		expr, offset, err = consumeExpr(f, offset)
		if err != nil {
			return nil, originalOffset, err
		}

		exprs = append(exprs, expr)

		offset, err = consume(f, offset, []string{lexer.TokenComma})
		if err != nil {
			break
		}
	}

	offset, err = consume(f, offset, []string{lexer.TokenParenClose})
	if err != nil {
		return nil, originalOffset, err
	}

	return &ast.Call{
		FunctionName: f.Tokens[originalOffset].Value,
		Arguments:    exprs,
	}, offset, nil
}
