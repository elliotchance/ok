package parser

import (
	"errors"
	"strings"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/lexer"
)

func consumeArray(parser *Parser, offset int) (*ast.Array, int, error) {
	originalOffset := offset
	var err error
	node := &ast.Array{
		Pos: parser.File.Pos(offset),
	}

	var ty string
	ty, offset, err = consumeType(parser, offset)
	if err == nil {
		node.Kind = ty
	}

	// The type must be an array or the explicit "any" type. Otherwise "foo[0]"
	// would be interpreted as an array of type foo with one element instead of
	// an element accessor.
	if ty != "" && ty != "any" && !strings.HasPrefix(ty, "[]") {
		return nil, originalOffset, errors.New("invalid type for array")
	}

	offset, err = consume(parser.File, offset, []string{lexer.TokenSquareOpen})
	if err != nil {
		return nil, originalOffset, err
	}

	// Detect zero elements because consumeExprs will require at least one
	// expression.
	if parser.File.Tokens[offset].Kind == lexer.TokenSquareClose {
		return node, offset + 1, nil
	}

	node.Elements, offset, err = consumeExprs(parser, offset)
	if err != nil {
		return nil, originalOffset, err
	}

	offset, err = consume(parser.File, offset, []string{lexer.TokenSquareClose})
	if err != nil {
		return nil, originalOffset, err
	}

	return node, offset, nil
}
