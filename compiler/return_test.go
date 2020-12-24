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

func TestReturn(t *testing.T) {
	for testName, test := range map[string]struct {
		node     ast.Node
		expected []vm.Instruction
		err      error
	}{
		"return-zero": {
			node: &ast.Return{},
			expected: []vm.Instruction{
				&vm.Return{},
			},
		},
		"return-number": {
			node: &ast.Return{
				Exprs: []ast.Node{
					asttest.NewLiteralNumber("123"),
				},
			},
			expected: []vm.Instruction{
				&vm.AssignSymbol{
					Result: "1",
					Symbol: "number123",
				},
				&vm.Return{
					Results: []vm.Register{"1"},
				},
			},
		},
		"return-multiple-values": {
			node: &ast.Return{
				Exprs: []ast.Node{
					asttest.NewLiteralNumber("123"),
					asttest.NewLiteralString("foo"),
				},
			},
			expected: []vm.Instruction{
				&vm.AssignSymbol{
					Result: "1",
					Symbol: "number123",
				},
				&vm.AssignSymbol{
					Result: "2",
					Symbol: "stringfoo",
				},
				&vm.Return{
					Results: []vm.Register{"1", "2"},
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
			if test.err != nil {
				assert.EqualError(t, err, test.err.Error())
			} else {
				require.NoError(t, err)
				assert.Equal(t, test.expected, compiledFunc.Instructions.Instructions)
			}
		})
	}
}
