package parser

import (
	"ok/ast"
	"ok/lexer"
)

func consumeFunc(f *File, offset int) (*ast.Func, int, bool, error) {
	var err error
	originalOffset := offset

	offset, err = consume(f, offset, []string{
		lexer.TokenFunc, lexer.TokenIdentifier,
		lexer.TokenParenOpen, lexer.TokenParenClose,
		lexer.TokenCurlyOpen,
	})
	if err != nil {
		return nil, offset, originalOffset == offset, err
	}

	fn := &ast.Func{
		Name: f.Tokens[offset-4].Value,
	}

	for {
		var call *ast.Call
		call, offset, _ = consumeCall(f, offset)
		if call == nil {
			break
		}

		fn.Statements = append(fn.Statements, call)
	}

	offset, err = consume(f, offset, []string{lexer.TokenCurlyClose})
	if err != nil {
		return nil, offset, false, err
	}

	return fn, offset, false, nil
}
