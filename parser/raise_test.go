package parser_test

import (
	"fmt"
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/parser"
)

func TestRaise(t *testing.T) {
	for testName, test := range map[string]struct {
		str      string
		expected ast.Node
		errs     []error
	}{
		"raise-error-no-args": {
			str: "raise Error()",
			expected: &ast.Raise{
				Err: &ast.Call{
					Expr: &ast.Identifier{Name: "Error"},
				},
			},
		},
		"raise-error-one-arg": {
			str: "raise Error(123)",
			expected: &ast.Raise{
				Err: &ast.Call{
					Expr: &ast.Identifier{Name: "Error"},
					Arguments: []ast.Node{
						asttest.NewLiteralNumber("123"),
					},
				},
			},
		},
		"raise-error-two-args": {
			str: "raise MyError(123, \"foo\")",
			expected: &ast.Raise{
				Err: &ast.Call{
					Expr: &ast.Identifier{Name: "MyError"},
					Arguments: []ast.Node{
						asttest.NewLiteralNumber("123"),
						asttest.NewLiteralString("foo"),
					},
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
		})
	}
}
