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

func TestFor(t *testing.T) {
	for testName, test := range map[string]struct {
		nodes    []ast.Node
		expected []instruction.Instruction
		err      error
	}{
		"for-in-array-1": {
			nodes: []ast.Node{
				ast.NewBinary(
					&ast.Identifier{Name: "foo"},
					lexer.TokenAssign,
					ast.NewArrayNumbers([]string{"1.5", "2.3"}),
				),
				&ast.For{
					Condition: &ast.In{
						Key:  "k",
						Expr: &ast.Identifier{Name: "foo"},
					},
				},
			},
			expected: []instruction.Instruction{
				// alloc array
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("2"),
				},
				&instruction.ArrayAllocNumber{
					Size:   "1",
					Result: "2",
				},

				// set 2 elements
				&instruction.Assign{
					VariableName: "3",
					Value:        ast.NewLiteralNumber("0"),
				},
				&instruction.Assign{
					VariableName: "4",
					Value:        ast.NewLiteralNumber("1.5"),
				},
				&instruction.ArraySet{
					Array: "2",
					Index: "3",
					Value: "4",
				},
				&instruction.Assign{
					VariableName: "5",
					Value:        ast.NewLiteralNumber("1"),
				},
				&instruction.Assign{
					VariableName: "6",
					Value:        ast.NewLiteralNumber("2.3"),
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

				// for in
				&instruction.Assign{
					VariableName: "7",
					Value:        ast.NewLiteralNumber("0"),
				},
				&instruction.NextArray{
					Array:       "foo",
					Cursor:      "7",
					KeyResult:   "k",
					ValueResult: "",
					Result:      "8",
				},
				&instruction.JumpUnless{
					Condition: "8",
					To:        12,
				},
				&instruction.Jump{
					To: 9,
				},
			},
		},
		"for-in-array-2": {
			nodes: []ast.Node{
				ast.NewBinary(
					&ast.Identifier{Name: "foo"},
					lexer.TokenAssign,
					ast.NewArrayNumbers([]string{"1.5", "2.3"}),
				),
				&ast.For{
					Condition: &ast.In{
						Key:   "k",
						Value: "v",
						Expr:  &ast.Identifier{Name: "foo"},
					},
				},
			},
			expected: []instruction.Instruction{
				// alloc array
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("2"),
				},
				&instruction.ArrayAllocNumber{
					Size:   "1",
					Result: "2",
				},

				// set 2 elements
				&instruction.Assign{
					VariableName: "3",
					Value:        ast.NewLiteralNumber("0"),
				},
				&instruction.Assign{
					VariableName: "4",
					Value:        ast.NewLiteralNumber("1.5"),
				},
				&instruction.ArraySet{
					Array: "2",
					Index: "3",
					Value: "4",
				},
				&instruction.Assign{
					VariableName: "5",
					Value:        ast.NewLiteralNumber("1"),
				},
				&instruction.Assign{
					VariableName: "6",
					Value:        ast.NewLiteralNumber("2.3"),
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

				// for in
				&instruction.Assign{
					VariableName: "7",
					Value:        ast.NewLiteralNumber("0"),
				},
				&instruction.NextArray{
					Array:       "foo",
					Cursor:      "7",
					KeyResult:   "k",
					ValueResult: "v",
					Result:      "8",
				},
				&instruction.JumpUnless{
					Condition: "8",
					To:        12,
				},
				&instruction.Jump{
					To: 9,
				},
			},
		},
		"for-in-map-1": {
			nodes: []ast.Node{
				ast.NewBinary(
					&ast.Identifier{Name: "foo"},
					lexer.TokenAssign,
					ast.NewMapNumbers(map[string]string{"foo": "1.5", "bar": "2.3"}),
				),
				&ast.For{
					Condition: &ast.In{
						Key:  "k",
						Expr: &ast.Identifier{Name: "foo"},
					},
				},
			},
			expected: []instruction.Instruction{
				// alloc array
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("2"),
				},
				&instruction.MapAllocNumber{
					Size:   "1",
					Result: "2",
				},

				// set 2 elements
				&instruction.Assign{
					VariableName: "3",
					Value:        ast.NewLiteralString("bar"),
				},
				&instruction.Assign{
					VariableName: "4",
					Value:        ast.NewLiteralNumber("2.3"),
				},
				&instruction.MapSet{
					Map:   "2",
					Key:   "3",
					Value: "4",
				},
				&instruction.Assign{
					VariableName: "5",
					Value:        ast.NewLiteralString("foo"),
				},
				&instruction.Assign{
					VariableName: "6",
					Value:        ast.NewLiteralNumber("1.5"),
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

				// for in
				&instruction.Assign{
					VariableName: "7",
					Value:        ast.NewLiteralNumber("0"),
				},
				&instruction.NextMap{
					Map:         "foo",
					Cursor:      "7",
					KeyResult:   "k",
					ValueResult: "",
					Result:      "8",
				},
				&instruction.JumpUnless{
					Condition: "8",
					To:        12,
				},
				&instruction.Jump{
					To: 9,
				},
			},
		},
		"for-in-map-2": {
			nodes: []ast.Node{
				ast.NewBinary(
					&ast.Identifier{Name: "foo"},
					lexer.TokenAssign,
					ast.NewMapNumbers(map[string]string{"foo": "1.5", "bar": "2.3"}),
				),
				&ast.For{
					Condition: &ast.In{
						Key:   "k",
						Value: "v",
						Expr:  &ast.Identifier{Name: "foo"},
					},
				},
			},
			expected: []instruction.Instruction{
				// alloc array
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("2"),
				},
				&instruction.MapAllocNumber{
					Size:   "1",
					Result: "2",
				},

				// set 2 elements
				&instruction.Assign{
					VariableName: "3",
					Value:        ast.NewLiteralString("bar"),
				},
				&instruction.Assign{
					VariableName: "4",
					Value:        ast.NewLiteralNumber("2.3"),
				},
				&instruction.MapSet{
					Map:   "2",
					Key:   "3",
					Value: "4",
				},
				&instruction.Assign{
					VariableName: "5",
					Value:        ast.NewLiteralString("foo"),
				},
				&instruction.Assign{
					VariableName: "6",
					Value:        ast.NewLiteralNumber("1.5"),
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

				// for in
				&instruction.Assign{
					VariableName: "7",
					Value:        ast.NewLiteralNumber("0"),
				},
				&instruction.NextMap{
					Map:         "foo",
					Cursor:      "7",
					KeyResult:   "k",
					ValueResult: "v",
					Result:      "8",
				},
				&instruction.JumpUnless{
					Condition: "8",
					To:        12,
				},
				&instruction.Jump{
					To: 9,
				},
			},
		},
	} {
		t.Run(testName, func(t *testing.T) {
			compiledFunc, err := compiler.CompileFunc(newFunc(test.nodes...))
			if test.err != nil {
				assert.EqualError(t, err, test.err.Error())
			} else {
				require.NoError(t, err)
				assert.Equal(t, test.expected, compiledFunc.Instructions)
			}
		})
	}
}
