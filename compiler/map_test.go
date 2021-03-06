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
					Symbol: "0",
				},
				&vm.MapAlloc{
					Kind:   "2",
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
					Symbol: "0",
				},
				&vm.MapAlloc{
					Kind:   "3",
					Size:   "1",
					Result: "2",
				},

				// "a": 2
				&vm.AssignSymbol{
					Result: "3",
					Symbol: "1",
				},
				&vm.AssignSymbol{
					Result: "4",
					Symbol: "2",
				},
				&vm.MapSet{
					Map:   "2",
					Key:   "3",
					Value: "4",
				},

				// "b": 5
				&vm.AssignSymbol{
					Result: "5",
					Symbol: "3",
				},
				&vm.AssignSymbol{
					Result: "6",
					Symbol: "4",
				},
				&vm.MapSet{
					Map:   "2",
					Key:   "5",
					Value: "6",
				},

				// "c": 13
				&vm.AssignSymbol{
					Result: "7",
					Symbol: "5",
				},
				&vm.AssignSymbol{
					Result: "8",
					Symbol: "6",
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
					Symbol: "0",
				},
				&vm.MapAlloc{
					Kind:   "3",
					Size:   "1",
					Result: "2",
				},

				// "b": 123
				&vm.AssignSymbol{
					Result: "3",
					Symbol: "1",
				},
				&vm.AssignSymbol{
					Result: "4",
					Symbol: "2",
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
				Types:   types.Registry{},
				Symbols: map[vm.SymbolRegister]*vm.Symbol{},
			}, nil, nil, nil, nil)
			if test.err != nil {
				assert.EqualError(t, err, test.err.Error())
			} else {
				require.NoError(t, err)
				assert.Equal(t, test.expected, compiledFunc.Instructions.Instructions)
			}
		})
	}
}
