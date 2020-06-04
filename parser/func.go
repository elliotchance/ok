package parser

import (
	"ok/ast"
	"ok/lexer"
)

func consumeFunc(tokens []lexer.Token, offset int) (*ast.Func, int, error) {
	var err error
	offset, err = consume(tokens, offset, []string{
		lexer.TokenFunc, lexer.TokenIdentifier,
		lexer.TokenParenOpen, lexer.TokenParenClose,
		lexer.TokenCurlyOpen,
	})
	if err != nil {
		return nil, offset, err
	}

	fn := &ast.Func{
		Name: tokens[offset-4].Value,
	}

	for {
		var call *ast.Call
		call, offset, _ = consumeCall(tokens, offset)
		if call == nil {
			break
		}

		fn.Statements = append(fn.Statements, call)
	}

	offset, err = consume(tokens, offset, []string{lexer.TokenCurlyClose})
	if err != nil {
		return nil, offset, err
	}

	return fn, offset, nil
}
