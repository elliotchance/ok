package compiler_test

import (
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/compiler"
	"github.com/elliotchance/ok/vm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAssign(t *testing.T) {
	for testName, test := range map[string]struct {
		nodes    []ast.Node
		expected []vm.Instruction
		err      error
	}{
		"assign-1": {
			nodes: []ast.Node{
				&ast.Assign{
					Lefts: []ast.Node{
						&ast.Identifier{Name: "foo"},
					},
					Rights: []ast.Node{
						ast.NewLiteralNumber("1.5"),
					},
				},
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("1.5"),
				},
				&vm.Assign{
					VariableName: "foo",
					Register:     "1",
				},
			},
		},
		"assign-multiple-literals": {
			nodes: []ast.Node{
				&ast.Assign{
					Lefts: []ast.Node{
						&ast.Identifier{Name: "foo"},
						&ast.Identifier{Name: "bar"},
					},
					Rights: []ast.Node{
						ast.NewLiteralNumber("1.5"),
						ast.NewLiteralNumber("3.0"),
					},
				},
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("1.5"),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralNumber("3.0"),
				},
				&vm.Assign{
					VariableName: "foo",
					Register:     "1",
				},
				&vm.Assign{
					VariableName: "bar",
					Register:     "2",
				},
			},
		},
	} {
		t.Run(testName, func(t *testing.T) {
			compiledFunc, err := compiler.CompileFunc(newFunc(test.nodes...), nil)
			if test.err != nil {
				assert.EqualError(t, err, test.err.Error())
			} else {
				require.NoError(t, err)
				assert.Equal(t, test.expected, compiledFunc.Instructions)
			}
		})
	}
}
