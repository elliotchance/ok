package parser_test

import (
	"fmt"
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/lexer"
	"github.com/elliotchance/ok/parser"
	"github.com/elliotchance/ok/types"
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
				// Not "any" because it's redundant and not a runtime time.
				Kind: nil,
			},
		},
		"one-number": {
			str: `{"foo": 123}`,
			expected: &ast.Map{
				Elements: []*ast.KeyValue{
					{
						Key:   asttest.NewLiteralString("foo"),
						Value: asttest.NewLiteralNumber("123"),
					},
				},
			},
		},
		"mixed-literals": {
			str: `{"foo": 1, bar: "2", "b" + "az": true}`,
			expected: &ast.Map{
				Elements: []*ast.KeyValue{
					{
						Key:   asttest.NewLiteralString("foo"),
						Value: asttest.NewLiteralNumber("1"),
					},
					{
						Key:   &ast.Identifier{Name: "bar"},
						Value: asttest.NewLiteralString("2"),
					},
					{
						Key: asttest.NewBinary(
							asttest.NewLiteralString("b"),
							lexer.TokenPlus,
							asttest.NewLiteralString("az"),
						),
						Value: asttest.NewLiteralBool(true),
					},
				},
			},
		},
		"expressions": {
			str: `{"a": 1 + 2, "b": foo("bar") + 3}`,
			expected: &ast.Map{
				Elements: []*ast.KeyValue{
					{
						Key: asttest.NewLiteralString("a"),
						Value: asttest.NewBinary(
							asttest.NewLiteralNumber("1"),
							lexer.TokenPlus,
							asttest.NewLiteralNumber("2"),
						),
					},
					{
						Key: asttest.NewLiteralString("b"),
						Value: asttest.NewBinary(
							&ast.Call{
								FunctionName: "foo",
								Arguments: []ast.Node{
									asttest.NewLiteralString("bar"),
								},
							},
							lexer.TokenPlus,
							asttest.NewLiteralNumber("3"),
						),
					},
				},
			},
		},
		"numbers": {
			str: `{}number {"a": 1, "b": 2, "c": 3}`,
			expected: &ast.Map{
				Kind: types.NumberMap,
				Elements: []*ast.KeyValue{
					{
						Key:   asttest.NewLiteralString("a"),
						Value: asttest.NewLiteralNumber("1"),
					},
					{
						Key:   asttest.NewLiteralString("b"),
						Value: asttest.NewLiteralNumber("2"),
					},
					{
						Key:   asttest.NewLiteralString("c"),
						Value: asttest.NewLiteralNumber("3"),
					},
				},
			},
		},
		"get-index": {
			str: `foo["bar"]`,
			expected: &ast.Key{
				Expr: &ast.Identifier{Name: "foo"},
				Key:  asttest.NewLiteralString("bar"),
			},
		},
		"set-index": {
			str: `foo[bar] = 2`,
			expected: &ast.Assign{
				Lefts: []ast.Node{
					&ast.Key{
						Expr: &ast.Identifier{Name: "foo"},
						Key:  &ast.Identifier{Name: "bar"},
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
