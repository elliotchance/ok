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
		"array-number-index": {
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
		"map-number-key": {
			nodes: []ast.Node{
				ast.NewBinary(
					&ast.Identifier{Name: "foo"},
					lexer.TokenAssign,
					&ast.Map{
						Elements: []*ast.KeyValue{
							{
								Key:   ast.NewLiteralString("a"),
								Value: ast.NewLiteralNumber("123"),
							},
							{
								Key:   ast.NewLiteralString("b"),
								Value: ast.NewLiteralNumber("456"),
							},
						},
					},
				),
				&ast.Key{
					Expr: &ast.Identifier{Name: "foo"},
					Key:  ast.NewLiteralString("b"),
				},
			},
			expected: []instruction.Instruction{
				// alloc
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("2"),
				},
				&instruction.MapAllocNumber{
					Size:   "1",
					Result: "2",
				},

				// "a": 123
				&instruction.Assign{
					VariableName: "3",
					Value:        ast.NewLiteralString("a"),
				},
				&instruction.Assign{
					VariableName: "4",
					Value:        ast.NewLiteralNumber("123"),
				},
				&instruction.MapSet{
					Map:   "2",
					Key:   "3",
					Value: "4",
				},

				// "b": 456
				&instruction.Assign{
					VariableName: "5",
					Value:        ast.NewLiteralString("b"),
				},
				&instruction.Assign{
					VariableName: "6",
					Value:        ast.NewLiteralNumber("456"),
				},
				&instruction.MapSet{
					Map:   "2",
					Key:   "5",
					Value: "6",
				},

				// assign foo
				&instruction.Assign{
					VariableName: "foo",
					Register:     "2",
				},

				// get "b"
				&instruction.Assign{
					VariableName: "7",
					Value:        ast.NewLiteralString("b"),
				},
				&instruction.MapGet{
					Map:    "foo",
					Key:    "7",
					Result: "8",
				},
			},
		},
		"assign-array-number": {
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
		"assign-map-number": {
			nodes: []ast.Node{
				ast.NewBinary(
					&ast.Identifier{Name: "foo"},
					lexer.TokenAssign,
					&ast.Map{
						Elements: []*ast.KeyValue{
							{
								Key:   ast.NewLiteralString("a"),
								Value: ast.NewLiteralNumber("123"),
							},
							{
								Key:   ast.NewLiteralString("b"),
								Value: ast.NewLiteralNumber("456"),
							},
						},
					},
				),
				ast.NewBinary(
					&ast.Key{
						Expr: &ast.Identifier{Name: "foo"},
						Key:  ast.NewLiteralString("b"),
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
				&instruction.MapAllocNumber{
					Size:   "1",
					Result: "2",
				},

				// "a": 123
				&instruction.Assign{
					VariableName: "3",
					Value:        ast.NewLiteralString("a"),
				},
				&instruction.Assign{
					VariableName: "4",
					Value:        ast.NewLiteralNumber("123"),
				},
				&instruction.MapSet{
					Map:   "2",
					Key:   "3",
					Value: "4",
				},

				// "b": 456
				&instruction.Assign{
					VariableName: "5",
					Value:        ast.NewLiteralString("b"),
				},
				&instruction.Assign{
					VariableName: "6",
					Value:        ast.NewLiteralNumber("456"),
				},
				&instruction.MapSet{
					Map:   "2",
					Key:   "5",
					Value: "6",
				},

				// assign foo
				&instruction.Assign{
					VariableName: "foo",
					Register:     "2",
				},

				// foo["b"] = 2
				&instruction.Assign{
					VariableName: "7",
					Value:        ast.NewLiteralNumber("2"),
				},
				&instruction.Assign{
					VariableName: "8",
					Value:        ast.NewLiteralString("b"),
				},
				&instruction.MapSet{
					Map:   "foo",
					Key:   "8",
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
