package parser

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/lexer"
	"github.com/elliotchance/ok/types"
)

func consumeErrorScope(parser *Parser, offset int) (*ast.ErrorScope, int, error) {
	var err error
	originalOffset := offset

	offset, err = consume(parser.File, offset, []string{lexer.TokenTry})
	if err != nil {
		return nil, offset, err
	}

	node := &ast.ErrorScope{
		Pos: parser.File.Pos(originalOffset),
	}

	node.Statements, offset, err = consumeBlock(parser, offset)
	if err != nil {
		return nil, offset, err
	}

	// We can consume zero or more on clauses.
	for parser.File.Tokens[offset].Kind == lexer.TokenOn {
		var on *ast.On
		on, offset, err = consumeOn(parser, offset)
		if err != nil {
			return nil, originalOffset, err
		}

		node.On = append(node.On, on)
	}

	// The finally clause is optional
	if parser.File.Tokens[offset].Kind == lexer.TokenFinally {
		node.Finally, offset, err = consumeFinally(parser, offset)
		if err != nil {
			return nil, originalOffset, err
		}

		parser.AppendFinally(node.Finally)
	}

	return node, offset, nil
}

func consumeOn(parser *Parser, offset int) (*ast.On, int, error) {
	var err error
	originalOffset := offset

	offset, err = consume(parser.File, offset, []string{lexer.TokenOn})
	if err != nil {
		return nil, originalOffset, err
	}

	var ident *ast.Identifier
	ident, offset, err = consumeEntity(parser, offset)
	if err != nil {
		return nil, originalOffset, err
	}

	node := &ast.On{
		Type: types.NewUnresolvedInterface(ident.Name),
		Pos:  parser.File.Pos(originalOffset),
	}

	node.Statements, offset, err = consumeBlock(parser, offset)
	if err != nil {
		return nil, offset, err
	}

	return node, offset, nil
}

func consumeFinally(parser *Parser, offset int) (*ast.Finally, int, error) {
	var err error
	originalOffset := offset

	offset, err = consume(parser.File, offset, []string{lexer.TokenFinally})
	if err != nil {
		return nil, originalOffset, err
	}

	node := &ast.Finally{
		Pos: parser.File.Pos(originalOffset),
	}

	node.Statements, offset, err = consumeBlock(parser, offset)
	if err != nil {
		return nil, offset, err
	}

	return node, offset, nil
}
