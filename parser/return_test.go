package parser_test

import (
	"fmt"
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/parser"

	"github.com/stretchr/testify/assert"
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
					ast.NewLiteralNumber("123"),
				},
			},
		},
		"return-two-values": {
			str: "return 123, 'a'\n",
			expected: &ast.Return{
				Exprs: []ast.Node{
					ast.NewLiteralNumber("123"),
					ast.NewLiteralChar('a'),
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
		})
	}
}
