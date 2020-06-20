package parser_test

import (
	"fmt"
	"ok/ast"
	"ok/lexer"
	"ok/parser"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArray(t *testing.T) {
	for testName, test := range map[string]struct {
		str      string
		expected ast.Node
		errs     []error
	}{
		"unknown-empty": {
			str:      "[]",
			expected: &ast.Array{},
		},
		"any-empty": {
			str: "any []",
			expected: &ast.Array{
				Kind: "any",
			},
		},
		"one-number": {
			str: "[123]",
			expected: &ast.Array{
				Elements: []ast.Node{
					ast.NewLiteralNumber("123"),
				},
			},
		},
		"mixed-literals": {
			str: `[1, "2", true]`,
			expected: &ast.Array{
				Elements: []ast.Node{
					ast.NewLiteralNumber("1"),
					ast.NewLiteralString("2"),
					ast.NewLiteralBool(true),
				},
			},
		},
		"expressions": {
			str: `[1 + 2, foo("bar")]`,
			expected: &ast.Array{
				Elements: []ast.Node{
					ast.NewBinary(
						ast.NewLiteralNumber("1"),
						lexer.TokenPlus,
						ast.NewLiteralNumber("2"),
					),
					&ast.Call{
						FunctionName: "foo",
						Arguments: []ast.Node{
							ast.NewLiteralString("bar"),
						},
					},
				},
			},
		},
		"numbers": {
			str: `[]number [1, 2, 3]`,
			expected: &ast.Array{
				Kind: "[]number",
				Elements: []ast.Node{
					ast.NewLiteralNumber("1"),
					ast.NewLiteralNumber("2"),
					ast.NewLiteralNumber("3"),
				},
			},
		},
		"get-index": {
			str: `foo[0]`,
			expected: &ast.Key{
				Expr: &ast.Identifier{Name: "foo"},
				Key:  ast.NewLiteralNumber("0"),
			},
		},
		"set-index": {
			str: `foo[1] = 2`,
			expected: ast.NewBinary(
				&ast.Key{
					Expr: &ast.Identifier{Name: "foo"},
					Key:  ast.NewLiteralNumber("1"),
				},
				lexer.TokenAssign,
				ast.NewLiteralNumber("2"),
			),
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
