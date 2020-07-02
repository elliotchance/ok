package compiler_test

import (
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/compiler"
	"github.com/elliotchance/ok/lexer"
	"github.com/elliotchance/ok/vm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCall(t *testing.T) {
	for testName, test := range map[string]struct {
		nodes    []ast.Node
		expected []vm.Instruction
		err      error
	}{
		"print-0": {
			nodes: []ast.Node{
				&ast.Call{
					FunctionName: "print",
				},
			},
			expected: []vm.Instruction{
				&vm.Print{},
			},
		},
		"len-1": {
			nodes: []ast.Node{
				&ast.Call{
					FunctionName: "len",
					Arguments: []ast.Node{
						ast.NewArrayNumbers(nil),
					},
				},
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("0"),
				},
				&vm.ArrayAllocNumber{
					Size:   "1",
					Result: "2",
				},
				&vm.Len{
					Argument: "2",
					Result:   "3",
				},
			},
		},
		"assign-print": {
			nodes: []ast.Node{
				&ast.Assign{
					Lefts: []ast.Node{
						&ast.Identifier{Name: "foo"},
					},
					Rights: []ast.Node{
						ast.NewLiteralNumber("1.5"),
					},
				},
				&ast.Call{
					FunctionName: "print",
					Arguments: []ast.Node{
						&ast.Identifier{Name: "foo"},
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
				&vm.Print{
					Arguments: []string{"foo"},
				},
			},
		},
		"assign-print-2": {
			nodes: []ast.Node{
				&ast.Assign{
					Lefts: []ast.Node{
						&ast.Identifier{Name: "foo"},
					},
					Rights: []ast.Node{
						ast.NewLiteralNumber("1.5"),
					},
				},
				&ast.Call{
					FunctionName: "print",
					Arguments: []ast.Node{
						&ast.Binary{
							Left:  &ast.Identifier{Name: "foo"},
							Op:    lexer.TokenPlus,
							Right: ast.NewLiteralNumber("2"),
						},
						&ast.Binary{
							Left:  ast.NewLiteralNumber("10"),
							Op:    lexer.TokenTimes,
							Right: &ast.Identifier{Name: "foo"},
						},
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
				&vm.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralNumber("2"),
				},
				&vm.Add{
					Left:   "foo",
					Right:  "2",
					Result: "3",
				},
				&vm.Assign{
					VariableName: "4",
					Value:        ast.NewLiteralNumber("10"),
				},
				&vm.Multiply{
					Left:   "4",
					Right:  "foo",
					Result: "5",
				},
				&vm.Print{
					Arguments: []string{"3", "5"},
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
