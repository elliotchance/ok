package parser

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/lexer"
)

func consumeBlock(parser *Parser, offset int) ([]ast.Node, int, error) {
	var err error
	originalOffset := offset

	offset, err = consume(parser, offset, []string{lexer.TokenCurlyOpen})
	if err != nil {
		return nil, offset, err
	}

	var statements []ast.Node
	for {
		if parser.tokens[offset].Kind == lexer.TokenCurlyClose {
			break
		}

		var statement ast.Node
		var hoist bool
		statement, offset, hoist, err = consumeStatement(parser, offset)
		if err != nil {
			parser.appendErrorAt(parser.pos(offset), err.Error())

			return nil, originalOffset, err
		}

		if hoist {
			statements = append([]ast.Node{statement}, statements...)
		} else {
			statements = append(statements, statement)
		}
	}

	offset, err = consume(parser, offset, []string{lexer.TokenCurlyClose})
	if err != nil {
		return nil, offset, err
	}

	return statements, offset, nil
}
