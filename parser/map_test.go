package parser_test

import (
	"fmt"
	"ok/ast"
	"ok/lexer"
	"ok/parser"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	for testName, test := range map[string]struct {
		str      string
		expected ast.Node
		errs     []error
	}{
		"any-empty": {
			str: "any {}",
			expected: &ast.Map{
				Kind: "any",
			},
		},
		"one-number": {
			str: `{"foo": 123}`,
			expected: &ast.Map{
				Elements: []*ast.KeyValue{
					{
						Key:   ast.NewLiteralString("foo"),
						Value: ast.NewLiteralNumber("123"),
					},
				},
			},
		},
		"mixed-literals": {
			str: `{"foo": 1, bar: "2", "b" + "az": true}`,
			expected: &ast.Map{
				Elements: []*ast.KeyValue{
					{
						Key:   ast.NewLiteralString("foo"),
						Value: ast.NewLiteralNumber("1"),
					},
					{
						Key:   &ast.Identifier{Name: "bar"},
						Value: ast.NewLiteralString("2"),
					},
					{
						Key: ast.NewBinary(
							ast.NewLiteralString("b"),
							lexer.TokenPlus,
							ast.NewLiteralString("az"),
						),
						Value: ast.NewLiteralBool(true),
					},
				},
			},
		},
		"expressions": {
			str: `{"a": 1 + 2, "b": foo("bar") + 3}`,
			expected: &ast.Map{
				Elements: []*ast.KeyValue{
					{
						Key: ast.NewLiteralString("a"),
						Value: ast.NewBinary(
							ast.NewLiteralNumber("1"),
							lexer.TokenPlus,
							ast.NewLiteralNumber("2"),
						),
					},
					{
						Key: ast.NewLiteralString("b"),
						Value: ast.NewBinary(
							&ast.Call{
								FunctionName: "foo",
								Arguments: []ast.Node{
									ast.NewLiteralString("bar"),
								},
							},
							lexer.TokenPlus,
							ast.NewLiteralNumber("3"),
						),
					},
				},
			},
		},
		"numbers": {
			str: `{}number {"a": 1, "b": 2, "c": 3}`,
			expected: &ast.Map{
				Kind: "{}number",
				Elements: []*ast.KeyValue{
					{
						Key:   ast.NewLiteralString("a"),
						Value: ast.NewLiteralNumber("1"),
					},
					{
						Key:   ast.NewLiteralString("b"),
						Value: ast.NewLiteralNumber("2"),
					},
					{
						Key:   ast.NewLiteralString("c"),
						Value: ast.NewLiteralNumber("3"),
					},
				},
			},
		},
		"get-index": {
			str: `foo["bar"]`,
			expected: &ast.Key{
				Expr: &ast.Identifier{Name: "foo"},
				Key:  ast.NewLiteralString("bar"),
			},
		},
		"set-index": {
			str: `foo[bar] = 2`,
			expected: ast.NewBinary(
				&ast.Key{
					Expr: &ast.Identifier{Name: "foo"},
					Key:  &ast.Identifier{Name: "bar"},
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
