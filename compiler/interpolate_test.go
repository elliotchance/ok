package compiler_test

import (
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/compiler"
	"github.com/elliotchance/ok/lexer"
	"github.com/elliotchance/ok/vm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInterpolate(t *testing.T) {
	for testName, test := range map[string]struct {
		node     ast.Node
		expected []vm.Instruction
	}{
		"simple": {
			node: &ast.Interpolate{
				Parts: []ast.Node{
					asttest.NewLiteralString("hi "),
					&ast.Group{
						Expr: asttest.NewBinary(
							asttest.NewLiteralNumber("1"),
							lexer.TokenPlus,
							asttest.NewLiteralNumber("2"),
						),
					},
					asttest.NewLiteralString(" there"),
				},
			},
			expected: []vm.Instruction{
				&vm.AssignSymbol{
					Result: "2",
					Symbol: "stringhi~~~",
				},
				&vm.AssignSymbol{
					Result: "3",
					Symbol: "number1",
				},
				&vm.AssignSymbol{
					Result: "4",
					Symbol: "number2",
				},
				&vm.Add{
					Left:   "3",
					Right:  "4",
					Result: "5",
				},
				&vm.AssignSymbol{
					Result: "6",
					Symbol: "string~~~there",
				},
				&vm.Interpolate{
					Result: "1",
					Args:   []vm.Register{"2", "5", "6"},
				},
			},
		},
	} {
		t.Run(testName, func(t *testing.T) {
			compiledFunc, err := compiler.CompileFunc(&ast.Func{
				Statements: []ast.Node{
					test.node,
				},
			}, &vm.File{
				Symbols: map[vm.SymbolRegister]*vm.Symbol{},
			}, nil, nil)
			require.NoError(t, err)
			assert.Equal(t, test.expected, compiledFunc.Instructions.Instructions)
		})
	}
}
