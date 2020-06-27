package parser

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/lexer"
)

func consumeIf(parser *Parser, offset int) (*ast.If, int, error) {
	var err error

	offset, err = consume(parser.File, offset, []string{lexer.TokenIf})
	if err != nil {
		return nil, offset, err
	}

	node := &ast.If{}

	node.Condition, offset, err = consumeExpr(parser, offset)
	if err != nil {
		return nil, offset, err
	}

	node.True, offset, err = consumeBlock(parser, offset)
	if err != nil {
		return nil, offset, err
	}

	// Else is optional.
	offset, err = consume(parser.File, offset, []string{lexer.TokenElse})
	if err != nil {
		return node, offset, nil // still success
	}

	node.False, offset, err = consumeBlock(parser, offset)
	if err != nil {
		return nil, offset, err
	}

	return node, offset, nil
}
