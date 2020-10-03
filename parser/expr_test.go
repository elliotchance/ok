package parser_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/lexer"
	"github.com/elliotchance/ok/parser"
	"github.com/elliotchance/ok/types"
	"github.com/stretchr/testify/assert"
)

func TestExpr(t *testing.T) {
	for testName, test := range map[string]struct {
		str      string
		expected ast.Node
		errs     []error
	}{
		"literal-true": {
			str: `true`,
			expected: &ast.Literal{
				Kind:  types.Bool,
				Value: "true",
			},
		},
		"literal-false": {
			str: `false`,
			expected: &ast.Literal{
				Kind:  types.Bool,
				Value: "false",
			},
		},
		"literal-char": {
			str: `'a'`,
			expected: &ast.Literal{
				Kind:  types.Char,
				Value: "a",
			},
		},
		"literal-zero-length-char": {
			str: `''`,
			expected: &ast.Literal{
				Kind:  types.Char,
				Value: "",
			},
			errs: []error{
				errors.New("a.ok:1:15 character literal cannot be empty"),
			},
		},
		"literal-number-zero": {
			str: `0`,
			expected: &ast.Literal{
				Kind:  types.Number,
				Value: "0",
			},
		},
		"literal-number-negative": {
			str:      `-3.20`,
			expected: asttest.NewLiteralNumber("-3.20"),
		},
		"compare-negative-numbers": {
			str: `-3 == -4`,
			expected: &ast.Binary{
				Left:  asttest.NewLiteralNumber("-3"),
				Op:    lexer.TokenEqual,
				Right: asttest.NewLiteralNumber("-4"),
			},
		},
		"numbers-plus": {
			str: `3 + 2`,
			expected: &ast.Binary{
				Left:  asttest.NewLiteralNumber("3"),
				Op:    lexer.TokenPlus,
				Right: asttest.NewLiteralNumber("2"),
			},
		},
		"numbers-minus": {
			str: `3 - 2`,
			expected: &ast.Binary{
				Left:  asttest.NewLiteralNumber("3"),
				Op:    lexer.TokenMinus,
				Right: asttest.NewLiteralNumber("2"),
			},
		},
		"numbers-times": {
			str: `3.0*2.1`,
			expected: &ast.Binary{
				Left:  asttest.NewLiteralNumber("3.0"),
				Op:    lexer.TokenTimes,
				Right: asttest.NewLiteralNumber("2.1"),
			},
		},
		"numbers-divide": {
			str: `3/2.0`,
			expected: &ast.Binary{
				Left:  asttest.NewLiteralNumber("3"),
				Op:    lexer.TokenDivide,
				Right: asttest.NewLiteralNumber("2.0"),
			},
		},
		"numbers-remainder": {
			str: `3 % 2`,
			expected: &ast.Binary{
				Left:  asttest.NewLiteralNumber("3"),
				Op:    lexer.TokenRemainder,
				Right: asttest.NewLiteralNumber("2"),
			},
		},
		"expr-3-linear-order": {
			str: `1 + 2 - 3`,
			expected: &ast.Binary{
				Left: asttest.NewLiteralNumber("1"),
				Op:   lexer.TokenPlus,
				Right: &ast.Binary{
					Left:  asttest.NewLiteralNumber("2"),
					Op:    lexer.TokenMinus,
					Right: asttest.NewLiteralNumber("3"),
				},
			},
		},
		"expr-3-non-linear-order": {
			str: `1 * 2 - 3`,
			expected: &ast.Binary{
				Left: &ast.Binary{
					Left:  asttest.NewLiteralNumber("1"),
					Op:    lexer.TokenTimes,
					Right: asttest.NewLiteralNumber("2"),
				},
				Op:    lexer.TokenMinus,
				Right: asttest.NewLiteralNumber("3"),
			},
		},
		"expr-3-grouping": {
			str: `1 * (2 - 3)`,
			expected: &ast.Binary{
				Left: asttest.NewLiteralNumber("1"),
				Op:   lexer.TokenTimes,
				Right: &ast.Group{
					Expr: &ast.Binary{
						Left:  asttest.NewLiteralNumber("2"),
						Op:    lexer.TokenMinus,
						Right: asttest.NewLiteralNumber("3"),
					},
				},
			},
		},
		"group-1": {
			str: `(2 - 3)`,
			expected: &ast.Group{
				Expr: &ast.Binary{
					Left:  asttest.NewLiteralNumber("2"),
					Op:    lexer.TokenMinus,
					Right: asttest.NewLiteralNumber("3"),
				},
			},
		},
		"group-2": {
			str: `(2)`,
			expected: &ast.Group{
				Expr: asttest.NewLiteralNumber("2"),
			},
		},
		"bool-and-bool": {
			str: `true and false`,
			expected: &ast.Binary{
				Left:  asttest.NewLiteralBool(true),
				Op:    lexer.TokenAnd,
				Right: asttest.NewLiteralBool(false),
			},
		},
		"bool-or-bool": {
			str: `true or false`,
			expected: &ast.Binary{
				Left:  asttest.NewLiteralBool(true),
				Op:    lexer.TokenOr,
				Right: asttest.NewLiteralBool(false),
			},
		},
		"not-bool": {
			str: `not true`,
			expected: &ast.Unary{
				Op:   lexer.TokenNot,
				Expr: asttest.NewLiteralBool(true),
			},
		},
		"bool-equal-bool": {
			str: `true==false`,
			expected: &ast.Binary{
				Left:  asttest.NewLiteralBool(true),
				Op:    lexer.TokenEqual,
				Right: asttest.NewLiteralBool(false),
			},
		},
		"bool-not-equal-bool": {
			str: `true != false`,
			expected: &ast.Binary{
				Left:  asttest.NewLiteralBool(true),
				Op:    lexer.TokenNotEqual,
				Right: asttest.NewLiteralBool(false),
			},
		},
		"bool-greater-than-bool": {
			str: `true>false`,
			expected: &ast.Binary{
				Left:  asttest.NewLiteralBool(true),
				Op:    lexer.TokenGreaterThan,
				Right: asttest.NewLiteralBool(false),
			},
		},
		"bool-greater-than-equal-bool": {
			str: `true>=false`,
			expected: &ast.Binary{
				Left:  asttest.NewLiteralBool(true),
				Op:    lexer.TokenGreaterThanEqual,
				Right: asttest.NewLiteralBool(false),
			},
		},
		"bool-less-than-bool": {
			str: `true < false`,
			expected: &ast.Binary{
				Left:  asttest.NewLiteralBool(true),
				Op:    lexer.TokenLessThan,
				Right: asttest.NewLiteralBool(false),
			},
		},
		"bool-less-than-equal-bool": {
			str: `true <= false`,
			expected: &ast.Binary{
				Left:  asttest.NewLiteralBool(true),
				Op:    lexer.TokenLessThanEqual,
				Right: asttest.NewLiteralBool(false),
			},
		},
		"increment": {
			str: "++a",
			expected: &ast.Unary{
				Op:   "++",
				Expr: &ast.Identifier{Name: "a"},
			},
		},
		"decrement": {
			str: "--a",
			expected: &ast.Unary{
				Op:   "--",
				Expr: &ast.Identifier{Name: "a"},
			},
		},
		"add-assign": {
			str: "a += 3",
			expected: &ast.Binary{
				Left:  &ast.Identifier{Name: "a"},
				Op:    lexer.TokenPlusAssign,
				Right: asttest.NewLiteralNumber("3"),
			},
		},
		"minus-assign": {
			str: "a -= 3",
			expected: &ast.Binary{
				Left:  &ast.Identifier{Name: "a"},
				Op:    lexer.TokenMinusAssign,
				Right: asttest.NewLiteralNumber("3"),
			},
		},
		"times-assign": {
			str: "a *= 3",
			expected: &ast.Binary{
				Left:  &ast.Identifier{Name: "a"},
				Op:    lexer.TokenTimesAssign,
				Right: asttest.NewLiteralNumber("3"),
			},
		},
		"divide-assign": {
			str: "a /= 3",
			expected: &ast.Binary{
				Left:  &ast.Identifier{Name: "a"},
				Op:    lexer.TokenDivideAssign,
				Right: asttest.NewLiteralNumber("3"),
			},
		},
		"remainder-assign": {
			str: "a %= 3",
			expected: &ast.Binary{
				Left:  &ast.Identifier{Name: "a"},
				Op:    lexer.TokenRemainderAssign,
				Right: asttest.NewLiteralNumber("3"),
			},
		},
		"minus-1": {
			str: "1 - 2 + 3",
			expected: &ast.Binary{
				Left: asttest.NewLiteralNumber("1"),
				Op:   lexer.TokenMinus,
				Right: &ast.Binary{
					Left:  asttest.NewLiteralNumber("2"),
					Op:    lexer.TokenPlus,
					Right: asttest.NewLiteralNumber("3"),
				},
			},
		},
		"minus-2": {
			str: "y - y + y",
			expected: &ast.Binary{
				Left: &ast.Identifier{Name: "y"},
				Op:   lexer.TokenMinus,
				Right: &ast.Binary{
					Left:  &ast.Identifier{Name: "y"},
					Op:    lexer.TokenPlus,
					Right: &ast.Identifier{Name: "y"},
				},
			},
		},
		"reduce-right-binary": {
			str: "i < a - b",
			expected: asttest.NewBinary(
				&ast.Identifier{Name: "i"},
				lexer.TokenLessThan,
				&ast.Binary{
					Left:  &ast.Identifier{Name: "a"},
					Op:    lexer.TokenMinus,
					Right: &ast.Identifier{Name: "b"},
				},
			),
		},
		"reduce-right-binary-2": {
			str: "i < len(a) - len(b)",
			expected: asttest.NewBinary(
				&ast.Identifier{Name: "i"},
				lexer.TokenLessThan,
				&ast.Binary{
					Left: &ast.Call{
						Expr: &ast.Identifier{Name: "len"},
						Arguments: []ast.Node{
							&ast.Identifier{Name: "a"},
						},
					},
					Op: lexer.TokenMinus,
					Right: &ast.Call{
						Expr: &ast.Identifier{Name: "len"},
						Arguments: []ast.Node{
							&ast.Identifier{Name: "b"},
						},
					},
				},
			),
		},
		"reduce-right-binary-3": {
			str: "len(a) - len(b) < i",
			expected: asttest.NewBinary(
				&ast.Binary{
					Left: &ast.Call{
						Expr: &ast.Identifier{Name: "len"},
						Arguments: []ast.Node{
							&ast.Identifier{Name: "a"},
						},
					},
					Op: lexer.TokenMinus,
					Right: &ast.Call{
						Expr: &ast.Identifier{Name: "len"},
						Arguments: []ast.Node{
							&ast.Identifier{Name: "b"},
						},
					},
				},
				lexer.TokenLessThan,
				&ast.Identifier{Name: "i"},
			),
		},
		"and-1": {
			str: "n >= a and n <= b",
			expected: asttest.NewBinary(
				asttest.NewBinary(
					&ast.Identifier{Name: "n"},
					lexer.TokenGreaterThanEqual,
					&ast.Identifier{Name: "a"},
				),
				lexer.TokenAnd,
				asttest.NewBinary(
					&ast.Identifier{Name: "n"},
					lexer.TokenLessThanEqual,
					&ast.Identifier{Name: "b"},
				),
			),
		},
		"and-2": {
			str: "n >= foo('A') and n <= bar('Z')",
			expected: asttest.NewBinary(
				asttest.NewBinary(
					&ast.Identifier{Name: "n"},
					lexer.TokenGreaterThanEqual,
					&ast.Call{
						Expr: &ast.Identifier{Name: "foo"},
						Arguments: []ast.Node{
							asttest.NewLiteralChar('A'),
						},
					},
				),
				lexer.TokenAnd,
				asttest.NewBinary(
					&ast.Identifier{Name: "n"},
					lexer.TokenLessThanEqual,
					&ast.Call{
						Expr: &ast.Identifier{Name: "bar"},
						Arguments: []ast.Node{
							asttest.NewLiteralChar('Z'),
						},
					},
				),
			),
		},
		"and-3": {
			str: "n >= number 'A' and n <= number 'Z'",
			expected: asttest.NewBinary(
				asttest.NewBinary(
					&ast.Identifier{Name: "n"},
					lexer.TokenGreaterThanEqual,
					&ast.Call{
						Expr: &ast.Identifier{Name: "number"},
						Arguments: []ast.Node{
							asttest.NewLiteralChar('A'),
						},
					},
				),
				lexer.TokenAnd,
				asttest.NewBinary(
					&ast.Identifier{Name: "n"},
					lexer.TokenLessThanEqual,
					&ast.Call{
						Expr: &ast.Identifier{Name: "number"},
						Arguments: []ast.Node{
							asttest.NewLiteralChar('Z'),
						},
					},
				),
			),
		},
		"call-1": {
			str: "n >= number 'A'",
			expected: asttest.NewBinary(
				&ast.Identifier{Name: "n"},
				lexer.TokenGreaterThanEqual,
				&ast.Call{
					Expr: &ast.Identifier{Name: "number"},
					Arguments: []ast.Node{
						asttest.NewLiteralChar('A'),
					},
				},
			),
		},
		"import-constant": {
			str: "math.Pi",
			expected: &ast.Key{
				Expr: &ast.Identifier{Name: "math"},
				Key:  &ast.Identifier{Name: "Pi"},
			},
		},
		"method": {
			str: "math.Foo()",
			expected: &ast.Call{
				Expr: &ast.Key{
					Expr: &ast.Identifier{Name: "math"},
					Key:  &ast.Identifier{Name: "Foo"},
				},
			},
		},
		"chained-key": {
			str: "math.Foo().Bar",
			expected: &ast.Key{ // math.Foo().Bar
				Expr: &ast.Call{ // math.Foo()
					Expr: &ast.Key{ // math.Foo
						Expr: &ast.Identifier{Name: "math"},
						Key:  &ast.Identifier{Name: "Foo"},
					},
				},
				Key: &ast.Identifier{Name: "Bar"},
			},
		},
		"chained-method": {
			str: "math.Foo().Bar()",
			expected: &ast.Call{ // math.Foo().Bar()
				Expr: &ast.Key{ // math.Foo().Bar
					Expr: &ast.Call{ // math.Foo()
						Expr: &ast.Key{ // math.Foo
							Expr: &ast.Identifier{Name: "math"},
							Key:  &ast.Identifier{Name: "Foo"},
						},
					},
					Key: &ast.Identifier{Name: "Bar"},
				},
			},
		},
		"get-index": {
			str: `foo[0]`,
			expected: &ast.Key{
				Expr: &ast.Identifier{Name: "foo"},
				Key:  asttest.NewLiteralNumber("0"),
			},
		},
		"cast-string-index": {
			str: `string foo[10]`,
			expected: &ast.Call{
				Expr: &ast.Identifier{Name: "string"},
				Arguments: []ast.Node{
					&ast.Key{
						Expr: &ast.Identifier{Name: "foo"},
						Key:  asttest.NewLiteralNumber("10"),
					},
				},
			},
		},
	} {
		t.Run(testName, func(t *testing.T) {
			str := fmt.Sprintf("func main() { %s }", test.str)
			p := parser.NewParser(0)
			p.ParseString(str, "a.ok")

			assertEqualErrors(t, test.errs, p.Errors())
			if test.expected == nil {
				asttest.AssertEqual(t, map[string]*ast.Func{}, p.Funcs())
			} else {
				asttest.AssertEqual(t, map[string]*ast.Func{
					"main": {
						Name:       "main",
						Statements: []ast.Node{test.expected},
					},
				}, p.Funcs())
			}
			assert.Nil(t, p.Comments())
		})
	}
}
