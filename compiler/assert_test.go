package compiler_test

import (
	"errors"
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

func TestAssert(t *testing.T) {
	for testName, test := range map[string]struct {
		node     *ast.Assert
		expected []vm.Instruction
		err      error
	}{
		"not-a-bool": {
			node: &ast.Assert{
				Expr: asttest.NewBinary(
					asttest.NewLiteralNumber("1"),
					lexer.TokenPlus,
					asttest.NewLiteralNumber("2"),
				),
			},
			err: errors.New("assert condition must be a bool but is number"),
		},
		"success": {
			node: &ast.Assert{
				Expr: asttest.NewBinary(
					asttest.NewLiteralNumber("1"),
					lexer.TokenEqual,
					asttest.NewLiteralNumber("2"),
				),
			},
			expected: []vm.Instruction{
				&vm.AssignSymbol{
					Result: "1",
					Symbol: "0",
				},
				&vm.AssignSymbol{
					Result: "2",
					Symbol: "1",
				},
				&vm.EqualNumber{
					Left:   "1",
					Right:  "2",
					Result: "3",
				},
				&vm.Assert{
					Left:  "1",
					Op:    "==",
					Right: "2",
					Final: "3",
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
				Types:   types.Registry{},
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
