package parser

import (
	"ok/ast"
	"ok/lexer"
)

func consumeExpr(parser *Parser, offset int) (ast.Node, int, error) {
	originalOffset := offset

	// Consume as many subexpressions and operators as possible.
	var parts []interface{}
	var err error
	didJustConsumeLiteral := false
	for offset < len(parser.File.Tokens) {
		tok := parser.File.Tokens[offset]

		// Bail out if we reach the end of the line.
		if parser.File.Tokens[offset-1].IsEndOfLine {
			break
		}

		// Try to consume a literal.
		var literal *ast.Literal
		literal, offset, err = consumeLiteral(parser.File, offset)
		if err == nil {
			err = validateLiteral(literal)
			if err != nil {
				// This kind of error should not stop the parsing.
				parser.Errors = append(parser.Errors, err)
			}

			parts = append(parts, literal)
			didJustConsumeLiteral = true
			continue
		}

		// Grouping "()"
		var group *ast.Group
		group, offset, err = consumeGroup(parser, offset)
		if err == nil {
			parts = append(parts, group)
			didJustConsumeLiteral = false
			continue
		}

		// Unary operator (can not be consumed directly after a literal).
		if !didJustConsumeLiteral {
			var unary *ast.Unary
			unary, offset, err = consumeUnary(parser, offset)
			if err == nil {
				parts = append(parts, unary)
				didJustConsumeLiteral = false
				continue
			}
		}

		// An identifier. It's important this happens last if all else fails.
		var identifier *ast.Identifier
		identifier, offset, err = consumeIdentifier(parser, offset)
		if err == nil {
			parts = append(parts, identifier)
			continue
		}

		// Otherwise it must be a a valid binary operator.
		switch tok.Kind {
		case lexer.TokenPlus, lexer.TokenMinus, lexer.TokenTimes,
			lexer.TokenDivide, lexer.TokenRemainder, lexer.TokenAnd,
			lexer.TokenOr, lexer.TokenEqual, lexer.TokenNotEqual,
			lexer.TokenGreaterThan, lexer.TokenGreaterThanEqual,
			lexer.TokenLessThan, lexer.TokenLessThanEqual:
			parts = append(parts, tok)
			offset++
			didJustConsumeLiteral = false
			continue
		}

		break
	}

	if len(parts) == 0 {
		// Nothing was consumed, this is an error.
		return nil, originalOffset, newTokenMismatch("expression",
			parser.File.Tokens[originalOffset-1].Kind,
			parser.File.Tokens[originalOffset].Kind)
	}

	return reduceExpr(parts), offset, nil
}

var operatorPrecedence = map[string]int{
	lexer.TokenOr: 1,

	lexer.TokenAnd: 2,

	lexer.TokenEqual:            3,
	lexer.TokenNotEqual:         3,
	lexer.TokenGreaterThan:      3,
	lexer.TokenGreaterThanEqual: 3,
	lexer.TokenLessThan:         3,
	lexer.TokenLessThanEqual:    3,

	lexer.TokenPlus:  4,
	lexer.TokenMinus: 4,

	lexer.TokenTimes:     5,
	lexer.TokenDivide:    5,
	lexer.TokenRemainder: 5,
}

func reduceExpr(parts []interface{}) ast.Node {
	if len(parts) == 1 {
		return parts[0]
	}

	if len(parts) == 3 {
		return &ast.Binary{
			Left:  parts[0],
			Op:    parts[1].(lexer.Token).Kind,
			Right: parts[2],
		}
	}

	leftPrecedence := operatorPrecedence[parts[1].(lexer.Token).Kind]
	rightPrecedence := operatorPrecedence[parts[len(parts)-2].(lexer.Token).Kind]
	if leftPrecedence > rightPrecedence {
		return &ast.Binary{
			Left:  reduceExpr(parts[:len(parts)-2]),
			Op:    parts[len(parts)-2].(lexer.Token).Kind,
			Right: parts[len(parts)-1].(*ast.Literal),
		}
	}

	return &ast.Binary{
		Left:  parts[0].(*ast.Literal),
		Op:    parts[1].(lexer.Token).Kind,
		Right: reduceExpr(parts[2:]),
	}
}
