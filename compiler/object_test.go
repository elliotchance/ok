package compiler_test

import (
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/compiler"
	"github.com/elliotchance/ok/types"
	"github.com/elliotchance/ok/vm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestObject(t *testing.T) {
	for testName, test := range map[string]struct {
		node     *ast.Func
		expected []vm.Instruction
		err      error
	}{
		"empty-object": {
			node: &ast.Func{
				Name:    "Foo",
				Returns: []*types.Type{types.NewUnresolvedInterface("Foo")},
			},
			expected: []vm.Instruction{
				// return instance
				&vm.Return{
					Results: []vm.Register{vm.StateRegister},
				},
			},
		},
		"one-variable": {
			node: &ast.Func{
				Name:    "Foo",
				Returns: []*types.Type{types.NewUnresolvedInterface("Foo")},
				Statements: []ast.Node{
					&ast.Assign{
						Lefts: []ast.Node{
							&ast.Identifier{Name: "bar"},
						},
						Rights: []ast.Node{
							asttest.NewLiteralNumber("123"),
						},
					},
				},
			},
			expected: []vm.Instruction{
				// bar = 123
				&vm.AssignSymbol{
					Result: "1",
					Symbol: "number123",
				},
				&vm.Assign{
					Result:   "bar",
					Register: "1",
				},

				// return instance
				&vm.Return{
					Results: []vm.Register{vm.StateRegister},
				},
			},
		},
	} {
		t.Run(testName, func(t *testing.T) {
			compiledFunc, err := compiler.CompileFunc(test.node,
				&vm.File{
					Symbols: map[vm.SymbolRegister]*vm.Symbol{},
				}, nil, nil)
			if test.err != nil {
				assert.EqualError(t, err, test.err.Error())
			} else {
				require.NoError(t, err)
				assert.Equal(t, test.expected, compiledFunc.Instructions.Instructions)
			}
		})
	}
}
