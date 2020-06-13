package parser

import (
	"ok/ast"
	"ok/lexer"
)

func consumeFunc(parser *Parser, offset int) (*ast.Func, int, bool, error) {
	var err error
	originalOffset := offset

	offset, err = consume(parser.File, offset, []string{
		lexer.TokenFunc, lexer.TokenIdentifier,
		lexer.TokenParenOpen, lexer.TokenParenClose,
		lexer.TokenCurlyOpen,
	})
	if err != nil {
		return nil, offset, originalOffset == offset, err
	}

	fn := &ast.Func{
		Name: parser.File.Tokens[offset-4].Value,
	}

	for {
		var statement ast.Node
		statement, offset, err = consumeStatement(parser, offset)
		if err != nil {
			break
		}

		fn.Statements = append(fn.Statements, statement)
	}

	offset, err = consume(parser.File, offset, []string{lexer.TokenCurlyClose})
	if err != nil {
		return nil, offset, false, err
	}

	return fn, offset, false, nil
}
