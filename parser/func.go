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
	})
	if err != nil {
		return nil, offset, originalOffset == offset, err
	}

	fn := &ast.Func{
		Name: parser.File.Tokens[offset-3].Value,
	}

	fn.Statements, offset, err = consumeBlock(parser, offset)
	if err != nil {
		return nil, offset, false, err
	}

	return fn, offset, false, nil
}
