package parser_test

import (
	"fmt"
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/parser"
)

func TestReturn(t *testing.T) {
	for testName, test := range map[string]struct {
		str      string
		expected ast.Node
		errs     []error
	}{
		"return-zero": {
			str:      "return\n",
			expected: &ast.Return{},
		},
		"return-number": {
			str: "return 123\n",
			expected: &ast.Return{
				Exprs: []ast.Node{
					asttest.NewLiteralNumber("123"),
				},
			},
		},
		"return-two-values": {
			str: "return 123, 'a'\n",
			expected: &ast.Return{
				Exprs: []ast.Node{
					asttest.NewLiteralNumber("123"),
					asttest.NewLiteralChar('a'),
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
