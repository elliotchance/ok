package compiler_test

import (
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/compiler"
	"github.com/elliotchance/ok/lexer"
	"github.com/elliotchance/ok/types"
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
					Expr: &ast.Identifier{Name: "print"},
				},
			},
			expected: []vm.Instruction{
				&vm.Print{},
			},
		},
		"len-1": {
			nodes: []ast.Node{
				&ast.Call{
					Expr: &ast.Identifier{Name: "len"},
					Arguments: []ast.Node{
						asttest.NewArrayNumbers(nil),
					},
				},
			},
			expected: []vm.Instruction{
				&vm.AssignSymbol{
					Result: "1",
					Symbol: "0",
				},
				&vm.ArrayAlloc{
					Size:   "1",
					Result: "2",
					Kind:   "2",
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
						asttest.NewLiteralNumber("1.5"),
					},
				},
				&ast.Call{
					Expr: &ast.Identifier{Name: "print"},
					Arguments: []ast.Node{
						&ast.Identifier{Name: "foo"},
					},
				},
			},
			expected: []vm.Instruction{
				&vm.AssignSymbol{
					Result: "1",
					Symbol: "0",
				},
				&vm.Assign{
					Result:   "foo",
					Register: "1",
				},
				&vm.Print{
					Arguments: []vm.Register{"foo"},
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
						asttest.NewLiteralNumber("1.5"),
					},
				},
				&ast.Call{
					Expr: &ast.Identifier{Name: "print"},
					Arguments: []ast.Node{
						&ast.Binary{
							Left:  &ast.Identifier{Name: "foo"},
							Op:    lexer.TokenPlus,
							Right: asttest.NewLiteralNumber("2"),
						},
						&ast.Binary{
							Left:  asttest.NewLiteralNumber("10"),
							Op:    lexer.TokenTimes,
							Right: &ast.Identifier{Name: "foo"},
						},
					},
				},
			},
			expected: []vm.Instruction{
				&vm.AssignSymbol{
					Result: "1",
					Symbol: "0",
				},
				&vm.Assign{
					Result:   "foo",
					Register: "1",
				},
				&vm.AssignSymbol{
					Result: "2",
					Symbol: "1",
				},
				&vm.Add{
					Left:   "foo",
					Right:  "2",
					Result: "3",
				},
				&vm.AssignSymbol{
					Result: "4",
					Symbol: "2",
				},
				&vm.Multiply{
					Left:   "4",
					Right:  "foo",
					Result: "5",
				},
				&vm.Print{
					Arguments: []vm.Register{"3", "5"},
				},
			},
		},
		"any-string": {
			nodes: []ast.Node{
				&ast.Call{
					Expr:      &ast.Identifier{Name: "any"},
					Arguments: []ast.Node{asttest.NewLiteralString("foo")},
				},
			},
			expected: []vm.Instruction{
				&vm.AssignSymbol{
					Result: "1",
					Symbol: "0",
				},
			},
		},
	} {
		t.Run(testName, func(t *testing.T) {
			compiledFunc, err := compiler.CompileFunc(newFunc(test.nodes...),
				&vm.File{
					Types:   types.Registry{},
					Symbols: map[vm.SymbolRegister]*vm.Symbol{},
				}, nil, nil, nil, nil)
			if test.err != nil {
				assert.EqualError(t, err, test.err.Error())
			} else {
				require.NoError(t, err)
				assert.Equal(t, test.expected, compiledFunc.Instructions.Instructions)
			}
		})
	}
}
