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
						Expr: &ast.Identifier{Name: "print"},
					},
				},
			},
		},
		"try-on-1": {
			str: "try { print() } on SomeError {}",
			expected: &ast.ErrorScope{
				Statements: []ast.Node{
					&ast.Call{
						Expr: &ast.Identifier{Name: "print"},
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
						Expr: &ast.Identifier{Name: "print"},
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
								Expr: &ast.Identifier{Name: "foo"},
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
						Expr: &ast.Identifier{Name: "print"},
					},
				},
				On: []*ast.On{
					{
						Type: types.NewUnresolvedInterface("SomethingElse"),
						Statements: []ast.Node{
							&ast.Call{
								Expr: &ast.Identifier{Name: "foo"},
							},
						},
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
		"try-finally": {
			str: "try { print() } finally { foo() }",
			expected: &ast.ErrorScope{
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
	} {
		t.Run(testName, func(t *testing.T) {
			str := fmt.Sprintf("func main() { %s }", test.str)
			p := parser.NewParser()
			p.ParseString(str, "a.ok")

			assertEqualErrors(t, test.errs, p.Errors())
			asttest.AssertEqual(t, map[string]*ast.Func{
				"main": newFunc(test.expected),
			}, p.Funcs())
			assert.Nil(t, p.Comments())
		})
	}
}
