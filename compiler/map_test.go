package compiler_test

import (
	"ok/ast"
	"ok/compiler"
	"ok/vm"
	"testing"

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
				Kind: "{}number",
			},
			expected: []vm.Instruction{
				// alloc
				&vm.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("0"),
				},
				&vm.MapAllocNumber{
					Size:   "1",
					Result: "2",
				},
			},
		},
		"implicit-numbers": {
			node: &ast.Map{
				Elements: []*ast.KeyValue{
					{
						Key:   ast.NewLiteralString("a"),
						Value: ast.NewLiteralNumber("2"),
					},
					{
						Key:   ast.NewLiteralString("b"),
						Value: ast.NewLiteralNumber("5"),
					},
					{
						Key:   ast.NewLiteralString("c"),
						Value: ast.NewLiteralNumber("13"),
					},
				},
			},
			expected: []vm.Instruction{
				// alloc
				&vm.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("3"),
				},
				&vm.MapAllocNumber{
					Size:   "1",
					Result: "2",
				},

				// "a": 2
				&vm.Assign{
					VariableName: "3",
					Value:        ast.NewLiteralString("a"),
				},
				&vm.Assign{
					VariableName: "4",
					Value:        ast.NewLiteralNumber("2"),
				},
				&vm.MapSet{
					Map:   "2",
					Key:   "3",
					Value: "4",
				},

				// "b": 5
				&vm.Assign{
					VariableName: "5",
					Value:        ast.NewLiteralString("b"),
				},
				&vm.Assign{
					VariableName: "6",
					Value:        ast.NewLiteralNumber("5"),
				},
				&vm.MapSet{
					Map:   "2",
					Key:   "5",
					Value: "6",
				},

				// "c": 13
				&vm.Assign{
					VariableName: "7",
					Value:        ast.NewLiteralString("c"),
				},
				&vm.Assign{
					VariableName: "8",
					Value:        ast.NewLiteralNumber("13"),
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
								Key:   ast.NewLiteralString("b"),
								Value: ast.NewLiteralNumber("123"),
							},
						},
					},
				},
			},
			expected: []vm.Instruction{
				// alloc
				&vm.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("1"),
				},
				&vm.MapAllocNumber{
					Size:   "1",
					Result: "2",
				},

				// "b": 123
				&vm.Assign{
					VariableName: "3",
					Value:        ast.NewLiteralString("b"),
				},
				&vm.Assign{
					VariableName: "4",
					Value:        ast.NewLiteralNumber("123"),
				},
				&vm.MapSet{
					Map:   "2",
					Key:   "3",
					Value: "4",
				},

				// assign a
				&vm.Assign{
					VariableName: "a",
					Register:     "2",
				},
			},
		},
	} {
		t.Run(testName, func(t *testing.T) {
			compiledFunc, err := compiler.CompileFunc(&ast.Func{
				Statements: []ast.Node{
					test.node,
				},
			}, nil)
			if test.err != nil {
				assert.EqualError(t, err, test.err.Error())
			} else {
				require.NoError(t, err)
				assert.Equal(t, test.expected, compiledFunc.Instructions)
			}
		})
	}
}
