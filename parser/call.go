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

	var literal *ast.Literal
	literal, offset, err = consumeLiteral(f, offset)
	if err != nil {
		return nil, originalOffset, newTokenMismatch("literal",
			f.Tokens[offset-1].Kind, f.Tokens[offset].Kind)
	}

	err = validateLiteral(literal)
	if err != nil {
		return nil, originalOffset, err
	}

	offset, err = consume(f, offset, []string{lexer.TokenParenClose})
	if err != nil {
		return nil, originalOffset, err
	}

	return &ast.Call{
		FunctionName: f.Tokens[offset-4].Value,
		Arguments: []*ast.Literal{
			literal,
		},
	}, offset, nil
}
