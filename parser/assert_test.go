package parser_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/lexer"
	"github.com/elliotchance/ok/parser"
	"github.com/stretchr/testify/assert"
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
				Expr: ast.NewBinary(
					ast.NewLiteralNumber("123"),
					lexer.TokenEqual,
					ast.NewLiteralNumber("456"),
				),
			},
		},
		"not-binary": {
			str:      "assert(false)",
			expected: &ast.Assert{},
			errs: []error{
				errors.New("only binary expressions are permitted in assertions"),
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
