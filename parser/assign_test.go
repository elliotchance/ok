package parser_test

import (
	"fmt"
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/parser"

	"github.com/stretchr/testify/assert"
)

func TestAssign(t *testing.T) {
	for testName, test := range map[string]struct {
		str      string
		expected *ast.Assign
	}{
		"single-literal": {
			str: "a = 123",
			expected: &ast.Assign{
				Lefts: []ast.Node{
					&ast.Identifier{Name: "a"},
				},
				Rights: []ast.Node{
					ast.NewLiteralNumber("123"),
				},
			},
		},
		"double-literal": {
			str: `a, b = 123, "foo"`,
			expected: &ast.Assign{
				Lefts: []ast.Node{
					&ast.Identifier{Name: "a"},
					&ast.Identifier{Name: "b"},
				},
				Rights: []ast.Node{
					ast.NewLiteralNumber("123"),
					ast.NewLiteralString("foo"),
				},
			},
		},
		"variable": {
			str: `foo = bar`,
			expected: &ast.Assign{
				Lefts: []ast.Node{
					&ast.Identifier{Name: "foo"},
				},
				Rights: []ast.Node{
					&ast.Identifier{Name: "bar"},
				},
			},
		},
	} {
		t.Run(testName, func(t *testing.T) {
			str := fmt.Sprintf("func main() { %s }", test.str)
			p := parser.ParseString(str)

			assert.Nil(t, p.Errors)
			assert.Equal(t, map[string]*ast.Func{
				"main": newFunc(test.expected),
			}, p.File.Funcs)
			assert.Nil(t, p.File.Comments)
		})
	}
}
