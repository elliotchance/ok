package parser

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/lexer"
)

func consumeAssign(parser *Parser, offset int) (*ast.Assign, int, error) {
	var assignable ast.Node
	var err error
	originalOffset := offset
	assign := &ast.Assign{}

	// There must always be at least one assignable on the left side.
	assignable, offset, err = consumeAssignable(parser, offset)
	if err != nil {
		return nil, originalOffset, err
	}
	assign.Lefts = []ast.Node{assignable}

	// Any more are optional.
	for parser.File.Tokens[offset].Kind == lexer.TokenComma {
		offset++ // skip ","

		assignable, offset, err = consumeAssignable(parser, offset)
		if err != nil {
			return nil, originalOffset, err
		}

		assign.Lefts = append(assign.Lefts, assignable)
	}

	// Need an equals followed by some another expression list. We rely on the
	// compiler to verify if the number of expressions on the right side is
	// acceptable.
	offset, err = consume(parser.File, offset, []string{lexer.TokenAssign})
	if err != nil {
		return nil, originalOffset, err
	}

	assign.Rights, offset, err = consumeExprs(parser, offset)
	if err != nil {
		return nil, originalOffset, err
	}

	return assign, offset, nil
}
