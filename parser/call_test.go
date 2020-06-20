package parser_test

import (
	"fmt"
	"ok/ast"
	"ok/parser"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCall(t *testing.T) {
	for testName, test := range map[string]struct {
		str      string
		expected *ast.Call
	}{
		"no-args": {
			str: "foo()",
			expected: &ast.Call{
				FunctionName: "foo",
			},
		},
		"one-arg": {
			str: `bar("baz")`,
			expected: &ast.Call{
				FunctionName: "bar",
				Arguments: []ast.Node{
					ast.NewLiteralString("baz"),
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
