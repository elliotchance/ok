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
					Finished: true,
					Type:     "1",
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
					Finished: true,
					Type:     "1",
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
					Type: "1",
				},
				&vm.Jump{
					To: 4,
				},

				&vm.On{
					Finished: true,
					Type:     "2",
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
							Type: "SomeError",
						},
						{
							Type: "SomethingElse",
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
					Type: "1",
				},
				&vm.Jump{
					To: 7,
				},

				&vm.On{
					Type: "2",
				},
				&vm.Print{},
				&vm.Jump{
					To: 7,
				},

				&vm.On{
					Finished: true,
					Type:     "3",
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
					Finished: true,
					Type:     "1",
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
				&vm.File{
					Types: types.Registry{},
				}, nil, map[string]*ast.Literal{
					"SomeError": {
						Kind: types.NewFunc(nil, []*types.Type{types.String}),
					},
					"SomethingElse": {
						Kind: types.NewFunc(nil, []*types.Type{types.Number}),
					},
				}, nil, nil)
			if test.err != nil {
				assert.EqualError(t, err, test.err.Error())
			} else {
				require.NoError(t, err)
				assert.Equal(t, test.expected, compiledFunc.Instructions.Instructions)
			}
		})
	}
}
