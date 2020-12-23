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
						asttest.NewLiteralNumber("1.5"),
					},
				},
			},
			expected: []vm.Instruction{
				&vm.AssignSymbol{
					Result: "1",
					Symbol: "number1.5",
				},
				&vm.Assign{
					Result:   "foo",
					Register: "1",
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
						asttest.NewLiteralNumber("1.5"),
						asttest.NewLiteralNumber("3.0"),
					},
				},
			},
			expected: []vm.Instruction{
				&vm.AssignSymbol{
					Result: "1",
					Symbol: "number1.5",
				},
				&vm.AssignSymbol{
					Result: "2",
					Symbol: "number3.0",
				},
				&vm.Assign{
					Result:   "foo",
					Register: "1",
				},
				&vm.Assign{
					Result:   "bar",
					Register: "2",
				},
			},
		},
		"key-assign": {
			nodes: []ast.Node{
				// TODO(elliot): This will fail once the compiler is smart
				//  enough to know foo cannot be used an object.
				&ast.Assign{
					Lefts: []ast.Node{
						&ast.Identifier{Name: "foo"},
					},
					Rights: []ast.Node{
						asttest.NewLiteralNumber("1.5"),
					},
				},
				&ast.Assign{
					Lefts: []ast.Node{
						&ast.Key{
							Expr: &ast.Identifier{Name: "foo"},
							Key:  &ast.Identifier{Name: "bar"},
						},
					},
					Rights: []ast.Node{
						asttest.NewLiteralNumber("1.5"),
					},
				},
			},
			expected: []vm.Instruction{
				&vm.AssignSymbol{
					Result: "1",
					Symbol: "number1.5",
				},
				&vm.Assign{
					Result:   "foo",
					Register: "1",
				},
				&vm.AssignSymbol{
					Result: "2",
					Symbol: "number1.5",
				},
				&vm.AssignSymbol{
					Result: "3",
					Symbol: "stringbar",
				},
				&vm.MapSet{
					Map:   "foo",
					Key:   "3",
					Value: "2",
				},
			},
		},
	} {
		t.Run(testName, func(t *testing.T) {
			compiledFunc, err := compiler.CompileFunc(newFunc(test.nodes...),
				&vm.File{
					Symbols: map[vm.SymbolRegister]*vm.Symbol{},
				}, nil)
			if test.err != nil {
				assert.EqualError(t, err, test.err.Error())
			} else {
				require.NoError(t, err)
				assert.Equal(t, test.expected, compiledFunc.Instructions.Instructions)
			}
		})
	}
}
