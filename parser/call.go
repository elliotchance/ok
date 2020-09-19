package parser

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/lexer"
)

func consumeTypeCast(parser *Parser, offset int) (*ast.Call, int, error) {
	originalOffset := offset

	// If we receive a type followed by a value, then we are casting a value.
	// This is picked up as a function call rather than a unary operation,
	// although maybe moving it to unary and refactoring array/map makes more
	// sense?
	var ty lexer.Token
	var err error
	ty, offset, err = consumeOneOf(parser.File, offset, typeTokens)
	if err != nil {
		return nil, originalOffset, err
	}

	var expr ast.Node
	expr, offset, err = consumeExpr(parser, offset, 1)
	if err != nil {
		return nil, originalOffset, err
	}

	call := &ast.Call{
		Expr:      &ast.Identifier{Name: ty.Value},
		Arguments: []ast.Node{expr},
	}

	return call, offset, nil
}

// consumeCall only consumes the arguments, it is expected to be called after an
// expression. That expression will need to be places back into the Call.Expr.
func consumeCall(parser *Parser, offset int) (*ast.Call, int, error) {
	originalOffset := offset
	var err error

	call := &ast.Call{
		Pos: parser.File.Pos(originalOffset),
	}

	offset, err = consume(parser.File, offset, []string{lexer.TokenParenOpen})
	if err != nil {
		return nil, originalOffset, err
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
