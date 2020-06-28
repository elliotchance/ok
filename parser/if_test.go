package parser_test

import (
	"fmt"
	"testing"

	"github.com/elliotchance/ok/ast"
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
					Right: ast.NewLiteralNumber("3"),
				},
			},
		},
		"if-2": {
			str: "if a == 3 { a = 1\n++a }",
			expected: &ast.If{
				Condition: &ast.Binary{
					Left:  &ast.Identifier{Name: "a"},
					Op:    lexer.TokenEqual,
					Right: ast.NewLiteralNumber("3"),
				},
				True: []ast.Node{
					&ast.Assign{
						Lefts: []ast.Node{
							&ast.Identifier{Name: "a"},
						},
						Rights: []ast.Node{
							ast.NewLiteralNumber("1"),
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
					Right: ast.NewLiteralNumber("3"),
				},
				True: []ast.Node{
					&ast.Assign{
						Lefts: []ast.Node{
							&ast.Identifier{Name: "a"},
						},
						Rights: []ast.Node{
							ast.NewLiteralNumber("1"),
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
					Right: ast.NewLiteralNumber("3"),
				},
				True: []ast.Node{
					&ast.Assign{
						Lefts: []ast.Node{
							&ast.Identifier{Name: "a"},
						},
						Rights: []ast.Node{
							ast.NewLiteralNumber("1"),
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
	} {
		t.Run(testName, func(t *testing.T) {
			str := fmt.Sprintf("func main() { %s }", test.str)
			p := parser.ParseString(str)

			assertEqualErrors(t, test.errs, p.Errors)
			if test.expected == nil {
				assert.Equal(t, map[string]*ast.Func{}, p.File.Funcs)
			} else {
				assert.Equal(t, map[string]*ast.Func{
					"main": {
						Name:       "main",
						Statements: []ast.Node{test.expected},
					},
				}, p.File.Funcs)
			}
			assert.Nil(t, p.File.Comments)
		})
	}
}
