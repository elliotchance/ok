package parser_test

import (
	"ok/ast"
	"ok/lexer"
	"ok/parser"
	"testing"

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
								ast.NewBinary(
									&ast.Identifier{Name: "a"},
									lexer.TokenEqual,
									ast.NewLiteralNumber("1"),
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
								ast.NewBinary(
									&ast.Identifier{Name: "a"},
									lexer.TokenEqual,
									ast.NewLiteralNumber("1"),
								),
							},
						},
						{
							Conditions: []ast.Node{
								ast.NewBinary(
									&ast.Identifier{Name: "a"},
									lexer.TokenGreaterThan,
									ast.NewLiteralNumber("2"),
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
								ast.NewBinary(
									&ast.Identifier{Name: "a"},
									lexer.TokenEqual,
									ast.NewLiteralNumber("1"),
								),
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
							FunctionName: "print",
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
								ast.NewBinary(
									&ast.Identifier{Name: "a"},
									lexer.TokenEqual,
									ast.NewLiteralNumber("1"),
								),
								ast.NewBinary(
									&ast.Identifier{Name: "a"},
									lexer.TokenLessThan,
									ast.NewLiteralNumber("10"),
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
								ast.NewLiteralNumber("1"),
								ast.NewLiteralNumber("2"),
								ast.NewLiteralNumber("3"),
							},
						},
					},
				},
			),
		},
	} {
		t.Run(testName, func(t *testing.T) {
			p := parser.ParseString(test.str)

			assertEqualErrors(t, test.errs, p.Errors)
			assert.Equal(t, map[string]*ast.Func{
				"main": test.expected,
			}, p.File.Funcs)
			assert.Equal(t, test.comments, p.File.Comments)
		})
	}
}
