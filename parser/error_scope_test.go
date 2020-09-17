package parser_test

import (
	"fmt"
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/parser"
	"github.com/elliotchance/ok/types"
	"github.com/stretchr/testify/assert"
)

func TestErrorScope(t *testing.T) {
	for testName, test := range map[string]struct {
		str      string
		expected *ast.ErrorScope
		errs     []error
	}{
		"only-empty-try": {
			str:      "try {}",
			expected: &ast.ErrorScope{},
		},
		"only-try": {
			str: "try { print() }",
			expected: &ast.ErrorScope{
				Statements: []ast.Node{
					&ast.Call{
						FunctionName: "print",
					},
				},
			},
		},
		"try-on-1": {
			str: "try { print() } on SomeError {}",
			expected: &ast.ErrorScope{
				Statements: []ast.Node{
					&ast.Call{
						FunctionName: "print",
					},
				},
				On: []*ast.On{
					{
						Type: types.NewUnresolvedInterface("SomeError"),
					},
				},
			},
		},
		"try-on-2": {
			str: "try { print() } on SomeError {} on SomethingElse { foo() }",
			expected: &ast.ErrorScope{
				Statements: []ast.Node{
					&ast.Call{
						FunctionName: "print",
					},
				},
				On: []*ast.On{
					{
						Type: types.NewUnresolvedInterface("SomeError"),
					},
					{
						Type: types.NewUnresolvedInterface("SomethingElse"),
						Statements: []ast.Node{
							&ast.Call{
								FunctionName: "foo",
							},
						},
					},
				},
			},
		},
		"try-on-finally": {
			str: "try { print() } on SomethingElse { foo() } finally { bar() }",
			expected: &ast.ErrorScope{
				Statements: []ast.Node{
					&ast.Call{
						FunctionName: "print",
					},
				},
				On: []*ast.On{
					{
						Type: types.NewUnresolvedInterface("SomethingElse"),
						Statements: []ast.Node{
							&ast.Call{
								FunctionName: "foo",
							},
						},
					},
				},
				Finally: &ast.Finally{
					Index: 0,
					Statements: []ast.Node{
						&ast.Call{
							FunctionName: "bar",
						},
					},
				},
			},
		},
		"try-finally": {
			str: "try { print() } finally { foo() }",
			expected: &ast.ErrorScope{
				Statements: []ast.Node{
					&ast.Call{
						FunctionName: "print",
					},
				},
				Finally: &ast.Finally{
					Index: 0,
					Statements: []ast.Node{
						&ast.Call{
							FunctionName: "foo",
						},
					},
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
			assert.Nil(t, p.File.Comments)
		})
	}
}
