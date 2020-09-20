package parser_test

import (
	"fmt"
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/lexer"
	"github.com/elliotchance/ok/parser"
	"github.com/stretchr/testify/assert"
)

func TestFor(t *testing.T) {
	for testName, test := range map[string]struct {
		str      string
		expected *ast.For
		comments []*ast.Comment
		errs     []error
	}{
		"for-empty-body": {
			str: "for true {}",
			expected: &ast.For{
				Condition: asttest.NewLiteralBool(true),
			},
		},
		"for-1": {
			str: "for true { print(a)\nprint(b) }",
			expected: &ast.For{
				Condition: asttest.NewLiteralBool(true),
				Statements: []ast.Node{
					&ast.Call{
						Expr: &ast.Identifier{Name: "print"},
						Arguments: []ast.Node{
							&ast.Identifier{Name: "a"},
						},
					},
					&ast.Call{
						Expr: &ast.Identifier{Name: "print"},
						Arguments: []ast.Node{
							&ast.Identifier{Name: "b"},
						},
					},
				},
			},
		},
		"for-empty-without-condition": {
			str:      "for {}",
			expected: &ast.For{},
		},
		"for-2": {
			str: "for { print(a)\nprint(b) }",
			expected: &ast.For{
				Statements: []ast.Node{
					&ast.Call{
						Expr: &ast.Identifier{Name: "print"},
						Arguments: []ast.Node{
							&ast.Identifier{Name: "a"},
						},
					},
					&ast.Call{
						Expr: &ast.Identifier{Name: "print"},
						Arguments: []ast.Node{
							&ast.Identifier{Name: "b"},
						},
					},
				},
			},
		},
		"break-1": {
			str: "for { break\nprint(a) }",
			expected: &ast.For{
				Statements: []ast.Node{
					&ast.Break{},
					&ast.Call{
						Expr: &ast.Identifier{Name: "print"},
						Arguments: []ast.Node{
							&ast.Identifier{Name: "a"},
						},
					},
				},
			},
		},
		"continue-1": {
			str: "for { print(a)\ncontinue }",
			expected: &ast.For{
				Statements: []ast.Node{
					&ast.Call{
						Expr: &ast.Identifier{Name: "print"},
						Arguments: []ast.Node{
							&ast.Identifier{Name: "a"},
						},
					},
					&ast.Continue{},
				},
			},
		},
		"for-3": {
			str: "for a = 0; a < 10; ++a { print(a) }",
			expected: &ast.For{
				Init: &ast.Assign{
					Lefts: []ast.Node{
						&ast.Identifier{Name: "a"},
					},
					Rights: []ast.Node{
						asttest.NewLiteralNumber("0"),
					},
				},
				Condition: &ast.Binary{
					Left:  &ast.Identifier{Name: "a"},
					Op:    lexer.TokenLessThan,
					Right: asttest.NewLiteralNumber("10"),
				},
				Next: &ast.Unary{
					Op:   lexer.TokenIncrement,
					Expr: &ast.Identifier{Name: "a"},
				},
				Statements: []ast.Node{
					&ast.Call{
						Expr: &ast.Identifier{Name: "print"},
						Arguments: []ast.Node{
							&ast.Identifier{Name: "a"},
						},
					},
				},
			},
		},
		"for-4-assign": {
			str: "for a = 0; a < 10; a = a + 1 { print(a) }",
			expected: &ast.For{
				Init: &ast.Assign{
					Lefts: []ast.Node{
						&ast.Identifier{Name: "a"},
					},
					Rights: []ast.Node{
						asttest.NewLiteralNumber("0"),
					},
				},
				Condition: &ast.Binary{
					Left:  &ast.Identifier{Name: "a"},
					Op:    lexer.TokenLessThan,
					Right: asttest.NewLiteralNumber("10"),
				},
				Next: &ast.Assign{
					Lefts: []ast.Node{
						&ast.Identifier{Name: "a"},
					},
					Rights: []ast.Node{
						&ast.Binary{
							Left:  &ast.Identifier{Name: "a"},
							Op:    lexer.TokenPlus,
							Right: asttest.NewLiteralNumber("1"),
						},
					},
				},
				Statements: []ast.Node{
					&ast.Call{
						Expr: &ast.Identifier{Name: "print"},
						Arguments: []ast.Node{
							&ast.Identifier{Name: "a"},
						},
					},
				},
			},
		},
		"for-in-1": {
			str: "for v in something { }",
			expected: &ast.For{
				Condition: &ast.In{
					Value: "v",
					Expr:  &ast.Identifier{Name: "something"},
				},
			},
		},
		"for-in-2": {
			str: "for v, k in something { }",
			expected: &ast.For{
				Condition: &ast.In{
					Value: "v",
					Key:   "k",
					Expr:  &ast.Identifier{Name: "something"},
				},
			},
		},
		"for-in-3": {
			str: "for i = 0; v, k in something; ++i { }",
			expected: &ast.For{
				Init: &ast.Assign{
					Lefts: []ast.Node{
						&ast.Identifier{Name: "i"},
					},
					Rights: []ast.Node{
						asttest.NewLiteralNumber("0"),
					},
				},
				Condition: &ast.In{
					Value: "v",
					Key:   "k",
					Expr:  &ast.Identifier{Name: "something"},
				},
				Next: &ast.Unary{
					Expr: &ast.Identifier{Name: "i"},
					Op:   lexer.TokenIncrement,
				},
			},
		},
	} {
		t.Run(testName, func(t *testing.T) {
			str := fmt.Sprintf("func main() { %s }", test.str)
			p := parser.NewParser()
			p.ParseString(str, "a.ok")

			assertEqualErrors(t, test.errs, p.Errors())
			asttest.AssertEqual(t, map[string]*ast.Func{
				"main": newFunc(test.expected),
			}, p.Funcs())
			assert.Equal(t, test.comments, p.Comments())
		})
	}
}
