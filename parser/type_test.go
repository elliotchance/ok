package parser_test

import (
	"fmt"
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/parser"

	"github.com/stretchr/testify/assert"
)

func TestType(t *testing.T) {
	for testName, test := range map[string]struct {
		str      string
		expected ast.Node
	}{
		"any": {
			str:      "any []",
			expected: &ast.Array{Kind: "any"},
		},
		"any-array": {
			str:      "[]any []",
			expected: &ast.Array{Kind: "[]any"},
		},
		"string-array": {
			str:      "[] string []",
			expected: &ast.Array{Kind: "[]string"},
		},
		"any-map": {
			str:      "{}any {}",
			expected: &ast.Map{Kind: "{}any"},
		},
		"string-map": {
			str:      "{} string {}",
			expected: &ast.Map{Kind: "{}string"},
		},
		"array-person": {
			str:      "[] Person []",
			expected: &ast.Array{Kind: "[]Person"},
		},
		"map-person": {
			str:      "{}Person {}",
			expected: &ast.Map{Kind: "{}Person"},
		},
	} {
		t.Run(testName, func(t *testing.T) {
			str := fmt.Sprintf("func main() { %s }", test.str)
			p := parser.ParseString(str, "a.ok")

			assert.Nil(t, p.Errors())
			asttest.AssertEqual(t, map[string]*ast.Func{
				"main": newFunc(test.expected),
			}, p.File.Funcs)
			assert.Nil(t, p.File.Comments)
		})
	}
}
