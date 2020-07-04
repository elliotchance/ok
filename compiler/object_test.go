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

func TestObject(t *testing.T) {
	for testName, test := range map[string]struct {
		node     *ast.Func
		expected []vm.Instruction
		err      error
	}{
		"empty-object": {
			node: &ast.Func{
				Name:    "Foo",
				Returns: []string{"Foo"},
			},
			expected: []vm.Instruction{
				// alloc instance
				&vm.Assign{
					VariableName: "1",
					Value:        asttest.NewLiteralNumber("0"),
				},
				&vm.MapAllocNumber{
					Size:   "1",
					Result: "2",
				},

				// return instance
				&vm.Return{
					Results: []string{"2"},
				},
			},
		},
		"one-variable": {
			node: &ast.Func{
				Name:    "Foo",
				Returns: []string{"Foo"},
				Statements: []ast.Node{
					&ast.Assign{
						Lefts: []ast.Node{
							&ast.Identifier{Name: "bar"},
						},
						Rights: []ast.Node{
							asttest.NewLiteralNumber("123"),
						},
					},
				},
			},
			expected: []vm.Instruction{
				// alloc instance
				&vm.Assign{
					VariableName: "1",
					Value:        asttest.NewLiteralNumber("0"),
				},
				&vm.MapAllocNumber{
					Size:   "1",
					Result: "2",
				},

				// bar = 123
				&vm.Assign{
					VariableName: "3",
					Value:        asttest.NewLiteralNumber("123"),
				},
				&vm.Assign{
					VariableName: "4",
					Value:        asttest.NewLiteralString("bar"),
				},
				&vm.MapSet{
					Map:   "2",
					Key:   "4",
					Value: "3",
				},

				// return instance
				&vm.Return{
					Results: []string{"2"},
				},
			},
		},
	} {
		t.Run(testName, func(t *testing.T) {
			compiledFunc, err := compiler.CompileFunc(test.node, nil)
			if test.err != nil {
				assert.EqualError(t, err, test.err.Error())
			} else {
				require.NoError(t, err)
				assert.Equal(t, test.expected, compiledFunc.Instructions)
			}
		})
	}
}
