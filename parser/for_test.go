package parser_test

import (
	"fmt"
	"ok/ast"
	"ok/lexer"
	"ok/parser"
	"testing"

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
				Condition: ast.NewLiteralBool(true),
			},
		},
		"for-1": {
			str: "for true { print(a)\nprint(b) }",
			expected: &ast.For{
				Condition: ast.NewLiteralBool(true),
				Statements: []ast.Node{
					&ast.Call{
						FunctionName: "print",
						Arguments: []ast.Node{
							&ast.Identifier{Name: "a"},
						},
					},
					&ast.Call{
						FunctionName: "print",
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
						FunctionName: "print",
						Arguments: []ast.Node{
							&ast.Identifier{Name: "a"},
						},
					},
					&ast.Call{
						FunctionName: "print",
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
						FunctionName: "print",
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
						FunctionName: "print",
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
						ast.NewLiteralNumber("0"),
					},
				},
				Condition: &ast.Binary{
					Left:  &ast.Identifier{Name: "a"},
					Op:    lexer.TokenLessThan,
					Right: ast.NewLiteralNumber("10"),
				},
				Next: &ast.Unary{
					Op:   lexer.TokenIncrement,
					Expr: &ast.Identifier{Name: "a"},
				},
				Statements: []ast.Node{
					&ast.Call{
						FunctionName: "print",
						Arguments: []ast.Node{
							&ast.Identifier{Name: "a"},
						},
					},
				},
			},
		},
		"for-in-1": {
			str: "for k in something { }",
			expected: &ast.For{
				Condition: &ast.In{
					Key:  "k",
					Expr: &ast.Identifier{Name: "something"},
				},
			},
		},
		"for-in-2": {
			str: "for k, v in something { }",
			expected: &ast.For{
				Condition: &ast.In{
					Key:   "k",
					Value: "v",
					Expr:  &ast.Identifier{Name: "something"},
				},
			},
		},
		"for-in-3": {
			str: "for i = 0; k, v in something; ++i { }",
			expected: &ast.For{
				Init: &ast.Assign{
					Lefts: []ast.Node{
						&ast.Identifier{Name: "i"},
					},
					Rights: []ast.Node{
						ast.NewLiteralNumber("0"),
					},
				},
				Condition: &ast.In{
					Key:   "k",
					Value: "v",
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
			p := parser.ParseString(str)

			assertEqualErrors(t, test.errs, p.Errors)
			assert.Equal(t, map[string]*ast.Func{
				"main": newFunc(test.expected),
			}, p.File.Funcs)
			assert.Equal(t, test.comments, p.File.Comments)
		})
	}
}
