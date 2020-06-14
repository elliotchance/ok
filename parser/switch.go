package parser

import (
	"ok/ast"
	"ok/lexer"
)

func consumeCase(parser *Parser, offset int) (*ast.Case, int, error) {
	var err error

	offset, err = consume(parser.File, offset, []string{lexer.TokenCase})
	if err != nil {
		return nil, offset, err
	}

	node := &ast.Case{}

	node.Conditions, offset, err = consumeExprs(parser, offset)
	if err != nil {
		return nil, offset, err
	}

	node.Statements, offset, err = consumeBlock(parser, offset)
	if err != nil {
		return nil, offset, err
	}

	return node, offset, nil
}

func consumeSwitch(parser *Parser, offset int) (*ast.Switch, int, error) {
	var err error

	offset, err = consume(parser.File, offset, []string{lexer.TokenSwitch})
	if err != nil {
		return nil, offset, err
	}

	node := &ast.Switch{}

	// An expression is optional.
	if parser.File.Tokens[offset].Kind != lexer.TokenCurlyOpen {
		node.Expr, offset, err = consumeExpr(parser, offset)
		if err != nil {
			return nil, offset, err
		}
	}

	offset, err = consume(parser.File, offset, []string{lexer.TokenCurlyOpen})
	if err != nil {
		return nil, offset, err
	}

	for {
		// Else is optional, but if we do consume it then we cannot expect any
		// cases after this since it must always appear at the end.
		if parser.File.Tokens[offset].Kind == lexer.TokenElse {
			offset++ // skip else

			node.Else, offset, err = consumeBlock(parser, offset)
			if err != nil {
				return nil, offset, err
			}
		}

		var caseStmt *ast.Case
		caseStmt, offset, err = consumeCase(parser, offset)
		if err != nil {
			break
		}

		node.Cases = append(node.Cases, caseStmt)
	}

	offset, err = consume(parser.File, offset, []string{lexer.TokenCurlyClose})
	if err != nil {
		return nil, offset, err
	}

	return node, offset, nil
}
