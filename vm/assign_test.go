package vm_test

import (
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/vm"

	"github.com/stretchr/testify/assert"
)

func TestAssign_Execute(t *testing.T) {
	t.Run("literal", func(t *testing.T) {
		registers := map[vm.Register]*ast.Literal{}
		ins := &vm.Assign{VariableName: "1", Value: asttest.NewLiteralNumber("1.5")}
		vm := &vm.VM{
			Stack: []map[vm.Register]*ast.Literal{registers},
		}
		assert.NoError(t, ins.Execute(nil, vm))
		assert.Equal(t, "1.5", registers["1"].Value)
	})

	t.Run("register", func(t *testing.T) {
		registers := map[vm.Register]*ast.Literal{
			"0": asttest.NewLiteralNumber("1.5"),
		}
		ins := &vm.Assign{VariableName: "1", Register: "0"}
		vm := &vm.VM{
			Stack: []map[vm.Register]*ast.Literal{registers},
		}
		assert.NoError(t, ins.Execute(nil, vm))
		assert.Equal(t, "1.5", registers["1"].Value)
	})
}
