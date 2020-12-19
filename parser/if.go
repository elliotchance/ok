package parser

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/lexer"
	"github.com/elliotchance/ok/types"
)

func consumeIf(parser *Parser, offset int) (*ast.If, int, error) {
	var err error
	originalOffset := offset

	offset, err = consume(parser, offset, []string{lexer.TokenIf})
	if err != nil {
		return nil, offset, err
	}

	node := &ast.If{
		Pos: parser.pos(originalOffset),
	}

	// Only if is allowed to contain the special binary condition using "is".
	//
	// TODO(elliot): In this future this could go anywhere a boolean is
	//  accepted. However, right now its more difficult because the compiler
	//  will treat the variable as that type within the True scope.
	node.Condition, offset, err = consumeIs(parser, offset)
	if err != nil {
		node.Condition, offset, err = consumeExpr(parser, offset, unlimitedTokens)
		if err != nil {
			return nil, offset, err
		}
	}

	node.True, offset, err = consumeBlock(parser, offset)
	if err != nil {
		return nil, offset, err
	}

	// Else is optional.
	offset, err = consume(parser, offset, []string{lexer.TokenElse})
	if err != nil {
		return node, offset, nil // still success
	}

	node.False, offset, err = consumeBlock(parser, offset)
	if err != nil {
		return nil, offset, err
	}

	return node, offset, nil
}

func consumeIs(parser *Parser, offset int) (*ast.Binary, int, error) {
	originalOffset := offset
	var err error
	var left ast.Node
	left, offset, err = consumeIdentifier(parser, offset)
	if err != nil {
		return nil, originalOffset, err
	}

	offset, err = consume(parser, offset, []string{lexer.TokenIs})
	if err != nil {
		return nil, originalOffset, err
	}

	var ty *types.Type
	ty, offset, err = consumeType(parser, offset)
	if err != nil {
		return nil, originalOffset, err
	}

	return &ast.Binary{
		Left:  left,
		Op:    lexer.TokenIs,
		Right: &ast.Identifier{Name: ty.String()},
	}, offset, nil
}
