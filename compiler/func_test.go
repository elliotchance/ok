package compiler_test

import (
	"errors"
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/compiler"
	"github.com/elliotchance/ok/parser"
	"github.com/elliotchance/ok/types"
	"github.com/elliotchance/ok/vm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func parseFunc(s string) *ast.Func {
	p := parser.NewParser(0)
	p.ParseString(s, "a.ok")
	return p.Funcs()["foo"]
}

func TestFunc(t *testing.T) {
	for testName, test := range map[string]struct {
		fn       *ast.Func
		expected []vm.Instruction
		err      error
	}{
		"no-statements": {
			fn: &ast.Func{},
		},
		"one-statement-no-args": {
			fn: &ast.Func{
				Statements: []ast.Node{
					&ast.Call{
						Expr: &ast.Identifier{Name: "print"},
					},
				},
			},
			expected: []vm.Instruction{
				&vm.Print{},
			},
		},
		"two-statements-with-args": {
			fn: &ast.Func{
				Statements: []ast.Node{
					&ast.Call{
						Expr: &ast.Identifier{Name: "print"},
					},
					&ast.Call{
						Expr: &ast.Identifier{Name: "print"},
						Arguments: []ast.Node{
							asttest.NewLiteralString("hello"),
						},
					},
				},
			},
			expected: []vm.Instruction{
				&vm.Print{},
				&vm.Assign{
					VariableName: "1",
					Value:        asttest.NewLiteralString("hello"),
				},
				&vm.Print{
					Arguments: []vm.Register{"1"},
				},
			},
		},
		"oos-exists-arg": {
			fn: parseFunc("func foo(baz number) { func bar() number { return ^baz } }"),
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "2",
					Value: &ast.Literal{
						Kind:  types.NewFunc(nil, []*types.Type{types.Number}),
						Value: "2",
					},
				},
				&vm.ParentScope{
					X: "2",
				},
				&vm.Assign{
					VariableName: "bar",
					Register:     "2",
				},
			},
		},
		"oos-exists-var": {
			fn: parseFunc("func foo() { baz = 0\n func bar() number { return ^baz } }"),
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value: &ast.Literal{
						Kind:  types.NewFunc(nil, []*types.Type{types.Number}),
						Value: "2",
					},
				},
				&vm.ParentScope{
					X: "1",
				},
				&vm.Assign{
					VariableName: "bar",
					Register:     "1",
				},
				&vm.Assign{
					VariableName: "2",
					Value: &ast.Literal{
						Kind:  types.Number,
						Value: "0",
						Pos:   "a.ok:1:20",
					},
				},
				&vm.Assign{
					VariableName: "baz",
					Register:     "2",
				},
			},
		},
		"oos-illegal": {
			fn:  parseFunc("func foo() { return ^bar }"),
			err: errors.New("bar does not exist in the parent scope"),
		},
		"oos-nonexistent": {
			fn:  parseFunc("func foo() { func bar() number { return ^baz } }"),
			err: errors.New("baz does not exist in the parent scope"),
		},
	} {
		t.Run(testName, func(t *testing.T) {
			compiledFunc, err := compiler.CompileFunc(test.fn,
				&vm.File{Funcs: map[string]*vm.CompiledFunc{}}, nil)
			if test.err != nil {
				assert.EqualError(t, err, test.err.Error())
			} else {
				require.NoError(t, err)
				assert.Equal(t, test.expected, compiledFunc.Instructions)
			}
		})
	}
}

func newFunc(statements ...ast.Node) *ast.Func {
	return &ast.Func{
		Statements: statements,
	}
}
