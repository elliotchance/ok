package parser_test

import (
	"fmt"
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/parser"

	"github.com/stretchr/testify/assert"
)

func TestType(t *testing.T) {
	for testName, test := range map[string]struct {
		str      string
		expected string
	}{
		"any": {
			str:      "any",
			expected: "any",
		},
		"bool": {
			str:      "bool",
			expected: "bool",
		},
		"char": {
			str:      "char",
			expected: "char",
		},
		"data": {
			str:      "data",
			expected: "data",
		},
		"number": {
			str:      "number",
			expected: "number",
		},
		"string": {
			str:      "string",
			expected: "string",
		},
		"any-array": {
			str:      "[]any",
			expected: "[]any",
		},
		"string-array": {
			str:      "[] string",
			expected: "[]string",
		},
		"any-map": {
			str:      "{}any",
			expected: "{}any",
		},
		"string-map": {
			str:      "{} string",
			expected: "{}string",
		},
	} {
		t.Run(testName, func(t *testing.T) {
			str := fmt.Sprintf("func main() { %s [] }", test.str)
			p := parser.ParseString(str)

			assert.Nil(t, p.Errors)
			assert.Equal(t, map[string]*ast.Func{
				"main": newFunc(&ast.Array{Kind: test.expected}),
			}, p.File.Funcs)
			assert.Nil(t, p.File.Comments)
		})
	}
}
