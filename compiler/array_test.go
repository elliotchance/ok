package compiler_test

import (
	"errors"
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/compiler"
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
			err:  errors.New("empty array needs to specify a type"),
		},
		"number-empty": {
			node: &ast.Array{
				Kind: "[]number",
			},
			expected: []vm.Instruction{
				// alloc
				&vm.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("0"),
				},
				&vm.ArrayAllocNumber{
					Size:   "1",
					Result: "2",
				},
			},
		},
		"implicit-numbers": {
			node: &ast.Array{
				Elements: []ast.Node{
					ast.NewLiteralNumber("2"),
					ast.NewLiteralNumber("5"),
					ast.NewLiteralNumber("13"),
				},
			},
			expected: []vm.Instruction{
				// alloc
				&vm.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("3"),
				},
				&vm.ArrayAllocNumber{
					Size:   "1",
					Result: "2",
				},

				// set 0
				&vm.Assign{
					VariableName: "3",
					Value:        ast.NewLiteralNumber("0"),
				},
				&vm.Assign{
					VariableName: "4",
					Value:        ast.NewLiteralNumber("2"),
				},
				&vm.ArraySet{
					Array: "2",
					Index: "3",
					Value: "4",
				},

				// set 1
				&vm.Assign{
					VariableName: "5",
					Value:        ast.NewLiteralNumber("1"),
				},
				&vm.Assign{
					VariableName: "6",
					Value:        ast.NewLiteralNumber("5"),
				},
				&vm.ArraySet{
					Array: "2",
					Index: "5",
					Value: "6",
				},

				// set 2
				&vm.Assign{
					VariableName: "7",
					Value:        ast.NewLiteralNumber("2"),
				},
				&vm.Assign{
					VariableName: "8",
					Value:        ast.NewLiteralNumber("13"),
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
							ast.NewLiteralNumber("1"),
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
				&vm.ArrayAllocNumber{
					Size:   "1",
					Result: "2",
				},

				// set 0
				&vm.Assign{
					VariableName: "3",
					Value:        ast.NewLiteralNumber("0"),
				},
				&vm.Assign{
					VariableName: "4",
					Value:        ast.NewLiteralNumber("1"),
				},
				&vm.ArraySet{
					Array: "2",
					Index: "3",
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
