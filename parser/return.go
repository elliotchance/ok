package parser

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/lexer"
)

func consumeReturn(parser *Parser, offset int) (*ast.Return, int, error) {
	originalOffset := offset
	var err error
	node := &ast.Return{}

	offset, err = consume(parser.File, offset, []string{lexer.TokenReturn})
	if err != nil {
		return nil, originalOffset, err
	}

	// Detect zero return values because consumeExprs will require at least one
	// expression.
	if parser.File.Tokens[offset-1].IsEndOfLine {
		return node, offset, nil
	}

	node.Exprs, offset, err = consumeExprs(parser, offset)
	if err != nil {
		return nil, originalOffset, err
	}

	return node, offset, nil
}
