package vm_test

import (
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/vm"

	"github.com/stretchr/testify/assert"
)

func TestAssignSymbol_Execute(t *testing.T) {
	registers := map[vm.Register]*ast.Literal{}
	ins := &vm.AssignSymbol{
		Result: "1",
		Symbol: "123",
	}
	vm := &vm.VM{
		Stack: []map[vm.Register]*ast.Literal{registers},
		Symbols: map[vm.SymbolRegister]*ast.Literal{
			"123": asttest.NewLiteralNumber("1.5"),
		},
	}
	assert.NoError(t, ins.Execute(nil, vm))
	assert.Equal(t, "1.5", registers["1"].Value)
}

func TestAssign_Execute(t *testing.T) {
	registers := map[vm.Register]*ast.Literal{
		"0": asttest.NewLiteralNumber("1.5"),
	}
	ins := &vm.Assign{
		Result:   "1",
		Register: "0",
	}
	vm := &vm.VM{
		Stack: []map[vm.Register]*ast.Literal{registers},
	}
	assert.NoError(t, ins.Execute(nil, vm))
	assert.Equal(t, "1.5", registers["1"].Value)
}
