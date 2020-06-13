package parser

import (
	"ok/ast"
	"ok/lexer"
)

// assign := identifier "=" expr
func consumeAssign(parser *Parser, offset int) (*ast.Assign, int, error) {
	originalOffset := offset

	var err error
	offset, err = consume(parser.File, offset, []string{
		lexer.TokenIdentifier,
		lexer.TokenAssign,
	})
	if err != nil {
		return nil, originalOffset, err
	}

	var expr ast.Node
	expr, offset, err = consumeExpr(parser, offset)
	if err != nil {
		return nil, originalOffset, err
	}

	return &ast.Assign{
		VariableName: parser.File.Tokens[originalOffset].Value,
		Expr:         expr,
	}, offset, nil
}
