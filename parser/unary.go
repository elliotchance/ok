package parser

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/lexer"
)

// unary := [ "not" | "-" | "++" | "--" ] expr
func consumeUnary(parser *Parser, offset int) (*ast.Unary, int, error) {
	originalOffset := offset

	var err error
	var t lexer.Token
	t, offset, err = consumeOneOf(parser, offset, []string{
		lexer.TokenNot,
		lexer.TokenMinus,
		lexer.TokenIncrement,
		lexer.TokenDecrement,
	})
	if err != nil {
		return nil, originalOffset, err
	}

	// We cannot simply use consumeExpr here, otherwise expressions that start
	// with a negative such as "-3 == -4" will be interpreted as "-(3 == -4)".
	// There are only a few types that can appear after a unary.
	//
	// TODO(elliot): We should not allow all of these in all cases. For example
	//  ++foo() should not be allowed.

	var expr ast.Node
	expr, offset, err = consumeLiteral(parser, offset)
	if err == nil {
		goto done
	}

	expr, offset, err = consumeCall(parser, offset)
	if err == nil {
		goto done
	}

	expr, offset, err = consumeGroup(parser, offset)
	if err == nil {
		goto done
	}

	expr, offset, err = consumeAssignable(parser, offset)
	if err == nil {
		goto done
	}

done:
	return &ast.Unary{
		Op:   t.Kind,
		Expr: expr,
		Pos:  parser.pos(originalOffset),
	}, offset, nil
}
