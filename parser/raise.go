package parser

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/lexer"
)

func consumeRaise(parser *Parser, offset int) (*ast.Raise, int, error) {
	var err error
	originalOffset := offset

	offset, err = consume(parser.File, offset, []string{lexer.TokenRaise})
	if err != nil {
		return nil, offset, err
	}

	node := &ast.Raise{
		Pos: parser.File.Pos(originalOffset),
	}

	node.Err, offset, err = consumeCall(parser, offset)
	if err != nil {
		return nil, offset, err
	}

	return node, offset, nil
}
