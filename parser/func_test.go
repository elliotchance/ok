package parser_test

import (
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/parser"
	"github.com/elliotchance/ok/types"
	"github.com/stretchr/testify/assert"
)

func TestFunc(t *testing.T) {
	for testName, test := range map[string]struct {
		str      string
		expected map[string]*ast.Func
	}{
		"zero-args-zero-returns": {
			str: "func foo() {}",
			expected: map[string]*ast.Func{
				"foo": {Name: "foo"},
			},
		},
		"one-arg-zero-returns": {
			str: "func foo(bar number) {}",
			expected: map[string]*ast.Func{
				"foo": {
					Name: "foo",
					Arguments: []*ast.Argument{
						{Name: "bar", Type: types.Number},
					},
				},
			},
		},
		"two-args-zero-returns": {
			str: "func foo(bar number, baz string) {}",
			expected: map[string]*ast.Func{
				"foo": {
					Name: "foo",
					Arguments: []*ast.Argument{
						{Name: "bar", Type: types.Number},
						{Name: "baz", Type: types.String},
					},
				},
			},
		},
		"three-args-zero-returns": {
			str: "func foo(bar number, baz string, qux [] bool) {}",
			expected: map[string]*ast.Func{
				"foo": {
					Name: "foo",
					Arguments: []*ast.Argument{
						{Name: "bar", Type: types.Number},
						{Name: "baz", Type: types.String},
						{Name: "qux", Type: types.BoolArray},
					},
				},
			},
		},
		"zero-args-one-return": {
			str: "func bar() number {}",
			expected: map[string]*ast.Func{
				"bar": {
					Name:    "bar",
					Returns: []*types.Type{types.Number},
				},
			},
		},
		"zero-args-two-returns": {
			str: "func bar() (number, string) {}",
			expected: map[string]*ast.Func{
				"bar": {
					Name:    "bar",
					Returns: []*types.Type{types.Number, types.String},
				},
			},
		},
		"three-compressed-args-zero-returns": {
			str: "func foo(bar, baz string, qux {}number) {}",
			expected: map[string]*ast.Func{
				"foo": {
					Name: "foo",
					Arguments: []*ast.Argument{
						{Name: "bar", Type: types.String},
						{Name: "baz", Type: types.String},
						{Name: "qux", Type: types.NumberMap},
					},
				},
			},
		},
		"multiple-functions": {
			str: "func foo() {}\nfunc bar() {}",
			expected: map[string]*ast.Func{
				"foo": {
					Name: "foo",
				},
				"bar": {
					Name: "bar",
				},
			},
		},
		"exported-function": {
			str: "func Abs() {}",
			expected: map[string]*ast.Func{
				"Abs": {
					Name: "Abs",
				},
			},
		},
		"try-finally-2": {
			str: "func main() {" +
				"try { print() } finally { foo() } " +
				"try { print() } finally { bar() }" +
				"}",
			expected: map[string]*ast.Func{
				"main": {
					Name: "main",
					Statements: []ast.Node{
						&ast.ErrorScope{
							Statements: []ast.Node{
								&ast.Call{
									Expr: &ast.Identifier{Name: "print"},
								},
							},
							Finally: &ast.Finally{
								Index: 0,
								Statements: []ast.Node{
									&ast.Call{
										Expr: &ast.Identifier{Name: "foo"},
									},
								},
							},
						},
						&ast.ErrorScope{
							Statements: []ast.Node{
								&ast.Call{
									Expr: &ast.Identifier{Name: "print"},
								},
							},
							Finally: &ast.Finally{
								Index: 1,
								Statements: []ast.Node{
									&ast.Call{
										Expr: &ast.Identifier{Name: "bar"},
									},
								},
							},
						},
					},
				},
			},
		},
		"try-finally-3": {
			str: "func main() {" +
				"try { print() } finally { foo() } " +
				"}" +
				"func main2() {" +
				"try { print() } finally { bar() }" +
				"}",
			expected: map[string]*ast.Func{
				"main": {
					Name: "main",
					Statements: []ast.Node{
						&ast.ErrorScope{
							Statements: []ast.Node{
								&ast.Call{
									Expr: &ast.Identifier{Name: "print"},
								},
							},
							Finally: &ast.Finally{
								Index: 0,
								Statements: []ast.Node{
									&ast.Call{
										Expr: &ast.Identifier{Name: "foo"},
									},
								},
							},
						},
					},
				},
				"main2": {
					Name: "main2",
					Statements: []ast.Node{
						&ast.ErrorScope{
							Statements: []ast.Node{
								&ast.Call{
									Expr: &ast.Identifier{Name: "print"},
								},
							},
							Finally: &ast.Finally{
								Index: 0,
								Statements: []ast.Node{
									&ast.Call{
										Expr: &ast.Identifier{Name: "bar"},
									},
								},
							},
						},
					},
				},
			},
		},
		"function-without-name": {
			str: "func () {}",
			expected: map[string]*ast.Func{
				"1": {Name: "1"},
			},
		},
	} {
		t.Run(testName, func(t *testing.T) {
			p := parser.NewParser()
			p.ParseString(test.str, "a.ok")

			assert.Nil(t, p.Errors())
			asttest.AssertEqual(t, test.expected, p.Funcs())
			assert.Nil(t, p.Comments())
		})
	}
}
