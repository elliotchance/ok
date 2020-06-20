package parser

import (
	"ok/ast"
	"ok/lexer"
)

func consumeFor(parser *Parser, offset int) (*ast.For, int, error) {
	var err error

	offset, err = consume(parser.File, offset, []string{lexer.TokenFor})
	if err != nil {
		return nil, offset, err
	}

	node := &ast.For{}

	// Condition is optional.
	if parser.File.Tokens[offset].Kind != lexer.TokenCurlyOpen {
		// Condition may be an "in" expression.
		node.Condition, offset, err = consumeIn(parser, offset)
		if err != nil {
			// So it must be an expression.
			node.Condition, offset, err = consumeExpr(parser, offset)
			if err != nil {
				return nil, offset, err
			}
		}
	}

	// If we encounter a ';' it means the Condition was actually an Init and we
	// should expected to see two more nodes.
	offset, err = consume(parser.File, offset, []string{lexer.TokenSemiColon})
	if err == nil {
		node.Init = node.Condition

		// Condition may be an "in" expression.
		node.Condition, offset, err = consumeIn(parser, offset)
		if err != nil {
			// So it must be an expression.
			node.Condition, offset, err = consumeExpr(parser, offset)
			if err != nil {
				return nil, offset, err
			}
		}

		offset, err = consume(parser.File, offset, []string{lexer.TokenSemiColon})
		if err != nil {
			return nil, offset, err
		}

		node.Next, offset, err = consumeExpr(parser, offset)
		if err != nil {
			return nil, offset, err
		}
	}

	node.Statements, offset, err = consumeBlock(parser, offset)
	if err != nil {
		return nil, offset, err
	}

	return node, offset, nil
}

func consumeIn(parser *Parser, offset int) (*ast.In, int, error) {
	originalOffset := offset
	var err error
	node := &ast.In{}

	offset, err = consume(parser.File, offset, []string{lexer.TokenIdentifier})
	if err != nil {
		return nil, originalOffset, err
	}
	node.Key = parser.File.Tokens[originalOffset].Value

	// Value is optional.
	offset, err = consume(parser.File, offset, []string{
		lexer.TokenComma, lexer.TokenIdentifier})
	if err == nil {
		node.Value = parser.File.Tokens[offset-1].Value
	}

	offset, err = consume(parser.File, offset, []string{lexer.TokenIn})
	if err != nil {
		return nil, originalOffset, err
	}

	node.Expr, offset, err = consumeExpr(parser, offset)
	if err != nil {
		return nil, originalOffset, err
	}

	return node, offset, nil
}
