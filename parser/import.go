package parser

import (
	"strings"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/lexer"
)

func consumeImport(parser *Parser, offset int) (*ast.Import, int, error) {
	originalOffset := offset
	var err error

	offset, err = consume(parser.File, offset, []string{
		lexer.TokenImport, lexer.TokenStringLiteral})
	if err != nil {
		return nil, originalOffset, err
	}

	packageName := parser.File.Tokens[offset-1].Value
	parts := strings.Split(packageName, "/")

	return &ast.Import{
		VariableName: parts[len(parts)-1],
		PackageName:  packageName,
		Pos:          parser.File.Pos(originalOffset),
	}, offset, nil
}
