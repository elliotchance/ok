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

func TestIf(t *testing.T) {
	for testName, test := range map[string]struct {
		str      string
		expected *ast.If
		errs     []error
	}{
		"if-1": {
			str: "if a == 3 { }",
			expected: &ast.If{
				Condition: &ast.Binary{
					Left:  &ast.Identifier{Name: "a"},
					Op:    lexer.TokenEqual,
					Right: asttest.NewLiteralNumber("3"),
				},
			},
		},
		"if-2": {
			str: "if a == 3 { a = 1\n++a }",
			expected: &ast.If{
				Condition: &ast.Binary{
					Left:  &ast.Identifier{Name: "a"},
					Op:    lexer.TokenEqual,
					Right: asttest.NewLiteralNumber("3"),
				},
				True: []ast.Node{
					&ast.Assign{
						Lefts: []ast.Node{
							&ast.Identifier{Name: "a"},
						},
						Rights: []ast.Node{
							asttest.NewLiteralNumber("1"),
						},
					},
					&ast.Unary{
						Op:   lexer.TokenIncrement,
						Expr: &ast.Identifier{Name: "a"},
					},
				},
			},
		},
		"if-else-1": {
			str: "if a == 3 { a = 1 } else { }",
			expected: &ast.If{
				Condition: &ast.Binary{
					Left:  &ast.Identifier{Name: "a"},
					Op:    lexer.TokenEqual,
					Right: asttest.NewLiteralNumber("3"),
				},
				True: []ast.Node{
					&ast.Assign{
						Lefts: []ast.Node{
							&ast.Identifier{Name: "a"},
						},
						Rights: []ast.Node{
							asttest.NewLiteralNumber("1"),
						},
					},
				},
			},
		},
		"if-else-2": {
			str: "if a == 3 { a = 1 } else { ++a }",
			expected: &ast.If{
				Condition: &ast.Binary{
					Left:  &ast.Identifier{Name: "a"},
					Op:    lexer.TokenEqual,
					Right: asttest.NewLiteralNumber("3"),
				},
				True: []ast.Node{
					&ast.Assign{
						Lefts: []ast.Node{
							&ast.Identifier{Name: "a"},
						},
						Rights: []ast.Node{
							asttest.NewLiteralNumber("1"),
						},
					},
				},
				False: []ast.Node{
					&ast.Unary{
						Op:   lexer.TokenIncrement,
						Expr: &ast.Identifier{Name: "a"},
					},
				},
			},
		},
		"if-is-1": {
			str: "if a is number { }",
			expected: &ast.If{
				Condition: &ast.Binary{
					Left:  &ast.Identifier{Name: "a"},
					Op:    lexer.TokenIs,
					Right: &ast.Identifier{Name: "number"},
				},
			},
		},
		"if-is-2": {
			str: "if foo is [] string { }",
			expected: &ast.If{
				Condition: &ast.Binary{
					Left:  &ast.Identifier{Name: "foo"},
					Op:    lexer.TokenIs,
					Right: &ast.Identifier{Name: "[]string"},
				},
			},
		},
		"if-is-3": {
			str: "if bar is {}bool { }",
			expected: &ast.If{
				Condition: &ast.Binary{
					Left:  &ast.Identifier{Name: "bar"},
					Op:    lexer.TokenIs,
					Right: &ast.Identifier{Name: "{}bool"},
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
					"1": {
						Name:       "main",
						Statements: []ast.Node{test.expected},
					},
				}, p.Funcs())
			}
			assert.Nil(t, p.Comments())
		})
	}
}
