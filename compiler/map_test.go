package compiler_test

import (
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/compiler"
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
				Kind: "{}number",
			},
			expected: []vm.Instruction{
				// alloc
				&vm.Assign{
					VariableName: "1",
					Value:        asttest.NewLiteralNumber("0"),
				},
				&vm.MapAlloc{
					Kind:   "{}any",
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
				&vm.Assign{
					VariableName: "1",
					Value:        asttest.NewLiteralNumber("3"),
				},
				&vm.MapAlloc{
					Kind:   "{}any",
					Size:   "1",
					Result: "2",
				},

				// "a": 2
				&vm.Assign{
					VariableName: "3",
					Value:        asttest.NewLiteralString("a"),
				},
				&vm.Assign{
					VariableName: "4",
					Value:        asttest.NewLiteralNumber("2"),
				},
				&vm.MapSet{
					Map:   "2",
					Key:   "3",
					Value: "4",
				},

				// "b": 5
				&vm.Assign{
					VariableName: "5",
					Value:        asttest.NewLiteralString("b"),
				},
				&vm.Assign{
					VariableName: "6",
					Value:        asttest.NewLiteralNumber("5"),
				},
				&vm.MapSet{
					Map:   "2",
					Key:   "5",
					Value: "6",
				},

				// "c": 13
				&vm.Assign{
					VariableName: "7",
					Value:        asttest.NewLiteralString("c"),
				},
				&vm.Assign{
					VariableName: "8",
					Value:        asttest.NewLiteralNumber("13"),
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
				&vm.Assign{
					VariableName: "1",
					Value:        asttest.NewLiteralNumber("1"),
				},
				&vm.MapAlloc{
					Kind:   "{}any",
					Size:   "1",
					Result: "2",
				},

				// "b": 123
				&vm.Assign{
					VariableName: "3",
					Value:        asttest.NewLiteralString("b"),
				},
				&vm.Assign{
					VariableName: "4",
					Value:        asttest.NewLiteralNumber("123"),
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
			}, &compiler.Compiled{})
			if test.err != nil {
				assert.EqualError(t, err, test.err.Error())
			} else {
				require.NoError(t, err)
				assert.Equal(t, test.expected, compiledFunc.Instructions)
			}
		})
	}
}
