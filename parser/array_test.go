package parser_test

import (
	"fmt"
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/lexer"
	"github.com/elliotchance/ok/parser"
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
					asttest.NewLiteralNumber("123"),
				},
			},
		},
		"mixed-literals": {
			str: `[1, "2", true]`,
			expected: &ast.Array{
				Elements: []ast.Node{
					asttest.NewLiteralNumber("1"),
					asttest.NewLiteralString("2"),
					asttest.NewLiteralBool(true),
				},
			},
		},
		"expressions": {
			str: `[1 + 2, foo("bar")]`,
			expected: &ast.Array{
				Elements: []ast.Node{
					asttest.NewBinary(
						asttest.NewLiteralNumber("1"),
						lexer.TokenPlus,
						asttest.NewLiteralNumber("2"),
					),
					&ast.Call{
						FunctionName: "foo",
						Arguments: []ast.Node{
							asttest.NewLiteralString("bar"),
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
					asttest.NewLiteralNumber("1"),
					asttest.NewLiteralNumber("2"),
					asttest.NewLiteralNumber("3"),
				},
			},
		},
		"get-index": {
			str: `foo[0]`,
			expected: &ast.Key{
				Expr: &ast.Identifier{Name: "foo"},
				Key:  asttest.NewLiteralNumber("0"),
			},
		},
		"set-index": {
			str: `foo[1] = 2`,
			expected: &ast.Assign{
				Lefts: []ast.Node{
					&ast.Key{
						Expr: &ast.Identifier{Name: "foo"},
						Key:  asttest.NewLiteralNumber("1"),
					},
				},
				Rights: []ast.Node{
					asttest.NewLiteralNumber("2"),
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
