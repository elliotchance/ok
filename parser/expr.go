package parser

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/lexer"
	"github.com/elliotchance/ok/types"
)

// unlimitedTokens is just some very large amount to be used when you do not
// need a limit for consumeExpr.
const unlimitedTokens = 1000000

// maxTokens can be unlimitedTokens.
func consumeExpr(parser *Parser, offset, maxTokens int) (ast.Node, int, error) {
	originalOffset := offset

	// Consume as many subexpressions and operators as possible.
	var parts []interface{}
	var err error
	didJustConsumeLiteral := false
	for consumed := 0; offset < len(parser.tokens) && consumed < maxTokens; consumed++ {
		tok := parser.tokens[offset]

		// Bail out if we reach the end of the line. However, be careful that we
		// don't look back to the previous token that is from a previous expr
		// read (outside the scope of this expression).
		if offset > originalOffset && parser.tokens[offset-1].IsEndOfLine {
			break
		}

		// Some tokens signal the end of the expression.
		if tok.Kind == lexer.TokenColon || tok.Kind == lexer.TokenComma {
			break
		}

		// Try to consume a function literal. Named function would have already
		// been consumed and translated into variable assignments at the
		// statement level. However, this doesn't mean that function literals
		// here are not allowed to have a name.
		var fn *ast.Func
		var fnAnon bool
		fn, offset, fnAnon, err = consumeFunc(parser, offset)
		if err == nil {
			// If the function literal has a name it will confuse the compiler
			// into thinking it's a package-level entity.
			if !fnAnon {
				parser.appendError(fn, "function literals cannot be named")
			}

			parts = append(parts, fn)
			didJustConsumeLiteral = true
			continue
		}

		// Try to consume a literal.
		var literal *ast.Literal
		literal, offset, err = consumeLiteral(parser, offset)
		if err == nil {
			parts = append(parts, literal)
			didJustConsumeLiteral = true
			continue
		}

		// Interpolate
		var interpolate *ast.Interpolate
		interpolate, offset, err = consumeInterpolate(parser, offset)
		if err == nil {
			parts = append(parts, interpolate)
			continue
		}

		// Array
		var array *ast.Array
		array, offset, err = consumeArray(parser, offset)
		if err == nil {
			parts = append(parts, array)
			didJustConsumeLiteral = false
			continue
		}

		// Map
		var m *ast.Map
		m, offset, err = consumeMap(parser, offset)
		if err == nil {
			parts = append(parts, m)
			didJustConsumeLiteral = false
			continue
		}

		// Type cast
		var call *ast.Call
		call, offset, err = consumeTypeCast(parser, offset)
		if err == nil {
			parts = append(parts, call)
			didJustConsumeLiteral = true
			continue
		}

		// Depending on whether the last token was an operator determines if "("
		// indicates a Call or Group.
		if len(parts) > 0 && operatorPrecedence[parser.tokens[offset-1].Kind] == 0 {
			var call *ast.Call
			call, offset, err = consumeCall(parser, offset)
			if err == nil {
				// Transform the last expression into a Call.
				call.Expr = parts[len(parts)-1].(ast.Node)
				parts = append(parts[:len(parts)-1], call)
				didJustConsumeLiteral = true
				continue
			}
		} else {
			// Grouping "()"
			var group *ast.Group
			group, offset, err = consumeGroup(parser, offset)
			if err == nil {
				parts = append(parts, group)
				didJustConsumeLiteral = false
				continue
			}
		}

		// Unary operator (can not be consumed directly after a literal).
		if !didJustConsumeLiteral {
			var unary *ast.Unary
			unary, offset, err = consumeUnary(parser, offset)
			if err == nil {
				// A unary can be a "-" number, this can be reduced now into a
				// number. This saves on some processing buts it's also required
				// for assertions that only work with simple binary operations.
				if lit, ok := unary.Expr.(*ast.Literal); ok &&
					lit.Kind.Kind == types.KindNumber &&
					unary.Op == lexer.TokenMinus {
					newLit := &ast.Literal{
						Kind:  types.Number,
						Value: "-" + lit.Value,
						Pos:   unary.Pos,
					}
					parts = append(parts, newLit)
				} else {
					parts = append(parts, unary)
				}

				didJustConsumeLiteral = false
				continue
			}
		}

		// An assignable. It's important this happens last if all else fails.
		var assignable ast.Node
		assignable, offset, err = consumeAssignable(parser, offset)
		if err == nil {
			parts = append(parts, assignable)
			didJustConsumeLiteral = true
			continue
		}

		// Otherwise it must be a a valid binary operator.
		switch tok.Kind {
		case
			// Arithmetic
			lexer.TokenPlus, lexer.TokenMinus, lexer.TokenTimes,
			lexer.TokenDivide, lexer.TokenRemainder,

			// Logical
			lexer.TokenAnd, lexer.TokenOr,

			// Comparison
			lexer.TokenEqual, lexer.TokenNotEqual,
			lexer.TokenGreaterThan, lexer.TokenGreaterThanEqual,
			lexer.TokenLessThan, lexer.TokenLessThanEqual,

			// Assignment
			lexer.TokenAssign, lexer.TokenPlusAssign, lexer.TokenMinusAssign,
			lexer.TokenTimesAssign, lexer.TokenDivideAssign,
			lexer.TokenRemainderAssign,

			// Access
			lexer.TokenDot:

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
			parser.tokens[originalOffset-1].Kind,
			parser.tokens[originalOffset].Kind)
	}

	return reduceExpr(parts), offset, nil
}

func consumeExprs(parser *Parser, offset int) ([]ast.Node, int, error) {
	// There must always be one expression.
	var expr ast.Node
	var err error
	expr, offset, err = consumeExpr(parser, offset, unlimitedTokens)
	if err != nil {
		return nil, offset, err
	}

	// Any others are optional, but if there is a comma it must be followed by
	// the next expression.
	exprs := []ast.Node{expr}
	for {
		if parser.tokens[offset].Kind != lexer.TokenComma {
			break
		}

		offset++ // skip comma
		expr, offset, err = consumeExpr(parser, offset, unlimitedTokens)
		if err != nil {
			return nil, offset, err
		}

		exprs = append(exprs, expr)
	}

	return exprs, offset, nil
}

var operatorPrecedence = map[string]int{
	lexer.TokenAssign:          1,
	lexer.TokenPlusAssign:      1,
	lexer.TokenMinusAssign:     1,
	lexer.TokenTimesAssign:     1,
	lexer.TokenDivideAssign:    1,
	lexer.TokenRemainderAssign: 1,

	lexer.TokenOr: 2,

	lexer.TokenAnd: 3,

	lexer.TokenEqual:            4,
	lexer.TokenNotEqual:         4,
	lexer.TokenGreaterThan:      4,
	lexer.TokenGreaterThanEqual: 4,
	lexer.TokenLessThan:         4,
	lexer.TokenLessThanEqual:    4,

	lexer.TokenPlus:  5,
	lexer.TokenMinus: 5,

	lexer.TokenTimes:     6,
	lexer.TokenDivide:    6,
	lexer.TokenRemainder: 6,

	lexer.TokenDot: 7,
}

func reduceExpr(parts []interface{}) ast.Node {
	if len(parts) == 1 {
		return parts[0].(ast.Node)
	}

	if len(parts) == 3 {
		return &ast.Binary{
			Left:  parts[0].(ast.Node),
			Op:    parts[1].(lexer.Token).Kind,
			Right: parts[2].(ast.Node),
		}
	}

	// We have to find the lowest precedence token to split on.
	winner := 0
	winnerPrecedence := 10
	for i := 1; i < len(parts); i += 2 {
		// Will always be true the first time.
		p := operatorPrecedence[parts[i].(lexer.Token).Kind]
		if p < winnerPrecedence {
			winner = i
			winnerPrecedence = p
		}
	}

	return &ast.Binary{
		Left:  reduceExpr(parts[:winner]),
		Op:    parts[winner].(lexer.Token).Kind,
		Right: reduceExpr(parts[winner+1:]),
	}
}
