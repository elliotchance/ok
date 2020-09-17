package parser_test

import (
	"fmt"
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/parser"
	"github.com/elliotchance/ok/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestType(t *testing.T) {
	for testName, test := range map[string]struct {
		str      string
		expected ast.Node
	}{
		"any": {
			str: "any []",
			expected: &ast.Array{
				// Not "any" because it's redundant and not a runtime time.
				Kind: nil,
			},
		},
		"any-array": {
			str:      "[]any []",
			expected: &ast.Array{Kind: types.AnyArray},
		},
		"string-array": {
			str:      "[] string []",
			expected: &ast.Array{Kind: types.StringArray},
		},
		"any-map": {
			str:      "{}any {}",
			expected: &ast.Map{Kind: types.AnyMap},
		},
		"string-map": {
			str:      "{} string {}",
			expected: &ast.Map{Kind: types.StringMap},
		},
		"array-person": {
			str:      "[] Person []",
			expected: &ast.Array{Kind: types.TypeFromString("[]Person")},
		},
		"map-person": {
			str:      "{}Person {}",
			expected: &ast.Map{Kind: types.TypeFromString("{}Person")},
		},
		"func-1": {
			str:      "{}func(number) {}",
			expected: &ast.Map{Kind: types.TypeFromString("{}func(number)")},
		},
		"func-2": {
			str:      "{}func (number) {}",
			expected: &ast.Map{Kind: types.TypeFromString("{}func(number)")},
		},
		"func-3": {
			str:      "[]func(number,Foo) []",
			expected: &ast.Array{Kind: types.TypeFromString("[]func(number, Foo)")},
		},
		"func-4": {
			str:      "[]func() []",
			expected: &ast.Array{Kind: types.TypeFromString("[]func()")},
		},
		"func-5": {
			str:      "[]func() number []",
			expected: &ast.Array{Kind: types.TypeFromString("[]func() number")},
		},
		"func-6": {
			str:      "[]func() (Foo,bar) []",
			expected: &ast.Array{Kind: types.TypeFromString("[]func() (Foo, bar)")},
		},
	} {
		t.Run(testName, func(t *testing.T) {
			str := fmt.Sprintf("func main() { %s }", test.str)
			p := parser.ParseString(str, "a.ok")

			require.Empty(t, p.Errors().String(), str)
			asttest.AssertEqual(t, map[string]*ast.Func{
				"main": newFunc(test.expected),
			}, p.File.Funcs)
			assert.Nil(t, p.File.Comments)
		})
	}
}
