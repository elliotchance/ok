package compiler_test

import (
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/compiler"
	"github.com/elliotchance/ok/types"
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
					Type: nil,
				},
			},
		},
		"only-try": {
			nodes: []ast.Node{
				&ast.ErrorScope{
					Statements: []ast.Node{
						&ast.Call{
							Expr: &ast.Identifier{Name: "print"},
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
					Type: nil,
				},
			},
		},
		"try-on-1": {
			nodes: []ast.Node{
				&ast.ErrorScope{
					Statements: []ast.Node{
						&ast.Call{
							Expr: &ast.Identifier{Name: "print"},
						},
					},
					On: []*ast.On{
						{
							Type: types.NewInterface("SomeError", nil),
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
					Type: types.NewInterface("SomeError", nil),
				},
				&vm.Jump{
					To: 4,
				},

				&vm.On{
					Type: nil,
				},
			},
		},
		"try-on-2": {
			nodes: []ast.Node{
				&ast.ErrorScope{
					Statements: []ast.Node{
						&ast.Call{
							Expr: &ast.Identifier{Name: "print"},
						},
					},
					On: []*ast.On{
						{
							Type: types.NewInterface("SomeError", nil),
						},
						{
							Type: types.NewInterface("SomethingElse", nil),
							Statements: []ast.Node{
								&ast.Call{
									Expr: &ast.Identifier{Name: "print"},
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
					Type: types.NewInterface("SomeError", nil),
				},
				&vm.Jump{
					To: 7,
				},

				&vm.On{
					Type: types.NewInterface("SomethingElse", nil),
				},
				&vm.Print{},
				&vm.Jump{
					To: 7,
				},

				&vm.On{
					Type: nil,
				},
			},
		},
		"try-finally": {
			nodes: []ast.Node{
				&ast.ErrorScope{
					Statements: []ast.Node{
						&ast.Call{
							Expr: &ast.Identifier{Name: "print"},
						},
					},
					Finally: &ast.Finally{
						Statements: []ast.Node{
							&ast.Call{
								Expr: &ast.Identifier{Name: "print"},
							},
						},
					},
				},
			},
			expected: []vm.Instruction{
				// The moment we enter the try, we enable the finally block.
				&vm.Finally{
					Index: 0,
					Run:   true,
				},

				&vm.Print{},
				&vm.Jump{
					To: 3,
				},

				&vm.On{
					Type: nil,
				},

				// If we enter the finally block we need to disable it, this
				// prevents it from running again when we return.
				&vm.Finally{
					Index: 0,
					Run:   false,
				},
				&vm.Print{},
			},
		},
	} {
		t.Run(testName, func(t *testing.T) {
			compiledFunc, err := compiler.CompileFunc(newFunc(test.nodes...),
				&vm.File{}, nil)
			if test.err != nil {
				assert.EqualError(t, err, test.err.Error())
			} else {
				require.NoError(t, err)
				assert.Equal(t, test.expected, compiledFunc.Instructions)
			}
		})
	}
}
