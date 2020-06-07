package parser

import (
	"ok/ast"
	"ok/lexer"
)

func consumeExpr(f *File, offset int) (ast.Node, int, error) {
	originalOffset := offset

	// Consume as many subexpressions and operators as possible.
	var parts []interface{}
	var err error
	for offset < len(f.Tokens) {
		// Try to consume a literal first. This is important because literals
		// may be multi token and contain an operator.
		var literal *ast.Literal
		literal, offset, err = consumeLiteral(f, offset)
		if err == nil {
			err = validateLiteral(literal)
			if err != nil {
				return nil, originalOffset, err
			}

			parts = append(parts, literal)
			continue
		}

		// Otherwise it must be a a valid binary operator.
		switch f.Tokens[offset].Kind {
		case lexer.TokenPlus, lexer.TokenMinus, lexer.TokenTimes,
			lexer.TokenDivide, lexer.TokenRemainder:
			parts = append(parts, f.Tokens[offset])
			offset++
			continue
		}

		break
	}

	if len(parts) == 0 {
		// Nothing was consumed, this is an error.
		return nil, originalOffset, newTokenMismatch("expression",
			f.Tokens[originalOffset-1].Kind, f.Tokens[originalOffset].Kind)
	}

	return reduceExpr(parts), offset, nil
}

var operatorPrecedence = map[string]int{
	lexer.TokenPlus:  1,
	lexer.TokenMinus: 1,

	lexer.TokenTimes:     2,
	lexer.TokenDivide:    2,
	lexer.TokenRemainder: 2,
}

func reduceExpr(parts []interface{}) ast.Node {
	if len(parts) == 1 {
		// TODO(elliot): Check a non-literal.
		return parts[0].(*ast.Literal)
	}

	if len(parts) == 2 {
		return &ast.Unary{
			Op:   lexer.TokenMinus,
			Expr: parts[1].(*ast.Literal),
		}
	}

	if len(parts) == 3 {
		return &ast.Binary{
			Left:  parts[0].(*ast.Literal),
			Op:    parts[1].(lexer.Token).Kind,
			Right: parts[2].(*ast.Literal),
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
