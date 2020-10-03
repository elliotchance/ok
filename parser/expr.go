package parser

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/lexer"
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

		if consumed%2 == 1 {
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
				lexer.TokenRemainderAssign:

				parts = append(parts, tok)
				offset++
				continue

			default:
				break
			}
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
			continue
		}

		// Try to consume a literal.
		var literal *ast.Literal
		literal, offset, err = consumeLiteral(parser, offset)
		if err == nil {
			parts = append(parts, literal)
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
			continue
		}

		// Map
		var m *ast.Map
		m, offset, err = consumeMap(parser, offset)
		if err == nil {
			parts = append(parts, m)
			continue
		}

		// Type cast
		var call *ast.Call
		call, offset, err = consumeTypeCast(parser, offset)
		if err == nil {
			parts = append(parts, call)
			continue
		}

		// Grouping "()"
		var group *ast.Group
		group, offset, err = consumeGroup(parser, offset)
		if err == nil {
			parts = append(parts, group)
			continue
		}

		// A chained expression. It's important this happens last if all else
		// fails.
		var expr ast.Node
		expr, offset, err = consumeChainedExpr(parser, offset)
		if err == nil {
			parts = append(parts, expr)
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

func consumeChainedExpr(parser *Parser, offset int) (ast.Node, int, error) {
	originalOffset := offset
	var expr ast.Node
	var err error

	// Optional unary expression.
	var unary lexer.Token
	unary, offset, _ = consumeOneOf(parser, offset, []string{
		lexer.TokenNot,
		lexer.TokenMinus,
		lexer.TokenIncrement,
		lexer.TokenDecrement,
	})

	expr, offset, err = consumeIdentifier(parser, offset)
	if err != nil {
		return nil, originalOffset, err
	}

	for offset < len(parser.tokens) {
		tok := parser.tokens[offset]

		if tok.Kind == lexer.TokenDot {
			offset++ // skip "."

			var ident *ast.Identifier
			ident, offset, err = consumeIdentifier(parser, offset)
			if err != nil {
				return nil, originalOffset, err
			}

			expr = &ast.Key{
				Expr: expr,
				Key:  ident,
				Pos:  ident.Pos,
			}
			continue
		}

		if tok.Kind == lexer.TokenSquareOpen {
			offset++ // skip "["

			var key ast.Node
			key, offset, err = consumeExpr(parser, offset, unlimitedTokens)
			if err != nil {
				return nil, originalOffset, err
			}

			offset, err = consume(parser, offset, []string{lexer.TokenSquareClose})
			if err != nil {
				return nil, originalOffset, err
			}

			expr = &ast.Key{
				Expr: expr,
				Key:  key,
				Pos:  key.Position(),
			}
			continue
		}

		if tok.Kind == lexer.TokenParenOpen {
			var call *ast.Call
			call, offset, err = consumeCall(parser, offset)
			if err != nil {
				return nil, originalOffset, err
			}

			call.Expr = expr
			expr = call
			continue
		}

		break
	}

	// If there was a unary operator at the beginning, it must be applied to the
	// final result because "-foo()" = "-(foo())" and not "(-foo)()".
	if unary.Kind != "" {
		expr = &ast.Unary{
			Op:   unary.Value,
			Expr: expr,
			Pos:  expr.Position(),
		}
	}

	return expr, offset, nil
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
