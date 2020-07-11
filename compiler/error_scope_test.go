package compiler_test

import (
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/compiler"
	"github.com/elliotchance/ok/vm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestErrorScope(t *testing.T) {
	for testName, test := range map[string]struct {
		nodes    []ast.Node
		expected []vm.Instruction
		err      error
	}{
		"only-empty-try": {
			nodes: []ast.Node{
				&ast.ErrorScope{},
			},
			expected: []vm.Instruction{
				&vm.Jump{
					To: 1,
				},

				&vm.On{
					Type: "",
				},
			},
		},
		"only-try": {
			nodes: []ast.Node{
				&ast.ErrorScope{
					Statements: []ast.Node{
						&ast.Call{
							FunctionName: "print",
						},
					},
				},
			},
			expected: []vm.Instruction{
				&vm.Print{},
				&vm.Jump{
					To: 2,
				},

				&vm.On{
					Type: "",
				},
			},
		},
		"try-on-1": {
			nodes: []ast.Node{
				&ast.ErrorScope{
					Statements: []ast.Node{
						&ast.Call{
							FunctionName: "print",
						},
					},
					On: []*ast.On{
						{
							Type: "SomeError",
						},
					},
				},
			},
			expected: []vm.Instruction{
				&vm.Print{},
				&vm.Jump{
					To: 4,
				},

				&vm.On{
					Type: "SomeError",
				},
				&vm.Jump{
					To: 4,
				},

				&vm.On{
					Type: "",
				},
			},
		},
		"try-on-2": {
			nodes: []ast.Node{
				&ast.ErrorScope{
					Statements: []ast.Node{
						&ast.Call{
							FunctionName: "print",
						},
					},
					On: []*ast.On{
						{
							Type: "SomeError",
						},
						{
							Type: "SomethingElse",
							Statements: []ast.Node{
								&ast.Call{
									FunctionName: "print",
								},
							},
						},
					},
				},
			},
			expected: []vm.Instruction{
				&vm.Print{},
				&vm.Jump{
					To: 7,
				},

				&vm.On{
					Type: "SomeError",
				},
				&vm.Jump{
					To: 7,
				},

				&vm.On{
					Type: "SomethingElse",
				},
				&vm.Print{},
				&vm.Jump{
					To: 7,
				},

				&vm.On{
					Type: "",
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
