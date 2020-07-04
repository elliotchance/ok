package parser_test

import (
	"fmt"
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/lexer"
	"github.com/elliotchance/ok/parser"
)

func TestInterpolate(t *testing.T) {
	for testName, test := range map[string]struct {
		str      string
		expected *ast.Interpolate
		errs     []error
	}{
		"1": {
			str: `"{name}"`,
			expected: &ast.Interpolate{
				Parts: []ast.Node{
					&ast.Group{
						Expr: &ast.Identifier{Name: "name"},
					},
				},
			},
		},
		"2": {
			str: `"hi {1+2} there"`,
			expected: &ast.Interpolate{
				Parts: []ast.Node{
					asttest.NewLiteralString("hi "),
					&ast.Group{
						Expr: asttest.NewBinary(
							asttest.NewLiteralNumber("1"),
							lexer.TokenPlus,
							asttest.NewLiteralNumber("2"),
						),
					},
					asttest.NewLiteralString(" there"),
				},
			},
		},
	} {
		t.Run(testName, func(t *testing.T) {
			str := fmt.Sprintf("func main() { %s }", test.str)
			p := parser.ParseString(str, "a.ok")

			assertEqualErrors(t, test.errs, p.Errors())
			asttest.AssertEqual(t, map[string]*ast.Func{
				"main": newFunc(test.expected),
			}, p.File.Funcs)
		})
	}
}
