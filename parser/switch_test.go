package parser_test

import (
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/lexer"
	"github.com/elliotchance/ok/parser"
	"github.com/stretchr/testify/assert"
)

func TestSwitch(t *testing.T) {
	for testName, test := range map[string]struct {
		str      string
		expected *ast.Func
		comments []*ast.Comment
		errs     []error
	}{
		"switch-1": {
			str: "func main() { switch {} }",
			expected: newFunc(
				&ast.Switch{},
			),
		},
		"switch-2": {
			str: "func main() { switch { case a == 1 {} } }",
			expected: newFunc(
				&ast.Switch{
					Cases: []*ast.Case{
						{
							Conditions: []ast.Node{
								asttest.NewBinary(
									&ast.Identifier{Name: "a"},
									lexer.TokenEqual,
									asttest.NewLiteralNumber("1"),
								),
							},
						},
					},
				},
			),
		},
		"switch-3": {
			str: "func main() { switch { case a == 1 {} case a > 2 {} } }",
			expected: newFunc(
				&ast.Switch{
					Cases: []*ast.Case{
						{
							Conditions: []ast.Node{
								asttest.NewBinary(
									&ast.Identifier{Name: "a"},
									lexer.TokenEqual,
									asttest.NewLiteralNumber("1"),
								),
							},
						},
						{
							Conditions: []ast.Node{
								asttest.NewBinary(
									&ast.Identifier{Name: "a"},
									lexer.TokenGreaterThan,
									asttest.NewLiteralNumber("2"),
								),
							},
						},
					},
				},
			),
		},
		"switch-4": {
			str: "func main() { switch { case a == 1 { print(a) } else { } } }",
			expected: newFunc(
				&ast.Switch{
					Cases: []*ast.Case{
						{
							Conditions: []ast.Node{
								asttest.NewBinary(
									&ast.Identifier{Name: "a"},
									lexer.TokenEqual,
									asttest.NewLiteralNumber("1"),
								),
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
					Else: []ast.Node(nil),
				},
			),
		},
		"switch-5": {
			str: "func main() { switch { else { print(b) } } }",
			expected: newFunc(
				&ast.Switch{
					Else: []ast.Node{
						&ast.Call{
							Expr: &ast.Identifier{Name: "print"},
							Arguments: []ast.Node{
								&ast.Identifier{Name: "b"},
							},
						},
					},
				},
			),
		},
		"switch-6": {
			str: "func main() { switch { case a == 1, a < 10 {} } }",
			expected: newFunc(
				&ast.Switch{
					Cases: []*ast.Case{
						{
							Conditions: []ast.Node{
								asttest.NewBinary(
									&ast.Identifier{Name: "a"},
									lexer.TokenEqual,
									asttest.NewLiteralNumber("1"),
								),
								asttest.NewBinary(
									&ast.Identifier{Name: "a"},
									lexer.TokenLessThan,
									asttest.NewLiteralNumber("10"),
								),
							},
						},
					},
				},
			),
		},
		"switch-value-1": {
			str: "func main() { switch a {} }",
			expected: newFunc(
				&ast.Switch{
					Expr: &ast.Identifier{Name: "a"},
				},
			),
		},
		"switch-value-2": {
			str: "func main() { switch a { case 1, 2, 3 {} } }",
			expected: newFunc(
				&ast.Switch{
					Expr: &ast.Identifier{Name: "a"},
					Cases: []*ast.Case{
						{
							Conditions: []ast.Node{
								asttest.NewLiteralNumber("1"),
								asttest.NewLiteralNumber("2"),
								asttest.NewLiteralNumber("3"),
							},
						},
					},
				},
			),
		},
	} {
		t.Run(testName, func(t *testing.T) {
			p := parser.ParseString(test.str, "a.ok")

			assertEqualErrors(t, test.errs, p.Errors())
			asttest.AssertEqual(t, map[string]*ast.Func{
				"main": test.expected,
			}, p.File.Funcs)
			assert.Equal(t, test.comments, p.File.Comments)
		})
	}
}
