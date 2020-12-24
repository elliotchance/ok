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

func TestMap(t *testing.T) {
	for testName, test := range map[string]struct {
		node     ast.Node
		expected []vm.Instruction
		err      error
	}{
		"number-empty": {
			node: &ast.Map{
				Kind: types.NumberMap,
			},
			expected: []vm.Instruction{
				// alloc
				&vm.AssignSymbol{
					Result: "1",
					Symbol: "number0",
				},
				&vm.MapAlloc{
					Kind:   "{}number",
					Size:   "1",
					Result: "2",
				},
			},
		},
		"implicit-numbers": {
			node: &ast.Map{
				Elements: []*ast.KeyValue{
					{
						Key:   asttest.NewLiteralString("a"),
						Value: asttest.NewLiteralNumber("2"),
					},
					{
						Key:   asttest.NewLiteralString("b"),
						Value: asttest.NewLiteralNumber("5"),
					},
					{
						Key:   asttest.NewLiteralString("c"),
						Value: asttest.NewLiteralNumber("13"),
					},
				},
			},
			expected: []vm.Instruction{
				// alloc
				&vm.AssignSymbol{
					Result: "1",
					Symbol: "number3",
				},
				&vm.MapAlloc{
					Kind:   "{}number",
					Size:   "1",
					Result: "2",
				},

				// "a": 2
				&vm.AssignSymbol{
					Result: "3",
					Symbol: "stringa",
				},
				&vm.AssignSymbol{
					Result: "4",
					Symbol: "number2",
				},
				&vm.MapSet{
					Map:   "2",
					Key:   "3",
					Value: "4",
				},

				// "b": 5
				&vm.AssignSymbol{
					Result: "5",
					Symbol: "stringb",
				},
				&vm.AssignSymbol{
					Result: "6",
					Symbol: "number5",
				},
				&vm.MapSet{
					Map:   "2",
					Key:   "5",
					Value: "6",
				},

				// "c": 13
				&vm.AssignSymbol{
					Result: "7",
					Symbol: "stringc",
				},
				&vm.AssignSymbol{
					Result: "8",
					Symbol: "number13",
				},
				&vm.MapSet{
					Map:   "2",
					Key:   "7",
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
					&ast.Map{
						Elements: []*ast.KeyValue{
							{
								Key:   asttest.NewLiteralString("b"),
								Value: asttest.NewLiteralNumber("123"),
							},
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
				&vm.MapAlloc{
					Kind:   "{}number",
					Size:   "1",
					Result: "2",
				},

				// "b": 123
				&vm.AssignSymbol{
					Result: "3",
					Symbol: "stringb",
				},
				&vm.AssignSymbol{
					Result: "4",
					Symbol: "number123",
				},
				&vm.MapSet{
					Map:   "2",
					Key:   "3",
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
