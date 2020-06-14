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

	node.Condition, offset, err = consumeExpr(parser, offset)
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

	offset, err = consume(parser.File, offset, []string{
		lexer.TokenSwitch, lexer.TokenCurlyOpen,
	})
	if err != nil {
		return nil, offset, err
	}

	node := &ast.Switch{}

	for {
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

	// Else is optional.
	offset, err = consume(parser.File, offset, []string{lexer.TokenElse})
	if err != nil {
		return node, offset, nil // still success
	}

	node.Else, offset, err = consumeBlock(parser, offset)
	if err != nil {
		return nil, offset, err
	}

	return node, offset, nil
}
