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

func TestMap(t *testing.T) {
	for testName, test := range map[string]struct {
		node     ast.Node
		expected []instruction.Instruction
		err      error
	}{
		"number-empty": {
			node: &ast.Map{
				Kind: "{}number",
			},
			expected: []instruction.Instruction{
				// alloc
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("0"),
				},
				&instruction.MapAllocNumber{
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
			expected: []instruction.Instruction{
				// alloc
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("3"),
				},
				&instruction.MapAllocNumber{
					Size:   "1",
					Result: "2",
				},

				// "a": 2
				&instruction.Assign{
					VariableName: "3",
					Value:        ast.NewLiteralString("a"),
				},
				&instruction.Assign{
					VariableName: "4",
					Value:        ast.NewLiteralNumber("2"),
				},
				&instruction.MapSet{
					Map:   "2",
					Key:   "3",
					Value: "4",
				},

				// "b": 5
				&instruction.Assign{
					VariableName: "5",
					Value:        ast.NewLiteralString("b"),
				},
				&instruction.Assign{
					VariableName: "6",
					Value:        ast.NewLiteralNumber("5"),
				},
				&instruction.MapSet{
					Map:   "2",
					Key:   "5",
					Value: "6",
				},

				// "c": 13
				&instruction.Assign{
					VariableName: "7",
					Value:        ast.NewLiteralString("c"),
				},
				&instruction.Assign{
					VariableName: "8",
					Value:        ast.NewLiteralNumber("13"),
				},
				&instruction.MapSet{
					Map:   "2",
					Key:   "7",
					Value: "8",
				},
			},
		},
		"assign-implicit-numbers": {
			node: ast.NewBinary(
				&ast.Identifier{Name: "a"},
				lexer.TokenAssign,
				&ast.Map{
					Elements: []*ast.KeyValue{
						{
							Key:   ast.NewLiteralString("b"),
							Value: ast.NewLiteralNumber("123"),
						},
					},
				},
			),
			expected: []instruction.Instruction{
				// alloc
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("1"),
				},
				&instruction.MapAllocNumber{
					Size:   "1",
					Result: "2",
				},

				// "b": 123
				&instruction.Assign{
					VariableName: "3",
					Value:        ast.NewLiteralString("b"),
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
