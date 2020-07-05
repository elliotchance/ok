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

func TestKey(t *testing.T) {
	for testName, test := range map[string]struct {
		nodes    []ast.Node
		expected []vm.Instruction
		err      error
	}{
		"array-number-index": {
			nodes: []ast.Node{
				&ast.Assign{
					Lefts: []ast.Node{
						&ast.Identifier{Name: "foo"},
					},
					Rights: []ast.Node{
						&ast.Array{
							Elements: []ast.Node{
								asttest.NewLiteralNumber("123"),
								asttest.NewLiteralNumber("456"),
							},
						},
					},
				},
				&ast.Key{
					Expr: &ast.Identifier{Name: "foo"},
					Key:  asttest.NewLiteralNumber("1"),
				},
			},
			expected: []vm.Instruction{
				// alloc
				&vm.Assign{
					VariableName: "1",
					Value:        asttest.NewLiteralNumber("2"),
				},
				&vm.ArrayAlloc{
					Size:   "1",
					Result: "2",
					Kind:   "[]number",
				},

				// set 0
				&vm.Assign{
					VariableName: "3",
					Value:        asttest.NewLiteralNumber("0"),
				},
				&vm.Assign{
					VariableName: "4",
					Value:        asttest.NewLiteralNumber("123"),
				},
				&vm.ArraySet{
					Array: "2",
					Index: "3",
					Value: "4",
				},

				// set 1
				&vm.Assign{
					VariableName: "5",
					Value:        asttest.NewLiteralNumber("1"),
				},
				&vm.Assign{
					VariableName: "6",
					Value:        asttest.NewLiteralNumber("456"),
				},
				&vm.ArraySet{
					Array: "2",
					Index: "5",
					Value: "6",
				},

				// assign foo
				&vm.Assign{
					VariableName: "foo",
					Register:     "2",
				},

				// get 1
				&vm.Assign{
					VariableName: "7",
					Value:        asttest.NewLiteralNumber("1"),
				},
				&vm.ArrayGet{
					Array:  "foo",
					Index:  "7",
					Result: "8",
				},
			},
		},
		"map-number-key": {
			nodes: []ast.Node{
				&ast.Assign{
					Lefts: []ast.Node{
						&ast.Identifier{Name: "foo"},
					},
					Rights: []ast.Node{
						&ast.Map{
							Elements: []*ast.KeyValue{
								{
									Key:   asttest.NewLiteralString("a"),
									Value: asttest.NewLiteralNumber("123"),
								},
								{
									Key:   asttest.NewLiteralString("b"),
									Value: asttest.NewLiteralNumber("456"),
								},
							},
						},
					},
				},
				&ast.Key{
					Expr: &ast.Identifier{Name: "foo"},
					Key:  asttest.NewLiteralString("b"),
				},
			},
			expected: []vm.Instruction{
				// alloc
				&vm.Assign{
					VariableName: "1",
					Value:        asttest.NewLiteralNumber("2"),
				},
				&vm.MapAllocNumber{
					Size:   "1",
					Result: "2",
				},

				// "a": 123
				&vm.Assign{
					VariableName: "3",
					Value:        asttest.NewLiteralString("a"),
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

				// "b": 456
				&vm.Assign{
					VariableName: "5",
					Value:        asttest.NewLiteralString("b"),
				},
				&vm.Assign{
					VariableName: "6",
					Value:        asttest.NewLiteralNumber("456"),
				},
				&vm.MapSet{
					Map:   "2",
					Key:   "5",
					Value: "6",
				},

				// assign foo
				&vm.Assign{
					VariableName: "foo",
					Register:     "2",
				},

				// get "b"
				&vm.Assign{
					VariableName: "7",
					Value:        asttest.NewLiteralString("b"),
				},
				&vm.MapGet{
					Map:    "foo",
					Key:    "7",
					Result: "8",
				},
			},
		},
		"object-property": {
			nodes: []ast.Node{
				&ast.Assign{
					Lefts: []ast.Node{
						&ast.Identifier{Name: "foo"},
					},
					Rights: []ast.Node{
						// We can use a map here because constructors actually
						// return a map.
						&ast.Map{
							Elements: []*ast.KeyValue{
								{
									Key:   asttest.NewLiteralString("a"),
									Value: asttest.NewLiteralNumber("123"),
								},
								{
									Key:   asttest.NewLiteralString("b"),
									Value: asttest.NewLiteralNumber("456"),
								},
							},
						},
					},
				},
				&ast.Key{
					Expr: &ast.Identifier{Name: "foo"},
					Key:  asttest.NewLiteralString("b"),
				},
			},
			expected: []vm.Instruction{
				// alloc
				&vm.Assign{
					VariableName: "1",
					Value:        asttest.NewLiteralNumber("2"),
				},
				&vm.MapAllocNumber{
					Size:   "1",
					Result: "2",
				},

				// "a": 123
				&vm.Assign{
					VariableName: "3",
					Value:        asttest.NewLiteralString("a"),
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

				// "b": 456
				&vm.Assign{
					VariableName: "5",
					Value:        asttest.NewLiteralString("b"),
				},
				&vm.Assign{
					VariableName: "6",
					Value:        asttest.NewLiteralNumber("456"),
				},
				&vm.MapSet{
					Map:   "2",
					Key:   "5",
					Value: "6",
				},

				// assign foo
				&vm.Assign{
					VariableName: "foo",
					Register:     "2",
				},

				// get "b"
				&vm.Assign{
					VariableName: "7",
					Value:        asttest.NewLiteralString("b"),
				},
				&vm.MapGet{
					Map:    "foo",
					Key:    "7",
					Result: "8",
				},
			},
		},
		"assign-array-number": {
			nodes: []ast.Node{
				&ast.Assign{
					Lefts: []ast.Node{
						&ast.Identifier{Name: "foo"},
					},
					Rights: []ast.Node{
						&ast.Array{
							Elements: []ast.Node{
								asttest.NewLiteralNumber("123"),
								asttest.NewLiteralNumber("456"),
							},
						},
					},
				},
				&ast.Assign{
					Lefts: []ast.Node{
						&ast.Key{
							Expr: &ast.Identifier{Name: "foo"},
							Key:  asttest.NewLiteralNumber("1"),
						},
					},
					Rights: []ast.Node{
						asttest.NewLiteralNumber("2"),
					},
				},
			},
			expected: []vm.Instruction{
				// alloc
				&vm.Assign{
					VariableName: "1",
					Value:        asttest.NewLiteralNumber("2"),
				},
				&vm.ArrayAlloc{
					Size:   "1",
					Result: "2",
					Kind:   "[]number",
				},

				// set 0
				&vm.Assign{
					VariableName: "3",
					Value:        asttest.NewLiteralNumber("0"),
				},
				&vm.Assign{
					VariableName: "4",
					Value:        asttest.NewLiteralNumber("123"),
				},
				&vm.ArraySet{
					Array: "2",
					Index: "3",
					Value: "4",
				},

				// set 1
				&vm.Assign{
					VariableName: "5",
					Value:        asttest.NewLiteralNumber("1"),
				},
				&vm.Assign{
					VariableName: "6",
					Value:        asttest.NewLiteralNumber("456"),
				},
				&vm.ArraySet{
					Array: "2",
					Index: "5",
					Value: "6",
				},

				// assign foo
				&vm.Assign{
					VariableName: "foo",
					Register:     "2",
				},

				// foo[1] = 2
				&vm.Assign{
					VariableName: "7",
					Value:        asttest.NewLiteralNumber("2"),
				},
				&vm.Assign{
					VariableName: "8",
					Value:        asttest.NewLiteralNumber("1"),
				},
				&vm.ArraySet{
					Array: "foo",
					Index: "8",
					Value: "7",
				},
			},
		},
		"assign-map-number": {
			nodes: []ast.Node{
				&ast.Assign{
					Lefts: []ast.Node{
						&ast.Identifier{Name: "foo"},
					},
					Rights: []ast.Node{
						&ast.Map{
							Elements: []*ast.KeyValue{
								{
									Key:   asttest.NewLiteralString("a"),
									Value: asttest.NewLiteralNumber("123"),
								},
								{
									Key:   asttest.NewLiteralString("b"),
									Value: asttest.NewLiteralNumber("456"),
								},
							},
						},
					},
				},
				&ast.Assign{
					Lefts: []ast.Node{
						&ast.Key{
							Expr: &ast.Identifier{Name: "foo"},
							Key:  asttest.NewLiteralString("b"),
						},
					},
					Rights: []ast.Node{
						asttest.NewLiteralNumber("2"),
					},
				},
			},
			expected: []vm.Instruction{
				// alloc
				&vm.Assign{
					VariableName: "1",
					Value:        asttest.NewLiteralNumber("2"),
				},
				&vm.MapAllocNumber{
					Size:   "1",
					Result: "2",
				},

				// "a": 123
				&vm.Assign{
					VariableName: "3",
					Value:        asttest.NewLiteralString("a"),
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

				// "b": 456
				&vm.Assign{
					VariableName: "5",
					Value:        asttest.NewLiteralString("b"),
				},
				&vm.Assign{
					VariableName: "6",
					Value:        asttest.NewLiteralNumber("456"),
				},
				&vm.MapSet{
					Map:   "2",
					Key:   "5",
					Value: "6",
				},

				// assign foo
				&vm.Assign{
					VariableName: "foo",
					Register:     "2",
				},

				// foo["b"] = 2
				&vm.Assign{
					VariableName: "7",
					Value:        asttest.NewLiteralNumber("2"),
				},
				&vm.Assign{
					VariableName: "8",
					Value:        asttest.NewLiteralString("b"),
				},
				&vm.MapSet{
					Map:   "foo",
					Key:   "8",
					Value: "7",
				},
			},
		},
		"string-number-index": {
			nodes: []ast.Node{
				&ast.Assign{
					Lefts: []ast.Node{
						&ast.Identifier{Name: "foo"},
					},
					Rights: []ast.Node{
						asttest.NewLiteralString("bar"),
					},
				},
				&ast.Key{
					Expr: &ast.Identifier{Name: "foo"},
					Key:  asttest.NewLiteralNumber("1"),
				},
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        asttest.NewLiteralString("bar"),
				},
				&vm.Assign{
					VariableName: "foo",
					Register:     "1",
				},

				// foo[1]
				&vm.Assign{
					VariableName: "2",
					Value:        asttest.NewLiteralNumber("1"),
				},
				&vm.StringIndex{
					String: "foo",
					Index:  "2",
					Result: "3",
				},
			},
		},
	} {
		t.Run(testName, func(t *testing.T) {
			compiledFunc, err := compiler.CompileFunc(&ast.Func{
				Statements: test.nodes,
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
