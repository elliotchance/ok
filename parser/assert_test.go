package parser_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/lexer"
	"github.com/elliotchance/ok/parser"
)

func TestAssert(t *testing.T) {
	for testName, test := range map[string]struct {
		str      string
		expected ast.Node
		errs     []error
	}{
		"equals": {
			str: "assert(123 == 456)",
			expected: &ast.Assert{
				Expr: asttest.NewBinary(
					asttest.NewLiteralNumber("123"),
					lexer.TokenEqual,
					asttest.NewLiteralNumber("456"),
				),
			},
		},
		"not-binary": {
			str:      "assert(false)",
			expected: &ast.Assert{},
			errs: []error{
				errors.New("a.ok:1:15 only binary expressions are permitted in assertions"),
			},
		},
		"raise-type": {
			str: "assert(foo() raise Bar)",
			expected: &ast.AssertRaise{
				Call: &ast.Call{
					Expr: &ast.Identifier{Name: "foo"},
				},
				TypeOrValue: &ast.Identifier{Name: "Bar"},
			},
		},
		"raise-value": {
			str: `assert(foo() raise Error("uh oh"))`,
			expected: &ast.AssertRaise{
				Call: &ast.Call{
					Expr: &ast.Identifier{Name: "foo"},
				},
				TypeOrValue: &ast.Call{
					Expr: &ast.Identifier{Name: "Error"},
					Arguments: []ast.Node{
						asttest.NewLiteralString("uh oh"),
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
			asttest.AssertEqual(t, map[string]*ast.Func{
				"1": newFunc(test.expected),
			}, p.Funcs())
		})
	}
}
