package parser

import (
	"ok/ast"
	"ok/lexer"
)

func consumeImport(parser *Parser, offset int) (*ast.Import, int, error) {
	originalOffset := offset
	var err error

	offset, err = consume(parser.File, offset, []string{
		lexer.TokenImport, lexer.TokenStringLiteral})
	if err != nil {
		return nil, originalOffset, err
	}

	return &ast.Import{
		PackageName: parser.File.Tokens[offset-1].Value,
	}, offset, nil
}
