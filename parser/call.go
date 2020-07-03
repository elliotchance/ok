package parser

import (
	"fmt"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/lexer"
)

func consumeCall(parser *Parser, offset int) (*ast.Call, int, error) {
	originalOffset := offset
	var err error

	call := &ast.Call{
		Pos: parser.File.Pos(originalOffset),
	}

	offset, err = consume(parser.File, offset, []string{
		lexer.TokenIdentifier,
		lexer.TokenDot,
		lexer.TokenIdentifier,
		lexer.TokenParenOpen,
	})
	if err != nil {
		offset, err = consume(parser.File, offset, []string{
			lexer.TokenIdentifier,
			lexer.TokenParenOpen,
		})
		if err != nil {
			return nil, originalOffset, err
		}

		call.FunctionName = parser.File.Tokens[offset-2].Value
	} else {
		call.FunctionName = fmt.Sprintf("%s.%s",
			parser.File.Tokens[offset-4].Value,
			parser.File.Tokens[offset-2].Value)
	}

	// Catch zero arguments.
	offset, err = consume(parser.File, offset, []string{lexer.TokenParenClose})
	if err == nil {
		return call, offset, nil
	}

	call.Arguments, offset, err = consumeExprs(parser, offset)
	if err != nil {
		return nil, originalOffset, err
	}

	offset, err = consume(parser.File, offset, []string{lexer.TokenParenClose})
	if err != nil {
		return nil, originalOffset, err
	}

	return call, offset, nil
}
