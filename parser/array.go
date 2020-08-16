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

	// If "any" is placed in front of the array we need to remove it. Apart from
	// the fact it's redundant (since everything is an "any"). It also isn't a
	// valid runtime type, we need to resolve the real type of the array below.
	if node.Kind == "any" {
		node.Kind = ""
	}

	// The type must be an array or the explicit "[]any" type. Otherwise "foo[0]"
	// would be interpreted as an array of type foo with one element instead of
	// an element accessor.
	if node.Kind != "" && !strings.HasPrefix(ty, "[]") {
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
