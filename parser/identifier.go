package parser

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/lexer"
)

func consumeIdentifier(parser *Parser, offset int) (*ast.Identifier, int, error) {
	originalOffset := offset

	var err error
	offset, err = consume(parser, offset, []string{lexer.TokenIdentifier})
	if err != nil {
		return nil, originalOffset, err
	}

	return &ast.Identifier{
		Name: parser.tokens[originalOffset].Value,
		Pos:  parser.pos(originalOffset),
	}, offset, nil
}

// An entity can be an local identifier or imported identifier. Such as "foo" or
// "foo.Bar".
func consumeEntity(parser *Parser, offset int) (*ast.Identifier, int, error) {
	originalOffset := offset

	var err error
	offset, err = consume(parser, offset, []string{lexer.TokenIdentifier})
	if err != nil {
		return nil, originalOffset, err
	}

	identifier := &ast.Identifier{
		Name: parser.tokens[offset-1].Value,
		Pos:  parser.pos(originalOffset),
	}

	offset, err = consume(parser, offset, []string{
		lexer.TokenDot, lexer.TokenIdentifier})
	if err == nil {
		identifier.Name += "." + parser.tokens[offset-1].Value
	}

	return identifier, offset, nil
}
