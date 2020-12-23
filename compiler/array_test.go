package compiler_test

import (
	"errors"
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/compiler"
	"github.com/elliotchance/ok/types"
	"github.com/elliotchance/ok/vm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestArray(t *testing.T) {
	for testName, test := range map[string]struct {
		node     ast.Node
		expected []vm.Instruction
		err      error
	}{
		"unknown-array-empty": {
			node: &ast.Array{},
			err:  errors.New(" empty array needs to specify a type"),
		},
		"number-empty": {
			node: &ast.Array{
				Kind: types.NumberArray,
			},
			expected: []vm.Instruction{
				// alloc
				&vm.AssignSymbol{
					Result: "1",
					Symbol: "number0",
				},
				&vm.ArrayAlloc{
					Size:   "1",
					Result: "2",
					Kind:   "[]number",
				},
			},
		},
		"implicit-numbers": {
			node: &ast.Array{
				Elements: []ast.Node{
					asttest.NewLiteralNumber("2"),
					asttest.NewLiteralNumber("5"),
					asttest.NewLiteralNumber("13"),
				},
			},
			expected: []vm.Instruction{
				// alloc
				&vm.AssignSymbol{
					Result: "1",
					Symbol: "number3",
				},
				&vm.ArrayAlloc{
					Size:   "1",
					Result: "2",
					Kind:   "[]number",
				},

				// set 0
				&vm.AssignSymbol{
					Result: "3",
					Symbol: "number0",
				},
				&vm.AssignSymbol{
					Result: "4",
					Symbol: "number2",
				},
				&vm.ArraySet{
					Array: "2",
					Index: "3",
					Value: "4",
				},

				// set 1
				&vm.AssignSymbol{
					Result: "5",
					Symbol: "number1",
				},
				&vm.AssignSymbol{
					Result: "6",
					Symbol: "number5",
				},
				&vm.ArraySet{
					Array: "2",
					Index: "5",
					Value: "6",
				},

				// set 2
				&vm.AssignSymbol{
					Result: "7",
					Symbol: "number2",
				},
				&vm.AssignSymbol{
					Result: "8",
					Symbol: "number13",
				},
				&vm.ArraySet{
					Array: "2",
					Index: "7",
					Value: "8",
				},
			},
		},
		"assign-implicit-numbers": {
			node: &ast.Assign{
				Lefts: []ast.Node{
					&ast.Identifier{Name: "a"},
				},
				Rights: []ast.Node{
					&ast.Array{
						Elements: []ast.Node{
							asttest.NewLiteralNumber("1"),
						},
					},
				},
			},
			expected: []vm.Instruction{
				// alloc
				&vm.AssignSymbol{
					Result: "1",
					Symbol: "number1",
				},
				&vm.ArrayAlloc{
					Size:   "1",
					Result: "2",
					Kind:   "[]number",
				},

				// set 0
				&vm.AssignSymbol{
					Result: "3",
					Symbol: "number0",
				},
				&vm.AssignSymbol{
					Result: "4",
					Symbol: "number1",
				},
				&vm.ArraySet{
					Array: "2",
					Index: "3",
					Value: "4",
				},

				// assign a
				&vm.Assign{
					Result:   "a",
					Register: "2",
				},
			},
		},
	} {
		t.Run(testName, func(t *testing.T) {
			compiledFunc, err := compiler.CompileFunc(&ast.Func{
				Statements: []ast.Node{
					test.node,
				},
			}, &vm.File{
				Types:   map[vm.TypeRegister]*types.Type{},
				Symbols: map[vm.SymbolRegister]*vm.Symbol{},
			}, nil)
			if test.err != nil {
				assert.EqualError(t, err, test.err.Error())
			} else {
				require.NoError(t, err)
				assert.Equal(t, test.expected, compiledFunc.Instructions.Instructions)
			}
		})
	}
}
