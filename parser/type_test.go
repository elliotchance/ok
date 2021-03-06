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
		error    string
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
		"any-string": {
			str: `any "foo"`,
			expected: &ast.Call{
				Expr:      &ast.Identifier{Name: "any"},
				Arguments: []ast.Node{asttest.NewLiteralString("foo")},
			},
		},
		"any-number": {
			str: `any 12.3`,
			expected: &ast.Call{
				Expr:      &ast.Identifier{Name: "any"},
				Arguments: []ast.Node{asttest.NewLiteralNumber("12.3")},
			},
		},
		"newline-not-allowed": {
			str:   "[]\nnumber",
			error: "a.ok:2:1 expecting statement; a.ok:1:1 expecting statement",
		},
	} {
		t.Run(testName, func(t *testing.T) {
			str := fmt.Sprintf("func main() { %s }", test.str)
			p := parser.NewParser(0)
			p.ParseString(str, "a.ok")

			if test.error != "" {
				assert.Equal(t, p.Errors().String(), test.error)
			} else {
				require.Empty(t, p.Errors().String(), str)
				asttest.AssertEqual(t, map[string]*ast.Func{
					"1": newFunc(test.expected),
				}, p.Funcs())
				assert.Nil(t, p.Comments())
			}
		})
	}
}
