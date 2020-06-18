package compiler_test

import (
	"errors"
	"ok/ast"
	"ok/compiler"
	"ok/instruction"
	"ok/lexer"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestArray(t *testing.T) {
	for testName, test := range map[string]struct {
		node     ast.Node
		expected []instruction.Instruction
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
			expected: []instruction.Instruction{
				// alloc
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("0"),
				},
				&instruction.ArrayAllocNumber{
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
			expected: []instruction.Instruction{
				// alloc
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("3"),
				},
				&instruction.ArrayAllocNumber{
					Size:   "1",
					Result: "2",
				},

				// set 0
				&instruction.Assign{
					VariableName: "3",
					Value:        ast.NewLiteralNumber("0"),
				},
				&instruction.Assign{
					VariableName: "4",
					Value:        ast.NewLiteralNumber("2"),
				},
				&instruction.ArraySet{
					Array: "2",
					Index: "3",
					Value: "4",
				},

				// set 1
				&instruction.Assign{
					VariableName: "5",
					Value:        ast.NewLiteralNumber("1"),
				},
				&instruction.Assign{
					VariableName: "6",
					Value:        ast.NewLiteralNumber("5"),
				},
				&instruction.ArraySet{
					Array: "2",
					Index: "5",
					Value: "6",
				},

				// set 2
				&instruction.Assign{
					VariableName: "7",
					Value:        ast.NewLiteralNumber("2"),
				},
				&instruction.Assign{
					VariableName: "8",
					Value:        ast.NewLiteralNumber("13"),
				},
				&instruction.ArraySet{
					Array: "2",
					Index: "7",
					Value: "8",
				},
			},
		},
		"assign-implicit-numbers": {
			node: ast.NewBinary(
				&ast.Identifier{Name: "a"},
				lexer.TokenAssign,
				&ast.Array{
					Elements: []ast.Node{
						ast.NewLiteralNumber("1"),
					},
				},
			),
			expected: []instruction.Instruction{
				// alloc
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("1"),
				},
				&instruction.ArrayAllocNumber{
					Size:   "1",
					Result: "2",
				},

				// set 0
				&instruction.Assign{
					VariableName: "3",
					Value:        ast.NewLiteralNumber("0"),
				},
				&instruction.Assign{
					VariableName: "4",
					Value:        ast.NewLiteralNumber("1"),
				},
				&instruction.ArraySet{
					Array: "2",
					Index: "3",
					Value: "4",
				},

				// assign a
				&instruction.Assign{
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
			})
			if test.err != nil {
				assert.EqualError(t, err, test.err.Error())
			} else {
				require.NoError(t, err)
				assert.Equal(t, test.expected, compiledFunc.Instructions)
			}
		})
	}
}
