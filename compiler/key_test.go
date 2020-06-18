package compiler_test

import (
	"ok/ast"
	"ok/compiler"
	"ok/instruction"
	"ok/lexer"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestKey(t *testing.T) {
	for testName, test := range map[string]struct {
		nodes    []ast.Node
		expected []instruction.Instruction
		err      error
	}{
		"number-index": {
			nodes: []ast.Node{
				ast.NewBinary(
					&ast.Identifier{Name: "foo"},
					lexer.TokenAssign,
					&ast.Array{
						Elements: []ast.Node{
							ast.NewLiteralNumber("123"),
							ast.NewLiteralNumber("456"),
						},
					},
				),
				&ast.Key{
					Expr: &ast.Identifier{Name: "foo"},
					Key:  ast.NewLiteralNumber("1"),
				},
			},
			expected: []instruction.Instruction{
				// alloc
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("2"),
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
					Value:        ast.NewLiteralNumber("123"),
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
					Value:        ast.NewLiteralNumber("456"),
				},
				&instruction.ArraySet{
					Array: "2",
					Index: "5",
					Value: "6",
				},

				// assign foo
				&instruction.Assign{
					VariableName: "foo",
					Register:     "2",
				},

				// get 1
				&instruction.Assign{
					VariableName: "7",
					Value:        ast.NewLiteralNumber("1"),
				},
				&instruction.ArrayGet{
					Array:  "foo",
					Index:  "7",
					Result: "8",
				},
			},
		},
		"assign-number-index": {
			nodes: []ast.Node{
				ast.NewBinary(
					&ast.Identifier{Name: "foo"},
					lexer.TokenAssign,
					&ast.Array{
						Elements: []ast.Node{
							ast.NewLiteralNumber("123"),
							ast.NewLiteralNumber("456"),
						},
					},
				),
				ast.NewBinary(
					&ast.Key{
						Expr: &ast.Identifier{Name: "foo"},
						Key:  ast.NewLiteralNumber("1"),
					},
					lexer.TokenAssign,
					ast.NewLiteralNumber("2"),
				),
			},
			expected: []instruction.Instruction{
				// alloc
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("2"),
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
					Value:        ast.NewLiteralNumber("123"),
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
					Value:        ast.NewLiteralNumber("456"),
				},
				&instruction.ArraySet{
					Array: "2",
					Index: "5",
					Value: "6",
				},

				// assign foo
				&instruction.Assign{
					VariableName: "foo",
					Register:     "2",
				},

				// foo[1] = 2
				&instruction.Assign{
					VariableName: "7",
					Value:        ast.NewLiteralNumber("2"),
				},
				&instruction.Assign{
					VariableName: "8",
					Value:        ast.NewLiteralNumber("1"),
				},
				&instruction.ArraySet{
					Array: "foo",
					Index: "8",
					Value: "7",
				},
			},
		},
	} {
		t.Run(testName, func(t *testing.T) {
			compiledFunc, err := compiler.CompileFunc(&ast.Func{
				Statements: test.nodes,
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
