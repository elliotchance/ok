package parser

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/lexer"
)

func consumeBlock(parser *Parser, offset int) ([]ast.Node, int, error) {
	var err error
	originalOffset := offset

	offset, err = consume(parser.File, offset, []string{lexer.TokenCurlyOpen})
	if err != nil {
		return nil, offset, err
	}

	var statements []ast.Node
	for {
		if parser.File.Tokens[offset].Kind == lexer.TokenCurlyClose {
			break
		}

		var statement ast.Node
		var hoist bool
		statement, offset, hoist, err = consumeStatement(parser, offset)
		if err != nil {
			parser.AppendErrorAt(parser.File.Pos(offset), err.Error())

			return nil, originalOffset, err
		}

		if hoist {
			statements = append([]ast.Node{statement}, statements...)
		} else {
			statements = append(statements, statement)
		}
	}

	offset, err = consume(parser.File, offset, []string{lexer.TokenCurlyClose})
	if err != nil {
		return nil, offset, err
	}

	return statements, offset, nil
}
