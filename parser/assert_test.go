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
