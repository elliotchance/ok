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

	// If we receive a type followed by a value, then we are casting a value.
	// This is picked up as a function call rather than a unary operation,
	// although maybe moving it to unary and refactoring array/map makes more
	// sense?
	var ty lexer.Token
	ty, offset, err = consumeOneOf(parser.File, offset, typeTokens)
	if err == nil {
		var expr ast.Node
		expr, offset, err = consumeExpr(parser, offset, 1)
		if err == nil {
			call.FunctionName = ty.Value
			call.Arguments = []ast.Node{expr}

			return call, offset, nil
		}
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
